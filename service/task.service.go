package service

import (
	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/repository/taskrepo"
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
