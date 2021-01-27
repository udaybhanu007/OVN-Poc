package domain

import (
	"go-echo-poc/app/helpers"
	"net/http"
	"strings"

	"github.com/gocql/gocql"
)

type usersDaoInterface interface {
	Get(*User) error
	Save(*User) (*gocql.UUID, error)
	Update(*User) error
	Delete(*User) error
}

var (
	UsersDaoService usersDaoInterface = &User{}
)

const (
	StatusActive = "active"
)

type User struct {
	Uuid        gocql.UUID `json:"uuid"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	DateCreated string     `json:"date_created"`
	Status      string     `json:"status"`
}

func (user *User) Validate() error {
	user.FirstName = strings.TrimSpace(user.FirstName)
	if user.FirstName == "" {
		return helpers.NewHTTPError(http.StatusBadRequest, "INVALID", "invalid first name")
	}
	user.LastName = strings.TrimSpace(user.LastName)
	if user.LastName == "" {
		return helpers.NewHTTPError(http.StatusBadRequest, "INVALID", "invalid last name")
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return helpers.NewHTTPError(http.StatusBadRequest, "INVALID", "invalid email address")
	}
	return nil
}
