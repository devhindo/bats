package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddress string
}

func (s *APIServer) RUN() {

	mux := http.NewServeMux()

	// user authentication
	mux.HandleFunc("POST /user/create", s.handleCreateUser)

	mux.HandleFunc("/", s.handleBase)
	mux.HandleFunc("GET /api/", s.handleAPIBaseRoute)
	mux.HandleFunc("/auth/login", s.handleLogin)
	mux.HandleFunc("/auth/refresh", s.handleRefreshToken)
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func (s *APIServer) handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Refresh Token")
	w.Write([]byte("Refresh Token"))
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Login")
		w.Write([]byte("Login"))
		return
	}
	fmt.Fprintf(w, "Login")
	w.Write([]byte("Login"))
}

func (s *APIServer) handleBase(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "heloo",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

}

func (s *APIServer) handleAPIBaseRoute(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"message": "heloo",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

	//w.Write([]byte("API Base Route"))
}
