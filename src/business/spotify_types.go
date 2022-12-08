package business

type SpotifyTopArtistsResponse struct {
	Items []SpotifyArtist `json:"items"`
	AdditionalFields
}

type SpotifyTopTracksResponse struct {
	Items []SpotifyTrack `json:"items"`
	AdditionalFields
}

type SpotifyArtist struct {
	ExternalUrls SpotifyExternalUrls `json:"external_urls"`
	Followers    SpotifyFollowers    `json:"followers"`
	Genres       []string            `json:"genres"`
	Href         string              `json:"href"`
	Id           string              `json:"id"`
	Images       SpotifyImages       `json:"images"`
	Name         string              `json:"name"`
	Popularity   int                 `json:"popularity"`
	Type         string              `json:"type"`
	Uri          string              `json:"uri"`
}

type SpotifyTrack struct {
	Album            SpotifyAlbum        `json:"album"`
	Artists          []SpotifyArtist     `json:"artists"`
	AvailableMarkets []string            `json:"available_markets"`
	DiscNumber       int                 `json:"disc_number"`
	DurationMs       int                 `json:"duration_ms"`
	Explicit         bool                `json:"explicit"`
	ExternalIds      SpotifyExternalIds  `json:"external_ids"`
	ExternalUrls     SpotifyExternalUrls `json:"external_urls"`
	Href             string              `json:"href"`
	Id               string              `json:"id"`
	IsLocal          bool                `json:"is_local"`
	Name             string              `json:"name"`
	Popularity       int                 `json:"popularity"`
	PreviewUrl       string              `json:"preview_url"`
	TrackNumber      int                 `json:"track_number"`
	Type             string              `json:"type"`
	Uri              string              `json:"uri"`
}

type SpotifyAlbum struct {
	AlbumType            string              `json:"album_type"`
	Artists              []Artist            `json:"artists"`
	AvailableMarkets     []string            `json:"available_markets"`
	ExternalUrls         SpotifyExternalUrls `json:"external_urls"`
	Href                 string              `json:"href"`
	Id                   string              `json:"id"`
	Images               SpotifyImages       `json:"images"`
	Name                 string              `json:"name"`
	ReleaseDate          string              `json:"release_date"`
	ReleaseDatePrecision string              `json:"release_date_precision"`
	TotalTracks          int                 `json:"total_tracks"`
	Type                 string              `json:"type"`
	Uri                  string              `json:"uri"`
}

type SpotifyFollowers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type SpotifyExternalUrls struct {
	Spotify string `json:"spotify"`
}

type SpotifyExternalIds struct {
	Isrc string `json:"isrc"`
}

type SpotifyImages []struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}
