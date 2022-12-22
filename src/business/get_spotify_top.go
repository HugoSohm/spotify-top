package business

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func getSpotifyTop(w http.ResponseWriter, r *http.Request, topType string) {
	// Define the request
	spotifyUrl := fmt.Sprintf("https://api.spotify.com/v1/me/top/%s", topType)
	spotifyRequest, _ := http.NewRequest("GET", spotifyUrl, nil)

	// Get the accessToken from request headers
	authorization := r.URL.Query().Get("access_token")
	if authorization == "" {
		NewError(w, http.StatusUnauthorized, "Missing access_token query string parameter")
	}

	// Set the headers using the accessToken
	spotifyRequest.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", authorization)},
	}

	// Get and define query parameters
	params := url.Values{}
	timeRange := r.URL.Query().Get("time_range")
	if timeRange != "" {
		if timeRange != "short_term" && timeRange != "medium_term" && timeRange != "long_term" {
			NewError(w, http.StatusBadRequest, "Invalid time_range (short_term, medium_term, long_term)")
		}
		params.Add("time_range", timeRange)
		spotifyRequest.URL.RawQuery = params.Encode()
	}

	// Execute the request
	res, err := http.DefaultClient.Do(spotifyRequest)
	if err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to execute query")
	}

	// Read the body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to read body")
	}

	switch topType {
	case "artists":
		// Jsonify the body
		var response SpotifyTopArtistsResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			NewError(w, http.StatusInternalServerError, "Failed to jsonify body")
		}

		// Transform response into data
		var result = ArtistTransformer(response.Items)
		if result == nil {
			NewError(w, http.StatusInternalServerError, "Data is empty (try refreshing the token)")
		}

		generateImage(w, r, result, "artists")
	case "tracks":
		// Jsonify the body
		var response SpotifyTopTracksResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			NewError(w, http.StatusInternalServerError, "Failed to jsonify body")
		}

		// Transform response into data
		var result = TrackTransformer(response.Items)
		if result == nil {
			NewError(w, http.StatusInternalServerError, "Data is empty (try refreshing the token)")
		}

		generateImage(w, r, result, "tracks")
	default:
		NewError(w, http.StatusInternalServerError, "Top type must be 'artists' or 'tracks'")
	}
}
