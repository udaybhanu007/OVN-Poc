package app

import (
	"net/http"
)

func ConfigureAndStartServer() {
	server := http.Server{
		Addr:    "localhost:4200", // DO NOT UPDATE
		Handler: nil,
	}
	server.ListenAndServe()
}
