package api

import (
	"net/http"
	"encoding/json"
	"fmt"
)


type APIServer struct{
	listenAddress string
}

func (s* APIServer) RUN() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleBase)
	mux.HandleFunc("GET /api/", s.handleAPIBaseRoute)
	mux.HandleFunc("POST /users/register/", s.handleRegisterUser)
	//mux.HandleFunc("POST /users/login", handleAddUsers)

	http.ListenAndServe(":8080", mux)
}

func (s* APIServer) handleBase(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "heloo",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

}

func (s* APIServer) handleAPIBaseRoute(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"message": "heloo",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

	//w.Write([]byte("API Base Route"))
}

func (s* APIServer) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add users")
	w.Write([]byte("Add users"))
}