package main

import (
	"log"

	database "github.com/nedssoft/go-basic-api/bin/db"
	"github.com/nedssoft/go-basic-api/config"
	"github.com/nedssoft/go-basic-api/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration completed")
	return nil
}

func main() {
	db, err := database.NewDB(config.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
		Migrate(db)
	}

}
