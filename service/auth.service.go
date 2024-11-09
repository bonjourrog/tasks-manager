package service

import (
	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/repository/authrepo"
)

type AuthService interface {
	Register(user entity.User) entity.MongoResult
}

type authService struct{}

var (
	_authRepo authrepo.AuthRepo
)

func NewAuthService(authRepo authrepo.AuthRepo) AuthService {
	_authRepo = authRepo
	return &authService{}
}
func (*authService) Register(user entity.User) entity.MongoResult {
	return _authRepo.Register(user)
}
