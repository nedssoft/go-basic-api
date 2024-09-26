package main

import (
	"log"

	database "github.com/nedssoft/learn-go/bin/db"
	"github.com/nedssoft/learn-go/cmd/api"
	"github.com/nedssoft/learn-go/config"
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