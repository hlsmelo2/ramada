package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_PORT    int
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	SECRET_KEY  string
)

func envLoad() {
	var _error error

	if _error = godotenv.Load(); _error != nil {
		log.Fatal(_error)
	}
}

func Load() {
	envLoad()
	var _error error

	API_PORT, _error = strconv.Atoi(os.Getenv("API_PORT"))
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	SECRET_KEY = os.Getenv("SECRET_KEY")

	if _error != nil {
		API_PORT = 3001
	}
}
