package main

import (
	"flag"
	"fmt"
	"go-echo-poc/app"
	"go-echo-poc/app/datasources/cassandra/users_db"
	"go-echo-poc/config"
	"os"
)

var profile = ""
var configServer = "http://localhost:8090"

func init() {
	name := flag.String("name", "echo-cassandra-demo", "Application name")
	profileTerminal := flag.String("profile", "", "Configuration profile URL")
	configServerTerminal := flag.String("config", "http://localhost:8090/", "Config server base url")
	flag.Parse()
	if profileTerminal != nil {
		profile = *profileTerminal
	}
	if configServerTerminal != nil {
		configServer = *configServerTerminal
	}
	if len(profile) == 0 {
		fmt.Println("profile flag is empty")
		os.Exit(1)
	}
	config.ApplicationName = *name
	config.ConfigProfile = profile
	config.ConfigServer = configServer
	config, configErr := config.LoadConfiguration()
	if configErr != nil {
		fmt.Println(configErr)
		os.Exit(1)
	}
	users_db.ConnectDB(config)
}

func main() {
	app.StartApplication()

}
