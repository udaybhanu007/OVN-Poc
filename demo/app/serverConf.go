package app

import (
	"net/http"
)

func ConfigureAndStartServer() {
	server := http.Server{
		Addr:    "0.0.0.0:4200",
		Handler: nil,
	}
	server.ListenAndServe()
}
