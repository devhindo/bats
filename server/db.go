package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qustavo/dotsql"
)


type DB struct {
	conn *sql.DB
}

func (db *DB) init() {
	
	dsn := "root:root@tcp(127.0.0.1:3306)/bats"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("error: db: error in opening connection: err: ", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("error: db: error in pinging connection: err: ", err)
	}

	db.conn = conn
	log.Println("db: connection established")
}

func (db *DB) createTables() {

	dot, err := dotsql.LoadFromFile("../scripts/db/schema.sql")
	if err != nil {
		log.Fatal("error: db: error in loading schema: err: ", err)
	}

	res, err := dot.Exec(db.conn, "db")
	if err != nil {
		log.Println("error: db: error in creating db: err: ", err)
	}
	log.Printf("db: db created: %v", res)

	res, err = dot.Exec(db.conn, "use")
	if err != nil {
		log.Fatal("error: db: error in using db: err: ", err)
	}
	log.Printf("db: using db: %v", res)
	
	_, _ = dot.Exec(db.conn, "users")
	_, _ = dot.Exec(db.conn, "addtestusers")

	log.Println("db: tables created")
}

func (db *DB)  userExists(username string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	var count int
	err := db.conn.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error: db: error in querying user: err: %v", err)
	}	
	return count > 0, nil
}
