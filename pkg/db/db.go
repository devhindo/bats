package db

import (

)

type Database interface {
	Connect() error
}

type Credentials interface {}

type DB struct {
	database Database
}
/*
func New(name string, credentials Credentials) *DB {
	var database Database
	if name == "mongo" {
		database = new()
	}
}

func (db *DB) Connect() {
	if db.name == "mongo" {
		connectMongo()
	}
	if db.name == "supabase" {
		connectSupabase()
	}
}
*/