package url_shortener

import (
	"context"

	domain "github.com/3FanYu/url-shortener-go/domain/url_shortener"
)

// EventRepository deals with events related video
type ShortUrlRepository interface {
	// BatchCreate creates list of events in batch
	FindOrCreate(ctx context.Context, shortUrl *domain.ShortUrl) (*domain.ShortUrl, error)
	GetByKey(ctx context.Context, key string) (*domain.ShortUrl, error)
}

type ShortUrlUsecase interface {
	// BatchCreate creates list of events in batch
	FindOrCreate(ctx context.Context, shortUrl *domain.ShortUrl) (*domain.ShortUrl, error)
	RedirectToShortUrl(ctx context.Context, key string) (*domain.ShortUrl, error)
}
