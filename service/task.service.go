package service

import (
	"github.com/bonjourrog/taskm/entity"
	taskrepo "github.com/bonjourrog/taskm/repository/taskRepo"
)

type TaskService interface {
	Create(task entity.Task) entity.MongoResult
}

type taskSerive struct{}

var (
	_taskRepo taskrepo.Task
)

func NewTaskService(taskRepo taskrepo.Task) TaskService {
	_taskRepo = taskRepo
	return &taskSerive{}
}

func (*taskSerive) Create(task entity.Task) entity.MongoResult {
	return _taskRepo.Create(task)
}
