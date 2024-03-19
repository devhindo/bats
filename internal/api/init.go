package api

import "github.com/devhindo/bats/pkg/db"

func INIT(db *db.DB) {
	api := APIServer{
		listenAddress: ":8080",
	}
	api.RUN()
}