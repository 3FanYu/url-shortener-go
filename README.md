# Setting up the project

1. Make sure you've installed docker and docker-compose
2. Make sure you're in the root directory of the project
3. Run `make up` to start services
4. Run below command to test out the API, feel free to switch up the "url" with your favorite website
```
curl --location --request POST 'localhost:3000/api/short_url' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "https://youtube.com"
}'
```
6. You should see return value as below
```
{"shortUrl":"http://localhost:3000/p0YDXMuR","originalUrl":"https://youtube.com"}%
```
7. Paste the `shortUrl` to your browser, and you should be redirected to your favorite website!


# API level design 
1. POST http://{host}/api/shorturl
    - Param {string} url
    - Return shortened url
2. GET http://{host}/{shortened_url}
    - Accept shortened_url
    - Redirect to original url if record found
    - Return 404 if no record found or record expired


# Implementation 
## Shortening url:
- [x] Check if received url exists in DB, return directly if true.
- [x] Generate an unique ID using timestamp concat with a counter
- [x] Use Base62 encode to generate short_url from unique ID.
- [x] Store the unique ID together with the original_url into cache.
- [x] Store the unique ID together with the original_url into database.
- [x] Return short url as a response to client
- [x] Database collection will be set up with TTL index, so expired data will be deleted automatically.


## Access original url using short url
- [x] Use Base62 decoding with the received short_url
- [x] Take the decoded result as the key to search in cache to get the original_url
- [x] Take the decoded result as the key to search in database to get the original_url
- [x] Response 302 and redirect client browser to original_url


## Road Map

### About redirection
We have two kinds of http status code to choose from: 301 and 302. 301 means “move permanently”, it will make the client browser cache the redirection. 302 means “move temporarily”, which will not cache the redirection. In this case we will use 302 because we want the short url to expire, it would be difficult to test if the browser caches the redirection. However, in real life people might make decision based on SEO issues.

### How to ensure uniqueness of short url upon generating?
We can use timestamp in the form of milliseconds together with Base62 encoding to ensure the uniqueness of every ID.
Timestamp in milliseconds will have the length of 13 digits right now. While most of the prefix digits will not be changing in the near future, we can make our timestamp shorter to have a more concise short url generated.
Timestamp in milliseconds will be 13 digits in length EX: 1677502170731
We can make 2023/1/1 00:00:00 UTC time as our salt: 1672502400000
Subtract the numbers to get less digits result 1677502170731 - 1672502400000 = 4999770731
While timestamp looks great, there is still a chance when two requests were received at the exact same millisecond, and eventually lead to duplicate unique ID. To deal with that, we can concat a counter behind the timestamp:
EX: Merge the timestamp with a 4 digits counter EX: 49997707310001
*Lock will need to be implemented on the counter to prevent race condition.
Generating short urls
Hash functions like MD5 are generally fast and simple to implement. However there is always a possibility of collision, where two different inputs produces the same output.
Converting Base10 numbers to Base62 on the other hand, does not have collision risk of hash functions, as long as the encoded strings and of consistent length. Nevertheless, Base62 encoding & decoding can be slower than the hash functions. While performance issue is not significant in this usecase Try it out: https://math.tools/calculator/base/10-62
 For this project I will go with Base62 for the purpose of collision risk. 

### ID as the index key in DB
Now that we have short urls generated by  unique_ids, we can store it into the database. We will be mapping a short url to its original url often, so my initial thought was to store short url as the index key, and map the url by querying: 
`select * from urls where short_url = <short_url>`
But as mentioned earlier, seting up index on integers will perform much better than string. Therefore we can store the unique_id as the index key instead. Everytime we receive request to map a short_url, we can decode the short_url back to its unique_id form, and query it like:
`select * from urls where uid = <unique_id>`

### Cache the mapping result
It’s not difficult to foresee the web app will be facing a lot of redirection request. Searching the short_url to original_url mapping directly from the database will get relatively expensive when facing high traffic, so it is better to cache the mapping result. There are many tools to choose from when it comes to memory caching, Redis is one of the no brainer that is used by many production services. But how should we cache it? One of the predictable behaviour of the user is to test the newly generated short url. Therefore we can cache the result into redis before even actually saving it into the database. In other cases, we can save the result of every redirect requests as a cache, however the TTL of each cache will be a very important factor if you don’t wish to store too much cache data when facing high traffic.

### Choosing the databse
Since the application involves heavy read and write, and not much table joining. I will be choosing a NoSQL database as my primary database due to faster read and write speed, high availability, and ease of scaling horizontally.

I will be using MongoDB for this project since it’s easier to setup locally. However, in production we could use cloud services such as AWS DynamoDB.

Short url expiration
My initial thought is to setup a cron job that runs and check in the database to delete the expired data. Since we will be storing the `created_at` field when we store the data, there will be no problem of querying and deleting the expired data.
EX: ShortUrlMapping.where(‘created_at < ?‘, Time.now - TTL).destroy

However, while choosing the database, I discovered a feature of MongoDB called TTL index that can achieve this feature. This way it is even easier to implement.
Ref: https://www.mongodb.com/docs/manual/core/index-ttl/

### Implementation 
Shortening url:
Check if received url exists in DB, return directly if true.
Generate an unique ID using timestamp concat with a counter
Use Base62 encode to generate short_url from unique ID.
Store the unique ID together with the original_url into cache and database.
Return short url as a response to client
Database collection will be set up with TTL index, so expired data will be deleted automatically.

Access original url using short url:
Use Base62 decoding with the received short_url
Take the decoded result as the key to search in cache or database to get the original_url
Response 302 and redirect client browser to original_url


### Go Beyond
Unfortunately, Implementation V2 still faces some problems when we scale:
Small chance of unique ID collision (happens when multiple services receive request at exact same time and while counters are luckily on the same number) 
Heavy writes.

To deal with problem 1, we will need to introduce a centralized service that maintains the configuration of all the counters inside every service. EX: Apache Zookeeper.
For instance, the centralized service can hold configuration of the counter’s range of counting for all the counters. EX: serviceA counter = 1~10000, serviceB counter = 10000 ~ 20000. That way we may never have duplicated unique ID be generated.

Regarding problem 2, we can implement the batch insertion mechanism, where we don’t directly write into the database. Instead, we will store the data inside of a cache and respond back to client directly after sending an event to a message queue. Later on the message queue will be pulled, then the data will eventually be inserted into the database. This way we won’t have to wait for database insertion before responding to our users.
