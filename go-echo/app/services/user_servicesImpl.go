package services

import (
	"fmt"
	"go-echo-poc/app/dao"
	"go-echo-poc/app/model"
	"go-echo-poc/app/utils"

	"github.com/gocql/gocql"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

const (
	StatusActive = "active"
)

type usersService struct{}

func (s *usersService) CreateUser(user model.User) (*utils.NotificationMessage, error) {
	if err := utils.Validate(user); err != nil {
		return nil, err
	}
	user.Status = StatusActive
	user.DateCreated = utils.GetNowDBFormat()

	uuid, err := dao.UsersDaoService.Save(&user)
	if err != nil {
		return nil, err
	}
	notification := utils.NotificationMessage{
		Message: fmt.Sprintf("user %v created", uuid),
	}
	return &notification, nil

}

func (s *usersService) GetUser(UUID gocql.UUID) (*model.User, error) {
	userObj := &model.User{Uuid: UUID}
	if err := dao.UsersDaoService.Get(userObj); err != nil {
		return nil, err
	}

	return userObj, nil
}

func (s *usersService) UpdateUser(user model.User) (*utils.NotificationMessage, error) {
	currentUser := &model.User{Uuid: user.Uuid}
	if err := dao.UsersDaoService.Get(currentUser); err != nil {
		return nil, err
	}

	if err := utils.Validate(user); err != nil {
		return nil, err
	}

	if err := dao.UsersDaoService.Update(&user); err != nil {
		return nil, err
	}
	notification := utils.NotificationMessage{
		Message: fmt.Sprintf("user %v updated", user.Uuid),
	}
	return &notification, nil
}

func (s *usersService) DeleteUser(UUID gocql.UUID) (*utils.NotificationMessage, error) {
	user := &model.User{Uuid: UUID}
	if err := dao.UsersDaoService.Get(user); err != nil {
		return nil, err
	}
	fmt.Println("delete start")
	if err := dao.UsersDaoService.Delete(user); err != nil {
		return nil, err
	}
	notification := utils.NotificationMessage{
		Message: fmt.Sprintf("user %v deleted", UUID),
	}
	return &notification, nil
}
