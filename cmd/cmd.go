package cmd

import (
	"bookkeeping-server/database"
	"bookkeeping-server/internal/router"
)

func RunServer() {
	database.Connect()
	r := router.New()
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
