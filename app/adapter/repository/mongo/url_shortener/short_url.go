package url_shortener

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "github.com/3FanYu/url-shortener-go/domain/url_shortener"
	uc "github.com/3FanYu/url-shortener-go/usecase/url_shortener"
)

type shortUrl struct {
	ID        string `bson:"_id,omitempty"`
	Key       string `bson:"key"`
	Url       string `bson:"url"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}

type shortUrlRepository struct {
	db *mongo.Database
}

const collection = "short_urls"
const urlTtl = 10

func NewRepository(db *mongo.Database) uc.ShortUrlRepository {
	return &shortUrlRepository{db: db}
}

func (r *shortUrlRepository) Create(ctx context.Context, shortUrl *domain.ShortUrl) (*domain.ShortUrl, error) {
	if err := r.createTtlIndex(ctx); err != nil {
		return nil, err
	}
	_, err := r.db.Collection(collection).InsertOne(ctx, shortUrl)
	if err != nil {
		return nil, err
	}
	return shortUrl, nil
}

func (r *shortUrlRepository) GetByKey(ctx context.Context, key string) (*domain.ShortUrl, error) {
	var record domain.ShortUrl
	filter := bson.D{{"key", key}}
	err := r.db.Collection(collection).FindOne(ctx, filter).Decode(&record)
	return &record, err
}

func (r *shortUrlRepository) createTtlIndex(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"createdat": 1},
		Options: options.Index().SetExpireAfterSeconds(urlTtl),
	}
	_, err := r.db.Collection(collection).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
