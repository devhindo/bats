package api

import (
	"net/http"
)

func RUN() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", handleAPIBaseRoute)
	mux.HandleFunc("POST /users/new", handleNewUsers)
	mux.HandleFunc("POST /users/login", handleAddUsers)

	http.ListenAndServe(":8080", mux)
}