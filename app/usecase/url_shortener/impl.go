package url_shortener

import (
	"context"
	"github.com/3FanYu/url-shortener-go/domain/url_shortener"
)

type impl struct {
	repo ShortUrlRepository
}

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

func (im *impl) Create(ctx context.Context, url string) error {
	return im.repo.Create(ctx, url)
}

func (im *impl) RedirectToShortUrl(ctx context.Context, key string) (*url_shortener.ShortUrl, error) {
	return im.repo.GetByKey(ctx, key)
}
