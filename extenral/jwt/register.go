package jwt

// implement register user

import (
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

func RegisterUser(u User, key []byte) {
	
	k := key 
	t := jwt.New(jwt.SigningMethodHS256)
	s := t.SignedString(key)

}