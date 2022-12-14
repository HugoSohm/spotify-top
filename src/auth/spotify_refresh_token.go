package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/HugoSohm/spotifytop-api/src/business"
	"io"
	"net/http"
	"os"
)

type spotifyRefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func SpotifyRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get refresh token from url
	refreshToken := r.URL.Query().Get("refresh_token")
	if refreshToken == "" {
		business.NewError(w, http.StatusBadRequest, "Missing refresh_token")
		return
	}

	// Define request
	spotifyRefreshTokenRequest, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)

	// Add query headers
	spotifyRefreshTokenRequest.Header = http.Header{
		"Accept":        {"application/json"},
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET")))))},
	}

	// Add query parameters
	params := spotifyRefreshTokenRequest.URL.Query()
	params.Add("refresh_token", refreshToken)
	params.Add("grant_type", "refresh_token")
	spotifyRefreshTokenRequest.URL.RawQuery = params.Encode()

	// Execute http query
	res, err := http.DefaultClient.Do(spotifyRefreshTokenRequest)
	if err != nil {
		business.NewError(w, http.StatusInternalServerError, "Failed to execute query")
		return
	}

	// Read the body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		business.NewError(w, http.StatusInternalServerError, "Failed to read body")
		return
	}

	// Jsonify the body
	var response spotifyRefreshTokenResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		business.NewError(w, http.StatusInternalServerError, "Failed to jsonify body")
		return
	}

	// Return the accessToken
	indentedResult, _ := json.MarshalIndent(response, "", "  ")
	_, _ = fmt.Fprintf(w, string(indentedResult))
}
