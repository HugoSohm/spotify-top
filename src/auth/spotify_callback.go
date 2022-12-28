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

type spotifyCallbackResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func SpotifyCallback(w http.ResponseWriter, r *http.Request) {
	// Get code from url
	code := r.URL.Query().Get("code")
	if code == "" {
		business.NewError(w, http.StatusBadRequest, "Missing code")
		return
	}

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
	var response spotifyCallbackResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		business.NewError(w, http.StatusInternalServerError, "Failed to jsonify body")
		return
	}

	var cookie = http.Cookie{
		Name:     "spotifytop_jwt",
		Value:    response.AccessToken,
		Path:     "/",
		Domain:   os.Getenv("FRONT_URL"),
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, os.Getenv("FRONT_URL"), http.StatusSeeOther)

	// Return the accessToken
	/*indentedResult, _ := json.MarshalIndent(response, "", "  ")
	_, _ = fmt.Fprintf(w, string(indentedResult))*/
}
