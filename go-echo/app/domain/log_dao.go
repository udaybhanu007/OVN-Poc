package domain

import (
	"fmt"
	"go-echo-poc/app/datasources/cassandra/users_db"
	"go-echo-poc/app/helpers"
	"net/http"
	"time"
)

const (
	insertQuery = "INSERT INTO audit_log (time, uuid, action, statusCode, message) VALUES (?, ?, ?, ?, ?)"
)

func (Log) Save(log *Log) error {

	fmt.Println(log)

	var currentTime = time.Now()
	if err := users_db.GetSession().Query(insertQuery,
		currentTime,
		log.Uuid,
		log.Action,
		log.StatusCode,
		log.Message,
	).Exec(); err != nil {
		return helpers.NewHTTPError(http.StatusInternalServerError, "error when tying to save log", err.Error())
	}
	return nil
}
