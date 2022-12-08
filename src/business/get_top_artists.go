package business

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func GetTopArtists(w http.ResponseWriter, r *http.Request) {
	// Define the request
	spotifyTopArtistsUrl, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/artists", nil)

	// Get the accessToken from request headers
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		NewError(w, http.StatusUnauthorized, "Missing authorization header")
		return
	}

	// Set the headers using the accessToken
	spotifyTopArtistsUrl.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {authorization},
	}

	// Get and define query parameters
	params := url.Values{}
	timeRange := r.URL.Query().Get("time_range")
	if timeRange != "" {
		if timeRange != "short_term" && timeRange != "medium_term" && timeRange != "long_term" {
			NewError(w, http.StatusBadRequest, "Invalid time_range (short_term, medium_term, long_term)")
			return
		}
		params.Add("time_range", timeRange)
		spotifyTopArtistsUrl.URL.RawQuery = params.Encode()
	}

	// Execute the request
	res, err := http.DefaultClient.Do(spotifyTopArtistsUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Read the response
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Jsonify the response
	var response SpotifyTopArtistsResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Transform response into data
	var result = ArtistTransformer(response.Items)

	// Format and display data
	indentedResult, _ := json.MarshalIndent(result, "", "  ")
	_, _ = fmt.Fprintf(w, string(indentedResult))
}
