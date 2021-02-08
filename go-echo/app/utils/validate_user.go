package utils

import (
	"go-echo-poc/app/model"
	"net/http"
	"strings"
)

func Validate(user model.User) error {
	user.FirstName = strings.TrimSpace(user.FirstName)
	if user.FirstName == "" {
		return NewHTTPError(http.StatusBadRequest, "INVALID", "invalid first name")
	}
	user.LastName = strings.TrimSpace(user.LastName)
	if user.LastName == "" {
		return NewHTTPError(http.StatusBadRequest, "INVALID", "invalid last name")
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return NewHTTPError(http.StatusBadRequest, "INVALID", "invalid email address")
	}
	return nil
}
