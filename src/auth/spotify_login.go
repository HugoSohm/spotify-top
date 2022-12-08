package auth

import (
	"net/http"
	"net/url"
	"os"
)

func SpotifyLogin(w http.ResponseWriter, r *http.Request) {
	// Define url
	spotifyAuthorizeUrl, _ := url.Parse("https://accounts.spotify.com/authorize?")

	// Add query parameters
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", "5343066bd0724a9fa65fabe78bae1a31")
	params.Add("scope", "user-read-email user-top-read")
	params.Add("redirect_uri", os.Getenv("SPOTIFY_REDIRECT_URI"))
	params.Add("state", "test")
	spotifyAuthorizeUrl.RawQuery = params.Encode()

	http.Redirect(w, r, spotifyAuthorizeUrl.String(), http.StatusSeeOther)
}
