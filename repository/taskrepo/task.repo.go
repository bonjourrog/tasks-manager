package taskrepo

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

type Task interface {
	Create(task entity.Task) entity.MongoResult
	GetAll(list_id primitive.ObjectID) entity.MongoResult
	UpdateTask(task_id primitive.ObjectID, task entity.Task) entity.MongoResult
	DeleteTask(task_id primitive.ObjectID) entity.MongoResult
}

type taskRepo struct{}

func NewTasksRepo() Task {
	return &taskRepo{}
}
func (*taskRepo) Create(task entity.Task) entity.MongoResult {
	var (
		_db     = db.NewMongoConnection()
		mResult entity.MongoResult
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("tasks")
	result, err := coll.InsertOne(context.TODO(), task)
	if err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Message: err.Error(),
		}
		return mResult
	}
	mResult = entity.MongoResult{
		Success:    true,
		Message:    "task created successfully",
		InsertedID: fmt.Sprintf("%v", result.InsertedID),
	}
	return mResult
}
func (*taskRepo) GetAll(list_id primitive.ObjectID) entity.MongoResult {
	var (
		_db     = db.NewMongoConnection()
		mResult entity.MongoResult
		tasks   []entity.Task
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv(("MONGO_DB"))).Collection("tasks")
	findResult, err := coll.Find(context.TODO(), bson.M{"list_id": list_id})
	if err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		}
		return mResult
	}
	if err = findResult.All(context.TODO(), &tasks); err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Message: err.Error(),
			Data:    tasks,
		}
		return mResult
	}
	mResult = entity.MongoResult{
		Success: true,
		Message: "Fetching tasks was successful",
		Data:    tasks,
	}
	return mResult

}
func (*taskRepo) UpdateTask(task_id primitive.ObjectID, task entity.Task) entity.MongoResult {
	var (
		_db     = db.NewMongoConnection()
		mResult entity.MongoResult
	)
	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	task.UpdatedAt = time.Now()
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("tasks")
	updateResult, err := coll.UpdateOne(context.TODO(), bson.M{"_id": task_id}, bson.M{"$set": task})
	if err != nil {
		mResult = entity.MongoResult{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
		return mResult
	}
	mResult = entity.MongoResult{
		Success: true,
		Message: "task updated successfully",
		Data:    task_id,
		Count:   updateResult.MatchedCount,
	}
	return mResult

}
func (*taskRepo) DeleteTask(task_id primitive.ObjectID) entity.MongoResult {
	var (
		_db    = db.NewMongoConnection()
		result entity.Task
	)

	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()

	coll := client.Database(os.Getenv("MONGO_DB")).Collection("tasks")

	err := coll.FindOne(context.TODO(), bson.M{"_id": task_id}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.MongoResult{
				Success: false,
				Message: "task not found",
				Data:    nil,
			}
		}
		return entity.MongoResult{
			Success: false,
			Message: fmt.Sprintf("Error fetching task: %s", err.Error()),
			Data:    nil,
		}
	}

	deleteResult, err := coll.DeleteOne(context.TODO(), bson.M{"_id": task_id})

	if err != nil {
		return entity.MongoResult{
			Success: false,
			Message: "error deleting task " + err.Error(),
			Data:    nil,
		}
	}
	if deleteResult.DeletedCount == 0 {
		return entity.MongoResult{
			Success: false,
			Message: "task not found for deleting",
			Data:    nil,
		}
	}

	return entity.MongoResult{
		Success: true,
		Message: "task deleted succefully",
		Data:    deleteResult.DeletedCount,
	}
}
