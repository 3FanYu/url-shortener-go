package url_shortener

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortUrl struct {
	ID        primitive.ObjectID
	Key       string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
