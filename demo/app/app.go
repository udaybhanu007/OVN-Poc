package app

import (
	"demo/controllers"
	"demo/helpers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	http.Handle("/adduser", helpers.RootHandler(controllers.AddUser))
	ConfigureAndStartServer()
}
