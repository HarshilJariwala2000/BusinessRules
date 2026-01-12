package constants

import (
	"os"
	"github.com/joho/godotenv"
	"log"
)

const (

)

type Config struct{
	DBHost string
	DBUser string
	DBPassword string
	DBPort string
	DBName string
	SSLMode string
}

var AppConfig *Config

func Load(){
	err := godotenv.Load()
	if err!=nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = &Config{
		DBHost: os.Getenv("PGHOST"),
		DBUser: os.Getenv("PGUSER"),
		DBPassword: os.Getenv("PGPASSWORD"),
		DBName: os.Getenv("PGDATABASE"),
		SSLMode: os.Getenv("PGSSLMODE"),
		DBPort: os.Getenv("PGPORT"),
	}
}