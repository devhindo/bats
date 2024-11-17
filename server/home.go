package main

import "net/http"

func (api *API) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Home"))
}