package main 

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"time"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return	json.NewEncoder(w).Encode(data)
}


// ip2location https://api.ip2location.io/?key={apiKey}&ip={ipAddress}


func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
    seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(result)
}