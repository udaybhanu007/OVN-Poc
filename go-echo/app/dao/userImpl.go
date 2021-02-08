package dao

import (
	"go-echo-poc/app/model"

	"github.com/gocql/gocql"
)

type usersDaoInterface interface {
	Get(*model.User) error
	Save(*model.User) (*gocql.UUID, error)
	Update(*model.User) error
	Delete(*model.User) error
}
