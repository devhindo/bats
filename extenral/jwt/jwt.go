package jwt

// implement register user

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type Params struct {
	key []byte
	t *jwt.Token
	s string
}



func createToken(u User, s []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"username": u.Username,
			"exp": time.Now().Add(time.Hour * 4).Unix(),
		})
	
	tokenString, err := token.SignedString(s)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, secret []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func RegisterUser(u User, key []byte) {
	
//	k := key 
//	t := jwt.New(jwt.SigningMethodHS256)
//	s := t.SignedString(key)
//
}