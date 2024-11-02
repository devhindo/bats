package main 

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return	json.NewEncoder(w).Encode(data)
}


// ip2location https://api.ip2location.io/?key={apiKey}&ip={ipAddress}
