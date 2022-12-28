package server

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	_ = godotenv.Load(".env")

	if os.Getenv("PORT") == "" || os.Getenv("FRONT_URL") == "" || os.Getenv("SPOTIFY_CLIENT_ID") == "" || os.Getenv("SPOTIFY_CLIENT_SECRET") == "" || os.Getenv("SPOTIFY_REDIRECT_URI") == "" {
		panic("Missing environment variables")
	}
}
