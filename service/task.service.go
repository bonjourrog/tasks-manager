package service

import (
	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/repository/taskrepo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService interface {
	Create(task entity.Task) entity.MongoResult
	GetAll(list_id primitive.ObjectID) entity.MongoResult
	Update(task_id primitive.ObjectID, task entity.Task) entity.MongoResult
	DeleteTask(task_id primitive.ObjectID) entity.MongoResult
}

type taskService struct{}

var (
	_taskRepo taskrepo.Task
)

func NewTaskService(taskRepo taskrepo.Task) TaskService {
	_taskRepo = taskRepo
	return &taskService{}
}

func (*taskService) Create(task entity.Task) entity.MongoResult {
	return _taskRepo.Create(task)
}
func (*taskService) GetAll(list_id primitive.ObjectID) entity.MongoResult {
	return _taskRepo.GetAll(list_id)
}
func (*taskService) Update(task_id primitive.ObjectID, task entity.Task) entity.MongoResult {
	return _taskRepo.UpdateTask(task_id, task)
}
func (*taskService) DeleteTask(task_id primitive.ObjectID) entity.MongoResult {
	return _taskRepo.DeleteTask(task_id)
}
