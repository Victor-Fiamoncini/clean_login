package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const connectionUri = "mongodb://mongo"

func GetCollection(collection string) *mongo.Collection {
	return nil
}
