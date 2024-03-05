package api

import (
	"fmt"
	"net/http"
)

func handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add users")
	w.Write([]byte("Add users"))
}