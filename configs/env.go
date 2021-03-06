package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ошибка загрузки .env файла")
	}

	return os.Getenv("MONGOURI")
}
