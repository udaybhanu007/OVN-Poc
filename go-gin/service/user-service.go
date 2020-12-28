package service

import (
	"fmt"
	"golang-gin-poc/entity"
)

// UserService interface
type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	DeleteRecord(int) []entity.User
	UpdateRecord(entity.User) []entity.User
}

type userService struct {
	users []entity.User
}

//New func
func New() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user
}
func (service *userService) FindAll() []entity.User {
	return service.users
}

func (service *userService) DeleteRecord(id int) []entity.User {
	for idx, item := range service.users {
		if item.ID == id {
			service.users = append(service.users[0:idx], service.users[idx+1:]...)
		}
	}
	return service.users
}

func (service *userService) UpdateRecord(user entity.User) []entity.User {
	fmt.Println(user)
	for idx, item := range service.users {
		if item.ID == user.ID {
			service.users[idx] = user
			// service.users[idx].FirstName = user.FirstName
			// service.users[idx].LastName = user.LastName
			// service.users[idx].Age = user.Age
			// service.users[idx].LastName = user.LastName
			// service.users[idx].UserAddress.Street = user.UserAddress.Street
			// service.users[idx].UserAddress.City = user.UserAddress.City
			// service.users[idx].UserAddress.Zipcode = user.UserAddress.Zipcode
		}
	}
	return service.users
}
