package service

import (
	"fmt"
	"golang-gin-poc/entity"
	"golang-gin-poc/helpers"
	"net/http"
)

// UserService interface
type UserService interface {
	Save(entity.User) (*helpers.Notification, *helpers.ApplicationError)
	FindAll() (*[]entity.User, *helpers.ApplicationError)
	FindUser(int) (*entity.User, *helpers.ApplicationError)
	DeleteRecord(int) (*helpers.Notification, *helpers.ApplicationError)
	UpdateRecord(entity.User) (*helpers.Notification, *helpers.ApplicationError)
}

type userService struct {
	users []entity.User
}

//New func
func New() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) (*helpers.Notification, *helpers.ApplicationError) {
	service.users = append(service.users, user)
	return &helpers.Notification{
		Message: fmt.Sprintf("UserID %d saved successfully", user.ID),
	}, nil

}
func (service *userService) FindAll() (*[]entity.User, *helpers.ApplicationError) {
	if len(service.users) > 0 {
		return &service.users, nil
	}
	return nil, &helpers.ApplicationError{
		Message:    fmt.Sprintf("No records found"),
		StatusCode: http.StatusOK,
	}

}

func (service *userService) FindUser(id int) (*entity.User, *helpers.ApplicationError) {
	for idx := range service.users {
		if service.users[idx].ID == id {
			return &service.users[idx], nil
		}
	}
	return nil, &helpers.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", id),
		StatusCode: http.StatusNotFound,
	}
}

func (service *userService) DeleteRecord(id int) (*helpers.Notification, *helpers.ApplicationError) {
	var isDeletedFlag bool = false
	for idx, item := range service.users {
		if item.ID == id {
			service.users = append(service.users[0:idx], service.users[idx+1:]...)
			isDeletedFlag = true
		}
	}
	if isDeletedFlag {
		return &helpers.Notification{
			Message: fmt.Sprintf("UserID %d deleted successfully", id),
		}, nil
	}
	return nil, &helpers.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", id),
		StatusCode: http.StatusNotFound,
	}
}

func (service *userService) UpdateRecord(user entity.User) (*helpers.Notification, *helpers.ApplicationError) {
	for idx, item := range service.users {
		if item.ID == user.ID {
			service.users[idx] = user
			return &helpers.Notification{
				Message: fmt.Sprintf("UserID %d updated successfully", user.ID),
			}, nil
		}
	}
	return nil, &helpers.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", user.ID),
		StatusCode: http.StatusNotFound,
	}
}
