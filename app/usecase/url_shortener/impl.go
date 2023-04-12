package url_shortener

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/3FanYu/url-shortener-go/domain/url_shortener"
	"github.com/go-redis/redis"
)

type impl struct {
	repo  ShortUrlRepository
	cache *redis.Client
}

// salt is 2023/01/01 00:00:00:000 timestamp
const salt = 1672502400000
const maxCount = 10000
const cacheTtl = 10 * time.Second

var (
	timeNow        = time.Now
	counter uint64 = 0
)

func NewUsecase(cacheSrv *redis.Client, repo ShortUrlRepository) ShortUrlUsecase {
	im := &impl{
		repo:  repo,
		cache: cacheSrv,
	}

	return im
}

func (im *impl) FindOrCreate(ctx context.Context, shortUrl *url_shortener.ShortUrl) (*url_shortener.ShortUrl, error) {
	shortUrl.Key = generateKey()
	r, err := im.repo.FindOrCreate(ctx, shortUrl)
	if err != nil {
		return nil, err
	}
	json, err := json.Marshal(shortUrl)
	if err != nil {
		log.Printf("JSON Marshal failed: %v", err)
	}
	err = im.cache.Set(r.Key, json, cacheTtl).Err()
	if err != nil {
		log.Printf("Cache Set failed: %v", err)
	}
	return r, nil
}

func (im *impl) RedirectToShortUrl(ctx context.Context, key string) (*url_shortener.ShortUrl, error) {
	var shortUrl url_shortener.ShortUrl
	val, err := im.cache.Get(key).Result()
	if err != nil {
		log.Printf("Cache Get failed: %v", err)
	}
	err = json.Unmarshal([]byte(val), &shortUrl)
	if err == nil {
		log.Println("Cache hit successfully")
		return &shortUrl, nil
	}
	log.Printf("JSON Unmarshal failed: %v", err)

	r, err := im.repo.GetByKey(ctx, key)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func generateKey() string {
	curTimeStamp := timeNow().Round(time.Millisecond).UnixNano() / 1000000

	atomic.AddUint64(&counter, 1)

	return big.NewInt((curTimeStamp-salt)*maxCount + (int64(counter % maxCount))).Text(62)
}
