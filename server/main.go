package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
	
	db := &DB{}
	db.init()
	db.createTables()
	
	runAPI()
}
