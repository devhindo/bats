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
	// "github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {

	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		log.Println("error: /signup: invalid credentails: err: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(cred)
	 log.Println("Handling signup")
}