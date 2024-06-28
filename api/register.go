package api

import (
	"net/http"

	"github.com/devhindo/bats/pkg/jwt"
)
// endpoint for registering new users using JWT authentication

func (s *APIServer) handleRegisterUser(w http.ResponseWriter, r *http.Request) {

}

var secretKey = []byte(os.env)