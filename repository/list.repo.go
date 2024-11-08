package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/bonjourrog/taskm/db"
	"github.com/bonjourrog/taskm/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListRepo interface {
	Create(list entity.List) entity.MongoResult
	FetchAll(user_id string) ([]entity.List, error)
}

type list struct{}

func NewListRepository() ListRepo {
	return &list{}
}
func (*list) Create(list entity.List) entity.MongoResult {
	var (
		response entity.MongoResult
		_db      = db.NewMongoConnection()
	)

	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("list")
	result, err := coll.InsertOne(context.TODO(), list)
	if err != nil {
		response = entity.MongoResult{
			Success:    false,
			Message:    err.Error(),
			InsertedID: "",
		}
		return response
	}
	response = entity.MongoResult{
		Success:    true,
		Message:    "List created successfuly",
		InsertedID: fmt.Sprintf("%v", result.InsertedID),
		Count:      0,
		Data:       nil,
	}
	return response
}
func (*list) FetchAll(user_id string) ([]entity.List, error) {
	var (
		_db   = db.NewMongoConnection()
		_list []entity.List
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("list")
	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"user_id", user_id}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "tasks"},
			{"localField", "_id"},
			{"foreignField", "list_id"},
			{"as", "tasks"},
		}}},
	}
	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &_list); err != nil {
		return nil, err
	}
	return _list, nil
}
