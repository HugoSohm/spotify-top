package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type (
	SpotifyTokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}
	SpotifyTopArtists struct {
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			Id     string   `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Type       string `json:"type"`
			Uri        string `json:"uri"`
		} `json:"items"`
	}
)

func spotifyLogin(w http.ResponseWriter, r *http.Request) {
	// Define url
	spotifyAuthorizeUrl, err := url.Parse("https://accounts.spotify.com/authorize?")
	if err != nil {
		log.Fatal(err)
	}

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

func spotifyCallback(_ http.ResponseWriter, r *http.Request) {
	// Get code from url
	code := r.URL.Query().Get("code")

	// Exchange code for accessToken
	spotifyExchangeCode(code)
}

func spotifyExchangeCode(code string) {
	// Define request
	spotifyTokenRequest, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)

	// Add query headers
	spotifyTokenRequest.Header = http.Header{
		"Accept":        {"application/json"},
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET")))))},
	}

	// Add query parameters
	params := spotifyTokenRequest.URL.Query()
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", os.Getenv("SPOTIFY_REDIRECT_URI"))
	spotifyTokenRequest.URL.RawQuery = params.Encode()

	// Execute http query
	res, err := http.DefaultClient.Do(spotifyTokenRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Read the body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Jsonify the body
	var response SpotifyTokenResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Get the accessToken
	spotifyTopArtists(response.AccessToken)
}

func spotifyTopArtists(accessToken string) {
	// Define the request
	spotifyTopArtistsUrl, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/artists", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the headers using the accessToken
	spotifyTopArtistsUrl.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", accessToken)},
	}

	// Execute the request
	res, err := http.DefaultClient.Do(spotifyTopArtistsUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Read the body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Jsonify the body
	var response SpotifyTopArtists
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Display the body
	log.Fatal(response.Items)
}

func main() {
	// Setup environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	// Setup http server
	mux := http.NewServeMux()
	mux.HandleFunc("/login", spotifyLogin)
	mux.HandleFunc("/callback", spotifyCallback)

	// Start http server
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start the server on :8080")
	}
}
