package url_shortener

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	domain "github.com/3FanYu/url-shortener-go/domain/url_shortener"
	uc "github.com/3FanYu/url-shortener-go/usecase/url_shortener"
)

type shortUrl struct {
	ID  string `bson:"_id,omitempty"`
	Key string `bson:"key"`
	Url string `bson:"url"`
}

type shortUrlRepository struct {
	db *mongo.Database
}

const collection = "short_urls"

func NewRepository(db *mongo.Database) uc.ShortUrlRepository {
	return &shortUrlRepository{db: db}
}

func (r *shortUrlRepository) Create(ctx context.Context, shortUrl *domain.ShortUrl) (*domain.ShortUrl, error) {
	_, err := r.db.Collection(collection).InsertOne(ctx, shortUrl)
	if err != nil {
		return nil, err
	}
	return shortUrl, nil
}

func (r *shortUrlRepository) GetByKey(ctx context.Context, key string) (*domain.ShortUrl, error) {
	var mapping domain.ShortUrl
	err := r.db.Collection(collection).FindOne(ctx, shortUrl{Key: key}).Decode(&mapping)
	return &mapping, err
}
