package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type spotifyTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func SpotifyExchange(w http.ResponseWriter, code string) {
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
	var response spotifyTokenResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Return the accessToken
	_, err = fmt.Fprintf(w, response.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
}
