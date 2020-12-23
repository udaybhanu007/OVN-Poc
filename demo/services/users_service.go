package services

import (
	"demo/domain"
	"demo/helpers"
)

func GetUser(userId int64) (*domain.User, *helpers.ApplicationError) {
	return domain.GetUser(userId)
}
