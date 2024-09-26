package config

import (
	"fmt"
	"os"
	 "github.com/lpernett/godotenv"
)

type Config struct {
  PublicHost string
	Port string
	DBUser string
	DBPassword string
	DBAddress string
	DBName string
}

func GetConnectionString() string {
	godotenv.Load(".env")
  return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",os.Getenv(("DB_HOST")), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
}
