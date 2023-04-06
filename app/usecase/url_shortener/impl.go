package url_shortener

import (
	"context"
	"github.com/3FanYu/url-shortener-go/domain/url_shortener"
)

type impl struct {
	repo ShortUrlRepository
}

func NewUrlShortenerUsecase(repo ShortUrlRepository) ShortUrlUsecase {
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

func (im *impl) Create(ctx context.Context, mapping *url_shortener.ShortUrl) error {
	return im.repo.Create(ctx, mapping)
}

func (im *impl) GetByKey(ctx context.Context, key string) (*url_shortener.ShortUrl, error) {
	return im.repo.GetByKey(ctx, key)
}
