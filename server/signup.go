package main

/*

user endpoints

/user/new
/user/login
/user/ban     yes?no?idk.

*/

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	tmpRegisteredUsers = make(Set)
)

func (api *API) handleSignUp(w http.ResponseWriter, r *http.Request) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("error: /signup: invalid credentails: err: ", err)
		http.Error(w, "error in signup credentials. " + "error: " + err.Error(), http.StatusBadRequest)
		return
	}

	// check if user already exists
	if exists, err := api.db.checkExsistance("users", "username", user.Username); err != nil {
		log.Println("error: /signup: error in checking user exists: err: ", err)
		http.Error(w, "error in checking user exists. " + "error: " + err.Error(), http.StatusInternalServerError)
		return
	} else if exists {
		log.Println("error: /signup: user already exists")
		http.Error(w, "user already exists", http.StatusConflict)
		return
	}

	// check if email already exists
	log.Println(user.Email)
	if exists, err := api.db.checkExsistance("users", "email", user.Email); err != nil {
		log.Println("error: /signup: error in checking email exists: err: ", err)
		http.Error(w, "error in checking email exists. " + "error: " + err.Error(), http.StatusInternalServerError)
		return
	} else if exists {
		log.Println("error: /signup: email already exists")
		http.Error(w, "email already exists", http.StatusConflict)
		return
	}

	// create a new user
	tokenString, err := createJWTToken(&user)
	if err != nil {
		log.Println("error: /signup: error in creating jwt token: err: ", err)
		http.Error(w, "error in creating jwt token. " + "error: " + err.Error(), http.StatusInternalServerError)
		return	
	}

	user.JWT = tokenString

	// generate otp 
	otp := generateRandomString(6)
	err = sendMailResend(user.Email, "Bats: Confirm Registration", "<html><body><p>Your OTP is: <strong>" + otp + "</strong></p><p>Valid for 5 minutes</p></body></html>")
	if err != nil {
		log.Println("error: /signup: error in sending mail: err: ", err)
		http.Error(w, "error in sending mail.", http.StatusInternalServerError)
		return
	}


	// send the token back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})


	log.Println("Handling signup")
}

func createJWTToken(c *User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": c.Username,
		"email": c.Email,
		"password": c.Password,
		"iss": "bats",
		"aud": "bats",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	
	tokenString, err := claims.SignedString(jwtSecret)
	if err != nil {
		log.Println("error: /signup: error in creating jwt token: err: ", err)
		return "", err
	}

	return tokenString, nil
} 