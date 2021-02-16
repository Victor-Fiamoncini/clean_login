package mocks

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoClient func
func NewMongoClient(url string) (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI(url))
}

// NewMongoDatabase func
func NewMongoDatabase(name string, client *mongo.Client) *mongo.Database {
	return client.Database(name)
}

// NewMongoSession func
func NewMongoSession(client *mongo.Client) (mongo.Session, error) {
	return client.StartSession()
}

// Connect func
func Connect(client *mongo.Client) error {
	return client.Connect(nil)
}

// FindOne func
func FindOne(ctx context.Context, coll mongo.Collection, filter interface{}) *mongo.SingleResult {
	return coll.FindOne(ctx, filter)
}

// InsertOne func
func InsertOne(ctx context.Context, coll mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
	return coll.InsertOne(ctx, document)
}

// Decode func
func Decode(result *mongo.SingleResult, value interface{}) error {
	return result.Decode(value)
}
