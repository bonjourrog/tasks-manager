package service

import (
	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListService interface {
	Create(list entity.List) entity.MongoResult
	FetchAll(user_id string) ([]entity.List, error)
	Delete(list_id primitive.ObjectID) entity.MongoResult
	Update(list_id primitive.ObjectID, list entity.List) entity.MongoResult
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
func (*listService) Delete(list_id primitive.ObjectID) entity.MongoResult {
	return _listRepo.Delete(list_id)
}
func (*listService) Update(list_id primitive.ObjectID, list entity.List) entity.MongoResult {
	return _listRepo.Update(list_id, list)
}
