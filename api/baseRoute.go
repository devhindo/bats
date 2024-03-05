package api

import (
	"net/http"
	"encoding/json"
)

func handleAPIBaseRoute(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"message": "heloo",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

	//w.Write([]byte("API Base Route"))
}