package services

import (
	"fmt"
	"go-echo-poc/app/domain"
	"go-echo-poc/app/helpers"

	"github.com/gocql/gocql"
)

type usersServiceInterface interface {
	CreateUser(domain.User) (*helpers.NotificationMessage, error)
	GetUser(gocql.UUID) (*domain.User, error)
	UpdateUser(domain.User) (*helpers.NotificationMessage, error)
	DeleteUser(gocql.UUID) (*helpers.NotificationMessage, error)
}

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

func (s *usersService) CreateUser(user domain.User) (*helpers.NotificationMessage, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = domain.StatusActive
	user.DateCreated = helpers.GetNowDBFormat()

	uuid, err := domain.UsersDaoService.Save(&user)
	if err != nil {
		return nil, err
	}
	notification := helpers.NotificationMessage{
		Message: fmt.Sprintf("user %v created", uuid),
	}
	return &notification, nil

}

func (s *usersService) GetUser(UUID gocql.UUID) (*domain.User, error) {
	userObj := &domain.User{Uuid: UUID}
	if err := domain.UsersDaoService.Get(userObj); err != nil {
		return nil, err
	}

	return userObj, nil
}

func (s *usersService) UpdateUser(user domain.User) (*helpers.NotificationMessage, error) {
	currentUser := &domain.User{Uuid: user.Uuid}
	if err := domain.UsersDaoService.Get(currentUser); err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := domain.UsersDaoService.Update(&user); err != nil {
		return nil, err
	}
	notification := helpers.NotificationMessage{
		Message: fmt.Sprintf("user %v updated", user.Uuid),
	}
	return &notification, nil
}

func (s *usersService) DeleteUser(UUID gocql.UUID) (*helpers.NotificationMessage, error) {
	user := &domain.User{Uuid: UUID}
	if err := domain.UsersDaoService.Get(user); err != nil {
		return nil, err
	}
	fmt.Println("delete start")
	if err := domain.UsersDaoService.Delete(user); err != nil {
		return nil, err
	}
	notification := helpers.NotificationMessage{
		Message: fmt.Sprintf("user %v deleted", UUID),
	}
	return &notification, nil
}
