package main

import (
	"log"

	"github.com/nedssoft/learn-go/config"
	"github.com/nedssoft/learn-go/models"
	"gorm.io/gorm"
	database "github.com/nedssoft/learn-go/bin/db"
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