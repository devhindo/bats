package main

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	logger *customLogger
)

func main() {

	log.SetFlags(log.Lshortfile)

	logger = newCustomLogger()
    logger.Println("This is a log message with a colored filename.")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
	
	db := &DB{}
	db.init()
	db.createTables()
	defer db.conn.Close()
		
	api := &API{db: db}
	api.runAPI()
}
