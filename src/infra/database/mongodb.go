package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user     = "root"
	password = "MongoDB2021!"
	url      = "mongodb://localhost:27017"
	database = "auth_clean_architecture"
)

// GetCollection func
func GetCollection(collection string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(url).SetAuth(options.Credential{
		Username: user,
		Password: password,
	}))

	if err != nil {
		panic(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
