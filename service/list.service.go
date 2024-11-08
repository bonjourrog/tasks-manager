package service

import (
	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/repository"
)

type ListService interface {
	Create(list entity.List) entity.MongoResult
}

type listService struct{}

var (
	_listRepo repository.ListRepo
)

func NewListService(listRepo repository.ListRepo) ListService {
	_listRepo = listRepo
	return &listService{}
}
func (*listService) Create(list entity.List) entity.MongoResult {
	return _listRepo.Create(list)
}
