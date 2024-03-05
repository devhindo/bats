package api

import (
	"net/http"
)

func handleAPIBaseRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Base Route"))
}