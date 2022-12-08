package server

import (
	"fmt"
	"github.com/HugoSohm/spotify-top/src/auth"
	"github.com/HugoSohm/spotify-top/src/business"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.SpotifyLogin)
	mux.HandleFunc("/callback", auth.SpotifyCallback)
	mux.HandleFunc("/top/artists", business.GetTopArtists)
	mux.HandleFunc("/top/tracks", business.GetTopTracks)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), mux)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to start the server on :%s", os.Getenv("PORT")))
	}
}
