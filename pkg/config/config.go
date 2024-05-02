package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(envFile string) {
	useSystemEnv := os.Getenv("USE_SYSTEM_ENV")

	if useSystemEnv == "true" {
		log.Print("using system environment variables")
	} else {
		log.Print("using .env environment variables")
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("cannot load environment file. reason: %v", err)
		}
	}
}
