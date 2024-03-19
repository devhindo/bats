package app

import (
	"github.com/devhindo/bats/pkg/db"
	"github.com/devhindo/bats/internal/api"
)

const (
	Version = "0.0.1"
	Author = "devhindo"
)

func RUN() {

	db := db.New("mongo")
	db.Connect()
	api.INIT(db)
}