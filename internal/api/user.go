package api

/*

	user endpoints

	/user/new
	/user/login
	/user/ban     yes?no?idk.

*/


import "net/http"


func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// parse user info into User struct
	}
}