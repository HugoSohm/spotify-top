package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	_ = godotenv.Load(".env")

	if os.Getenv("PORT") == "" || os.Getenv("SPOTIFY_CLIENT_ID") == "" || os.Getenv("SPOTIFY_CLIENT_SECRET") == "" || os.Getenv("SPOTIFY_REDIRECT_URI") == "" {
		log.Fatal("Missing environment variables")
	}
}
