package url_shortener

import "time"

type ShortUrl struct {
	Key       string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
