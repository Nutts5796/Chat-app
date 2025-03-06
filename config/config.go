package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl      string
	JWTSecret  string
	ServerPort string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	return Config{
		DBUrl:      os.Getenv("DB_URL"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
