package api

import (
	"fmt"
	"net/http"
)

func handleLoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login User")
	w.Write([]byte("Login User"))
}