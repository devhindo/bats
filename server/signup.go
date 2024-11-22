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
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	tmpRegisteredUsers = map[string]User{}
)



func (api *API) handleSignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling signupppppppp")
	

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }


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

	// save the token in http-only cookie
	http.SetCookie(w, &http.Cookie{
		Name: "jwt_token",
		Value: user.JWT,
		HttpOnly: true,
//		Secure: true, // set to true if using https
		Path: "/",
	})

	log.Println("jwt_token: " , user.JWT)
	// send the token back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "otp sent to email"})


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
	/*
	var tmpUser User
	tmpUser = User{
		Email: res.Email,
		Username: "tmp",
		Password: "tmp",
	}
	err = addUserToDB(api.db, &tmpUser) // remove later
	*/
	
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
	
	err = addUserToDB(api.db, &user)

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
		return nil, fmt.Errorf("error in parsing token: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
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

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "signed out successfully"})
}

func addUserToDB(db *DB, u *User) error {
	err := db.addRecord("users", []string{"username", "email", "password"}, []string{u.Username, u.Email, u.Password})
	if err != nil {
		return fmt.Errorf("error in adding user to database: %w", err)
	}
	return nil
}