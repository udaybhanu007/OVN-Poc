package app

import (
	"demo/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe("0.0.0.0:4200", nil); err != nil {
		panic(err)
	}
}
