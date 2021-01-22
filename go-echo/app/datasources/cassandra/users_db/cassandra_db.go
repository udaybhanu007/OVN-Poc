package users_db

import (
	"go-echo-poc/config"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func ConnectDB(configuration config.Configuration) {

	cluster := gocql.NewCluster(configuration.DbConfig.CassandraHost)

	cluster.Keyspace = configuration.DbConfig.CassandraKeyspace
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
