package app

import (
	"demo/auth"
	"demo/controllers"
	"demo/helpers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/", auth.HandleMain)
	http.HandleFunc("/login", auth.HandleGitHubLogin)
	http.HandleFunc("/users", controllers.GetUser)
	http.Handle("/adduser", helpers.RootHandler(controllers.AddUser))
	http.HandleFunc("/getToken", auth.GetAuthToken)
	ConfigureAndStartServer()
}
