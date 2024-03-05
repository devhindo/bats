package api

import (
	"net/http"
	"encoding/json"
)

func RUN() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleBase)
	mux.HandleFunc("GET /api/", handleAPIBaseRoute)
	mux.HandleFunc("POST /users/register/", handleRegisterUser)
	//mux.HandleFunc("POST /users/login", handleAddUsers)

	http.ListenAndServe(":8080", mux)
}


func HandleBase(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "heloo",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

}