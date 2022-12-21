package business

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func GetTopTracks(w http.ResponseWriter, r *http.Request) {
	// Define the request
	spotifyTopArtistsUrl, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/tracks", nil)

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
		NewError(w, http.StatusInternalServerError, "Failed to execute query")
		return
	}

	// Read the response
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to read body")
		return
	}

	// Jsonify the response
	var response SpotifyTopTracksResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to jsonify body")
		return
	}

	// Transform response into data
	var result = TrackTransformer(response.Items)
	if result == nil {
		NewError(w, http.StatusInternalServerError, "Data is empty (try refreshing the token)")
		return
	}

	generateImage(w, r, result, "morceaux")

	// Format and display data
	/*indentedResult, _ := json.MarshalIndent(result, "", "  ")
	_, _ = fmt.Fprintf(w, string(indentedResult))*/
}
