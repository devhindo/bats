package main

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
	"os"
	"strings"
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

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: user.JWT,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		//Expires: time.Now().Add(time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})

	log.Printf("Handling otp: %+v" , res)


}

// Function to verify JWT tokens
func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		log.Println("authHeader: ", authHeader)	

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

func (api *API) handleSignOut(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling signout")

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "",
		HttpOnly: true,
		Path: "/",
		MaxAge: -1,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signed out successfully"))
}