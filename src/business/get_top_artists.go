package business

import (
	"net/http"
)

func GetTopArtists(w http.ResponseWriter, r *http.Request) {
	// Get top artists from Spotify
	getSpotifyTop(w, r, "artists")
}
