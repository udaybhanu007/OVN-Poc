package model

import "github.com/gocql/gocql"

type User struct {
	Uuid        gocql.UUID `json:"uuid"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	DateCreated string     `json:"date_created"`
	Status      string     `json:"status"`
}
