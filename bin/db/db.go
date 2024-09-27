package db

import (
	"log"

	"github.com/nedssoft/go-basic-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(connString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration completed")
	return nil
}