package main

import (
	//	"github.com/devhindo/bats/internal/app"
	"fmt"

	"github.com/devhindo/bats/internal/env"
)

func main() {
	//app.RUN()
	fmt.Println(env.GetEnv("AHMED"))
}