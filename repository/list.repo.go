package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/bonjourrog/taskm/db"
	"github.com/bonjourrog/taskm/entity"
)

type ListRepo interface {
	Create(list entity.List) entity.MongoResult
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
