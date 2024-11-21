package main

import (
	"log"
	"net/http"
)

func (api *API) handleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("recieved a /home request from: ", r.RemoteAddr)
	w.Write([]byte("Welcome to Home"))
}