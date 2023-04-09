package url_shortener

import (
	"context"

	urlShortenerDomain "github.com/3FanYu/url-shortener-go/domain/url_shortener"
)

// EventRepository deals with events related video
type ShortUrlRepository interface {
	// BatchCreate creates list of events in batch
	Create(ctx context.Context, url string) error
	GetByKey(ctx context.Context, key string) (*urlShortenerDomain.ShortUrl, error)
}

type ShortUrlUsecase interface {
	// BatchCreate creates list of events in batch
	Create(ctx context.Context, url string) error
	RedirectToShortUrl(ctx context.Context, key string) (*urlShortenerDomain.ShortUrl, error)
}
