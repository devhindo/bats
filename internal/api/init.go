package api

import "github.com/devhindo/bats/pkg/db"

func INIT(port string, db db.Database) {
	api := APIServer{
		ListenAddress: port,
		DB: db,
	}
	api.RUN()
}