package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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
	mux.HandleFunc("/usr/verify", api.verifyUser)
	
	mux.Handle("POST /signout", JWTAuthMiddleware(http.HandlerFunc(api.handleSignOut)))
	mux.Handle("GET /home", JWTAuthMiddleware(http.HandlerFunc(api.handleHome)))
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
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func checkTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			log.Println("No token provided")
			return
		} 
		log.Println("Middleware Token: ", token)
		next.ServeHTTP(w, r)
	})
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		authHeader := r.Header.Get("Authorization")
		log.Println("authHeader: ", authHeader)	
		
		c, err := r.Cookie("jwt_token")
		if err != nil {
			log.Println("error printing cookie jwt")
		}
		log.Println("a7oooooooooo ", c )

		if authHeader == "" {
			log.Println("error: /signup: no authorization header")
			http.Error(w, "no authorization header", http.StatusUnauthorized)
			return
		}

		// Check if the authorization header is in the correct format
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := verifyToken(tokenString)
		if err != nil {
			log.Println("error: /signup: error in verifying token: err: ", err)
			http.Error(w, "error in verifying token. " + "error: " + err.Error(), http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			log.Println("error: /signup: invalid token")
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *API) verifyUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		log.Println("user is not authorizedddddd")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//log.Println("cookie: ", cookie)

	tokenString := cookie.Value
    token, err := verifyToken(tokenString)
	log.Println(err)
    if err != nil || !token.Valid {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    } 

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Println("claims: ", claims)

	username, ok := claims["sub"].(string)
	log.Println("username: ", username)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "authenticated"})
}