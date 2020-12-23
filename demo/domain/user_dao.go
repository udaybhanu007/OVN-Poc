package domain

import (
	"demo/helpers"
	"fmt"
	"net/http"
)

var (
	users = map[int64]*User{
		1: {ID: 1, FirstName: "test first name", LastName: "test Last name", Email: "testEmail@abc.com"},
		2: {ID: 2, FirstName: "second first name", LastName: "second Last name", Email: "secondEmail@abc.com"},
		3: {ID: 3, FirstName: "third first name", LastName: "third Last name", Email: "thirdEmail@abc.com"},
	}
)

func GetUser(userId int64) (*User, *helpers.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &helpers.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
