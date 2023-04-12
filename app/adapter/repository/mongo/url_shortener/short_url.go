package url_shortener

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "github.com/3FanYu/url-shortener-go/domain/url_shortener"
	uc "github.com/3FanYu/url-shortener-go/usecase/url_shortener"
)

type shortUrl struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Key       string             `bson:"key"`
	Url       string             `bson:"url"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type shortUrlRepository struct {
	db *mongo.Database
}

const collection = "short_urls"
const urlTtl = 10

func NewRepository(db *mongo.Database) uc.ShortUrlRepository {
	return &shortUrlRepository{db: db}
}

func (r *shortUrl) toDomain() *domain.ShortUrl {
	return &domain.ShortUrl{
		ID:        r.ID,
		Key:       r.Key,
		Url:       r.Url,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func toRepository(shortUrlDomain *domain.ShortUrl) *shortUrl {
	return &shortUrl{
		ID:        shortUrlDomain.ID,
		Key:       shortUrlDomain.Key,
		Url:       shortUrlDomain.Url,
		CreatedAt: shortUrlDomain.CreatedAt,
		UpdatedAt: shortUrlDomain.UpdatedAt,
	}
}

func (r *shortUrlRepository) FindOrCreate(ctx context.Context, shortUrlDomain *domain.ShortUrl) (*domain.ShortUrl, error) {
	var record shortUrl
	filter := bson.D{{Key: "url", Value: shortUrlDomain.Url}}
	err := r.db.Collection(collection).FindOne(ctx, filter).Decode(&record)
	if err == nil {
		return record.toDomain(), nil
	}
	if err := r.createTtlIndex(ctx); err != nil {
		return nil, err
	}

	insertRecord := toRepository(shortUrlDomain)
	_, err = r.db.Collection(collection).InsertOne(ctx, insertRecord)
	if err != nil {
		return nil, err
	}
	return insertRecord.toDomain(), nil
}

func (r *shortUrlRepository) GetByKey(ctx context.Context, key string) (*domain.ShortUrl, error) {
	var record shortUrl
	filter := bson.D{{Key: "key", Value: key}}
	err := r.db.Collection(collection).FindOne(ctx, filter).Decode(&record)
	return record.toDomain(), err
}

func (r *shortUrlRepository) createTtlIndex(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"created_at": 1},
		Options: options.Index().SetExpireAfterSeconds(urlTtl),
	}
	_, err := r.db.Collection(collection).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
