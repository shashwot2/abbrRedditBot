package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getenv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	dotenv := getenv("rsecret")
}
