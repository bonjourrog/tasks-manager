package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bonjourrog/taskm/db"
	"github.com/bonjourrog/taskm/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListRepo interface {
	Create(list entity.List) entity.MongoResult
	FetchAll(user_id string) ([]entity.List, error)
	Delete(list_id primitive.ObjectID) entity.MongoResult
	Update(list_id primitive.ObjectID, list entity.List) entity.MongoResult
}

type list struct{}

func NewListRepository() ListRepo {
	return &list{}
}
func (*list) Create(list entity.List) entity.MongoResult {
	var (
		response entity.MongoResult
		lists    []entity.List
		_db      = db.NewMongoConnection()
	)

	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("list")
	cursor, err := coll.Find(context.TODO(), bson.M{"user_id": list.UserID})
	if err != nil {
		response = entity.MongoResult{
			Success: false,
			Message: err.Error(),
		}
		return response
	}
	if err = cursor.All(context.TODO(), &lists); err != nil {
		response = entity.MongoResult{
			Success: false,
			Message: err.Error(),
		}
		return response
	}
	if len(lists) == 5 {
		response = entity.MongoResult{
			Success: false,
			Message: "have reached the limit for creating new lists",
			Data:    lists,
		}
		return response
	}
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
func (*list) Delete(list_id primitive.ObjectID) entity.MongoResult {
	var (
		_db      = db.NewMongoConnection()
		response entity.MongoResult
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("list")
	deleteResult, err := coll.DeleteOne(context.TODO(), bson.M{"_id": list_id})
	if err != nil {
		response = entity.MongoResult{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
		return response
	}
	response = entity.MongoResult{
		Success: true,
		Message: "list deleted successfully",
		Data:    deleteResult.DeletedCount,
	}
	return response
}
func (*list) Update(list_id primitive.ObjectID, list entity.List) entity.MongoResult {
	var (
		_db     = db.NewMongoConnection()
		respose entity.MongoResult
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("list")
	result, err := coll.UpdateOne(context.TODO(), bson.M{"_id": list_id}, bson.M{"$set": bson.M{
		"name":       list.Name,
		"color":      list.Color,
		"updated_at": time.Now(),
	}})
	if err != nil {
		respose = entity.MongoResult{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
		return respose
	}
	respose = entity.MongoResult{
		Success: true,
		Message: "list updated successfully",
		Data:    result.UpsertedID,
	}
	return respose
}
