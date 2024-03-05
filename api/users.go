package api

import (
	"fmt"
	"net/http"
)

func handleNewUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add users")
	w.Write([]byte("Add users"))
}