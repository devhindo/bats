package api

/*

	user endpoints

	/user/new
	/user/login
	/user/ban     yes?no?idk.

*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/devhindo/bats/internal/env"
)

var (
	secretKey = []byte(env.GetEnv("JWT_SECRET_KEY"))
)

func (s *APIServer) handleNewUser(w http.ResponseWriter, r *http.Request) {
	// check if user is already registered
	if r.Method == "POST" {
		// parse user info into User struct
	}
}

func (s *APIServer) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	log.Printf("The user request value %v", u)

	if true /* check username in db */ {
		tokenString, err := createToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
			fmt.Errorf("error creating token: %v", err)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "invalid credentials")
		}
	}

		
}

func (s *APIServer) handleNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// parse user info into User struct
	}
}

func createToken(u *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"email":    u.Email,
		"exp": 	time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}