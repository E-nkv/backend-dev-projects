package service

import "github.com/E-nkv/backend-dev-projects/restAPI/types"

type Service interface {
	GetUsers() ([]types.User, error)
	CreateUser(user *types.UserCreate) (int64, error)
	GetUser(ID int64) (types.User, error)
	DeleteUser(id int64) error
}
