package main

import (
	"log"

	database "github.com/nedssoft/go-basic-api/bin/db"
	"github.com/nedssoft/go-basic-api/cmd/api"
	"github.com/nedssoft/go-basic-api/config"
)

func main() {
	db, err := database.NewDB(config.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":3000", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
