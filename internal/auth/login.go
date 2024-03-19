package auth

import (
	"encoding/json"
	"net/http"

	"github.com/devhindo/bats/types"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u  types.User

	json.NewDecoder(r.Body).Decode(&u)

	if u.Username == "admin" && u.Password == "admin" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
		return
	}
}
