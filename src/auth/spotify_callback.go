package auth

import "net/http"

func SpotifyCallback(w http.ResponseWriter, r *http.Request) {
	// Get code from url
	code := r.URL.Query().Get("code")

	// Exchange code for accessToken
	SpotifyExchange(w, code)
}
