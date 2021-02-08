package dao

import (
	"net/http"

	"go-echo-poc/app/datasources/cassandra/users_db"
	"go-echo-poc/app/model"
	"go-echo-poc/app/utils"

	"github.com/gocql/gocql"
)

var (
	UsersDaoService usersDaoInterface = &usersDao{}
)

const (
	queryInsertUser = "INSERT INTO app_users (uid, first_name, last_name, email, status, date_created) VALUES (?, ?, ?, ?, ?, ?)"
	queryGetUser    = "SELECT first_name, last_name, email, date_created, status FROM app_users WHERE uid=?;"
	queryUpdateUser = "UPDATE app_users SET first_name=?, last_name=?, email=? WHERE uid=?"
	queryDeleteUser = "DELETE FROM app_users WHERE uid=?"
)

type usersDao struct{}

func (usersDao) Get(user *model.User) error {
	if err := users_db.GetSession().Query(queryGetUser, user.Uuid).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
	); err != nil {
		if err == gocql.ErrNotFound {
			return utils.NewHTTPError(http.StatusNotFound, "error when trying to get current UUID", "User not found")
		}
		return utils.NewHTTPError(http.StatusInternalServerError, "error when trying to get current UUID", err.Error())
	}
	return nil
}

func (usersDao) Save(user *model.User) (*gocql.UUID, error) {
	user.Uuid = gocql.TimeUUID()

	if err := users_db.GetSession().Query(queryInsertUser,
		user.Uuid,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Status,
		user.DateCreated,
	).Exec(); err != nil {
		return nil, utils.NewHTTPError(http.StatusInternalServerError, "error when tying to save user", err.Error())
	}
	return &user.Uuid, nil
}

func (usersDao) Update(user *model.User) error {
	if err := users_db.GetSession().Query(queryUpdateUser,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Uuid,
	).Exec(); err != nil {
		return utils.NewHTTPError(http.StatusInternalServerError, "error when trying to update UUID", err.Error())
	}
	return nil
}

func (usersDao) Delete(user *model.User) error {
	if err := users_db.GetSession().Query(queryDeleteUser,
		user.Uuid,
	).Exec(); err != nil {
		return utils.NewHTTPError(http.StatusInternalServerError, "error when trying to Delete UUID", err.Error())
	}
	return nil
}
