package business

func ArtistTransformer(spotifyArtists []SpotifyArtist) []string {
	var result []string

	for _, item := range spotifyArtists {
		artist := Artist{
			Name:       item.Name,
			Url:        item.Href,
			PictureUrl: item.Images[1].Url,
			Genres:     item.Genres,
		}
		result = append(result, artist.Name)
	}

	return result
}

func TrackTransformer(spotifyTracks []SpotifyTrack) []string {
	var result []string

	for _, item := range spotifyTracks {
		track := Track{
			Name:       item.Name,
			ArtistName: item.Artists[0].Name,
			Url:        item.Href,
			PictureUrl: item.Album.Images[1].Url,
		}
		result = append(result, track.Name)
	}

	return result
}
