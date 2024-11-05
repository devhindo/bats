package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


type DB struct {
	conn *sql.DB
}

func (db *DB) init() {
	
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"
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