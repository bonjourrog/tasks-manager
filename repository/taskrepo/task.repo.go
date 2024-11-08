package taskrepo

import (
	"context"
	"fmt"
	"os"

	"github.com/bonjourrog/taskm/db"
	"github.com/bonjourrog/taskm/entity"
)

type Task interface {
	Create(task entity.Task) entity.MongoResult
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
