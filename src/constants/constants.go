package constants

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadEnvOnce sync.Once

func GetConstant(key string) string {
	loadEnvOnce.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("No .env file found; using environment variables from runtime")
		}
	})

	return os.Getenv(key)
}
