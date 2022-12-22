package business

import (
	"net/http"
)

func GetTopTracks(w http.ResponseWriter, r *http.Request) {
	// Get top artists from Spotify
	getSpotifyTop(w, r, "tracks")
}
