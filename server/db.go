package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	
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
	_, _ = dot.Exec(db.conn, "posts")

	log.Println("db: tables created")
}

func (db *DB) checkExsistance(table string, column string, value string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, column)
	var count int
	err := db.conn.QueryRow(query, value).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error: db: error in querying user: err: %v", err)
	}	
	return count > 0, nil
}

func (db *DB) addRecord(tablename string, columns []string, values []string) error {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tablename, strings.Join(columns, ", "), "'"+strings.Join(values, "', '")+"'")
	log.Println(query)
	_, err := db.conn.Exec(query)
	if err != nil {
		return fmt.Errorf("error: db: error in inserting record: err: %v", err)
	}


	return nil
}