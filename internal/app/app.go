package app

import (

	"github.com/devhindo/bats/internal/api"
	"github.com/devhindo/bats/internal/env"
	"github.com/devhindo/bats/pkg/db"
)

const (
	Version = "0.0.1"
	Author = "devhindo"
)

func RUN() {

	db := db.NewMongoDB(db.MongoDB{
		ConnectionString: env.GetEnv("MONGO_CONNECTION_STRING")})

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	api.INIT(":8080", db)
}