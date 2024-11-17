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
	tmpRegisteredUsers = map[string]User{}
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
	user.Otp = otp
	log.Println("otp: ", otp)
	/*
	err = sendMailResend(user.Email, "Bats: Confirm Registration", "<html><body><p>Your OTP is: <strong>" + otp + "</strong></p><p>Valid for 5 minutes</p></body></html>")
	if err != nil {
		log.Println("error: /signup: error in sending mail: err: ", err)
		http.Error(w, "error in sending mail.", http.StatusInternalServerError)
		return
	}
	*/
	/*
	err = sendMailResend(user.Email, "Bats: Confirm Registration", "<html><body><p>Your OTP is: <strong>" + otp + "</strong></p><p>Valid for 5 minutes</p></body></html>")
	if err != nil {
		log.Println("error: /signup: error in sending mail: err: ", err)
		http.Error(w, "error in sending mail.", http.StatusInternalServerError)
		return
	}
	*/
	// add user to temporary registered users
	tmpRegisteredUsers[user.Email] = user

	// send the token back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "otp sent to email"})


	log.Printf("Handling signup: %+v" , user)
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

func (api *API) handleOTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling otp")

	var res struct {
		OTP string `json:"otp"`
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println("error: /signup/otp: invalid credentails: err: ", err)
		http.Error(w, "error in otp credentials. " + "error: " + err.Error(), http.StatusBadRequest)
		return
	}

	// check if user exists in temporary registered users

	user, ok := tmpRegisteredUsers[res.Email]
	if !ok {
		log.Println("error: /signup/otp: user not found in temporary registered users")
		http.Error(w, "user not found in temporary registered users", http.StatusNotFound)
		return
	}

	if user.Otp != res.OTP {
		log.Println("error: /signup/otp: invalid otp" + "invalid otp, entered one is " + res.OTP + " and correct one is " + user.Otp)
		http.Error(w, "invalid otp, entered", http.StatusUnauthorized)
		return
	}

	// add user to database
	err = nil
	if err != nil {
		log.Println("error: /signup/otp: error in adding user to database: err: ", err)
		http.Error(w, "error in adding user to database. " + "error: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// remove user from temporary registered users
	delete(tmpRegisteredUsers, res.Email)

	// send the token back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully", "token": user.JWT})

	log.Printf("Handling otp: %+v" , res)


}


