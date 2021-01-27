package domain

import (
	"net/http"

	"go-echo-poc/app/datasources/cassandra/users_db"
	"go-echo-poc/app/helpers"

	"github.com/gocql/gocql"
)

const (
	queryInsertUser = "INSERT INTO app_users (uid, first_name, last_name, email, status, date_created) VALUES (?, ?, ?, ?, ?, ?)"
	queryGetUser    = "SELECT first_name, last_name, email, date_created, status FROM app_users WHERE uid=?;"
	queryUpdateUser = "UPDATE app_users SET first_name=?, last_name=?, email=? WHERE uid=?"
	queryDeleteUser = "DELETE FROM app_users WHERE uid=?"
)

func (User) Get(user *User) error {
	if err := users_db.GetSession().Query(queryGetUser, user.Uuid).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
	); err != nil {
		if err == gocql.ErrNotFound {
			return helpers.NewHTTPError(http.StatusNotFound, "error when trying to get current UUID", "User not found")
		}
		return helpers.NewHTTPError(http.StatusInternalServerError, "error when trying to get current UUID", err.Error())
	}
	return nil
}

func (User) Save(user *User) (*gocql.UUID, error) {
	user.Uuid = gocql.TimeUUID()

	if err := users_db.GetSession().Query(queryInsertUser,
		user.Uuid,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Status,
		user.DateCreated,
	).Exec(); err != nil {
		return nil, helpers.NewHTTPError(http.StatusInternalServerError, "error when tying to save user", err.Error())
	}
	return &user.Uuid, nil
}

func (User) Update(user *User) error {
	if err := users_db.GetSession().Query(queryUpdateUser,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Uuid,
	).Exec(); err != nil {
		return helpers.NewHTTPError(http.StatusInternalServerError, "error when trying to update UUID", err.Error())
	}
	return nil
}

func (User) Delete(user *User) error {
	if err := users_db.GetSession().Query(queryDeleteUser,
		user.Uuid,
	).Exec(); err != nil {
		return helpers.NewHTTPError(http.StatusInternalServerError, "error when trying to Delete UUID", err.Error())
	}
	return nil
}
