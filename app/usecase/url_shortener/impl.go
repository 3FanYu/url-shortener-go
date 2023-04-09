package url_shortener

import (
	"context"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/3FanYu/url-shortener-go/domain/url_shortener"
)

type impl struct {
	repo ShortUrlRepository
}

// salt is 2023/01/01 00:00:00:000 timestamp
const salt = 1672502400000
const maxCount = 10000

var (
	timeNow        = time.Now
	counter uint64 = 0
)

func NewUsecase(repo ShortUrlRepository) ShortUrlUsecase {
	im := &impl{repo: repo}
	// im.cache = cacheSrv.Create([]cachekit.Setting{
	// 	{
	// 		Prefix: pfxRecord,
	// 		CacheAttributes: map[cachekit.Type]cachekit.Attribute{
	// 			cachekit.SharedCacheType: {TTL: time.Minute},
	// 			cachekit.LocalCacheType:  {TTL: 10 * time.Second},
	// 		},
	// 	},
	// })

	return im
}

func (im *impl) Create(ctx context.Context, shortUrl *url_shortener.ShortUrl) (*url_shortener.ShortUrl, error) {
	shortUrl.Key = generateKey()
	r, err := im.repo.Create(ctx, shortUrl)
	if err != nil{
		return nil, err
	}
	return addUrlPrefix(r), nil
}

func (im *impl) RedirectToShortUrl(ctx context.Context, key string) (*url_shortener.ShortUrl, error) {
	return im.repo.GetByKey(ctx, key)
}

func generateKey() string {
	curTimeStamp := timeNow().Round(time.Millisecond).UnixNano() / 1000000

	atomic.AddUint64(&counter, 1)

	return big.NewInt((curTimeStamp-salt)*maxCount + (int64(counter % maxCount))).Text(62)
}

func addUrlPrefix(shortUrl *url_shortener.ShortUrl) *url_shortener.ShortUrl {
	shortUrl.Key = "http://localhost/" + shortUrl.Key
	return shortUrl
}
