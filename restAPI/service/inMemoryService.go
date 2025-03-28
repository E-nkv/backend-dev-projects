package service

import (
	"fmt"
	"time"

	"github.com/E-nkv/backend-dev-projects/restAPI/errs"
	"github.com/E-nkv/backend-dev-projects/restAPI/types"
)

type InMemoryService struct {
	Users   []types.User
	lastID  int64 //simulate a sequence
	Service       //this makes the InMemoryService implement the Service interface. careful here: if we call a non-implemented method, the code panics
}

func (ims *InMemoryService) GetUsers() ([]types.User, error) {
	if time.Now().Unix()%3 == 0 {
		return nil, fmt.Errorf("casually cannot work when now is divisible by 3")
	}
	return ims.Users, nil
}
func (ims *InMemoryService) CreateUser(user *types.UserCreate) (int64, error) {
	if time.Now().Unix()%3 == 0 {
		return -1, fmt.Errorf("casually cannot work when now is divisible by 3")
	}
	userToAdd := types.User{ID: ims.lastID, Email: user.Email}
	ims.lastID++
	ims.Users = append(ims.Users, userToAdd)
	return userToAdd.ID, nil
}
func (ims *InMemoryService) GetUser(ID int64) (types.User, error) {
	for _, u := range ims.Users {
		if u.ID == ID {
			return u, nil
		}
	}
	return types.User{}, errs.ErrNotFound
}
func (ims *InMemoryService) DeleteUser(id int64) error {
	if time.Now().Unix()%3 == 0 {
		return fmt.Errorf("casually cannot work when now is divisible by 3")
	}
	for i, u := range ims.Users {
		if u.ID == id {
			ims.Users = append(ims.Users[:i], ims.Users[i+1:]...)
			return nil
		}
	}
	return errs.ErrNotFound
}
