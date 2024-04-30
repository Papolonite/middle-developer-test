package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("cannot load environment file. reason: %v", err)
	}

}
