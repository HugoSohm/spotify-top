package server

import (
	"fmt"
	"github.com/HugoSohm/spotify-top/src/auth"
	"github.com/HugoSohm/spotify-top/src/business"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to Spotify Top API")
}

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", auth.SpotifyLogin)
	router.HandleFunc("/callback", auth.SpotifyCallback)
	router.HandleFunc("/refresh-token", auth.SpotifyRefreshToken)
	router.HandleFunc("/top/artists", business.GetTopArtists)
	router.HandleFunc("/top/tracks", business.GetTopTracks)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./images")))

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the server on :%s", os.Getenv("PORT")))
	}
}
