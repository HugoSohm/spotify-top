package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}
	if os.Getenv("PORT") == "" || os.Getenv("SPOTIFY_CLIENT_ID") == "" || os.Getenv("SPOTIFY_CLIENT_SECRET") == "" || os.Getenv("SPOTIFY_REDIRECT_URI") == "" {
		log.Fatal("Missing environment variables")
	}
}
