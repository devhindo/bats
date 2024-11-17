package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type API struct {
	db *DB
}

func (api *API) runAPI() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /signup", api.handleSignUp)
	mux.HandleFunc("POST /signup/otp", api.handleOTP)
	mux.HandleFunc("POST /home", api.handleHome)

	/*
	mux.HandleFunc("GET /api/", s.handleAPIBaseRoute)
	mux.HandleFunc("/auth/login", s.handleLogin)
	mux.HandleFunc("/auth/refresh", s.handleRefreshToken)
	
	*/

	handler := corsMiddleware(mux)

	port := ":8080"
	log.Println("Server is running on port" + port)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling root")
	io.WriteString(w, "Hello Bats!")
}

func  handleBase(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "heli",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

}

func handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Refresh Token")
	w.Write([]byte("Refresh Token"))
}

func  handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Login")
		w.Write([]byte("Login"))
		return
	}
	fmt.Fprintf(w, "Login")
	w.Write([]byte("Login"))
}


func handleAPIBaseRoute(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"message": "heloo",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

	//w.Write([]byte("API Base Route"))
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight requests
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}