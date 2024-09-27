package main

import (
	"log"

	database "github.com/nedssoft/go-basic-api/bin/db"
	"github.com/nedssoft/go-basic-api/config"
)

func main() {
	db, err := database.NewDB(config.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
		database.Migrate(db)
	}

}
