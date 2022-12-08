package business

type Artist struct {
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	PictureUrl string   `json:"picture_url"`
	Genres     []string `json:"genres"`
}

type Track struct {
	Name       string `json:"name"`
	ArtistName string `json:"artist_name"`
	Url        string `json:"url"`
	PictureUrl string `json:"picture_url"`
}

type AdditionalFields struct {
	href     string
	limit    int
	next     string
	offset   int
	previous string
	total    int
}
