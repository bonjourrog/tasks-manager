package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConection interface {
	Connection() *mongo.Client
}

type mongoConnection struct{}

func NewMongoConnection() MongoConection {
	return &mongoConnection{}
}
func (*mongoConnection) Connection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	return client
}
