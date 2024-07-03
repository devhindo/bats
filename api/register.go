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
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwt logic

func createToken(user *User) (string, error) {
	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"iss": "bats",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	
	fmt.Printf("Token claims added: %+v\n", claims)
	
	fmt.Println(os.Getenv("SECRET_KEY"))
	
	tokenString, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "could'nt sign token", err
	}
	
	return tokenString, nil
	
}

func (s *APIServer) handleRegisterUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	log.Printf("The user request value %v", u)

	confirmationCode := generateConfirmationCode()
	confirmationMail := constructConfirmationMail(confirmationCode)


	err := SendMail(u.Email, confirmationMail)
	if err != nil {
		fmt.Println("couldn't send mail", err)
	}

	// save credentials in local db for 15 minutes.
	// if user confirmed using /register/confirm
	// then persist in the external db.


}

func (s *APIServer) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("handle login user")
	w.Header().Set("Content-Type", "application/json")
	
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	log.Printf("The user request value %v", u)
	
	if true { //check username in db  
		
		tokenString, err := createToken(&u)
		if err != nil {
			fmt.Println("can'ttttttttt create token")
			fmt.Printf("error creating token: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "invalid credentials")
		}

}

		


