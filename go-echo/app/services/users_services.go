package services

import (
	"go-echo-poc/app/model"
	"go-echo-poc/app/utils"

	"github.com/gocql/gocql"
)

type usersServiceInterface interface {
	CreateUser(model.User) (*utils.NotificationMessage, error)
	GetUser(gocql.UUID) (*model.User, error)
	UpdateUser(model.User) (*utils.NotificationMessage, error)
	DeleteUser(gocql.UUID) (*utils.NotificationMessage, error)
}
