package dao

import (
	"time"
)

type logDaoInterface interface {
	Save(*Log) error
}

var (
	LogDaoService logDaoInterface = &Log{}
)

type Log struct {
	Time       time.Time `json:"time"`
	Uuid       string    `json:"uuid"`
	Action     string    `json:"action"`
	StatusCode string    `json:"statusCode"`
	Message    string    `json:"message"`
}
