package service

import (
	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/repository"
)

type ListService interface {
	Create(list entity.List) entity.MongoResult
	FetchAll(user_id string) ([]entity.List, error)
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
func (*listService) FetchAll(user_id string) ([]entity.List, error) {
	return _listRepo.FetchAll(user_id)
}
