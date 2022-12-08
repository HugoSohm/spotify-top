package business

func ArtistTransformer(spotifyArtists []SpotifyArtist) []Artist {
	var result []Artist

	for _, item := range spotifyArtists {
		artist := Artist{
			Name:       item.Name,
			Url:        item.Href,
			PictureUrl: item.Images[1].Url,
			Genres:     item.Genres,
		}
		result = append(result, artist)
	}

	return result
}

func TrackTransformer(spotifyTracks []SpotifyTrack) []Track {
	var result []Track

	for _, item := range spotifyTracks {
		track := Track{
			Name:       item.Name,
			ArtistName: item.Artists[0].Name,
			Url:        item.Href,
			PictureUrl: item.Album.Images[1].Url,
		}
		result = append(result, track)
	}

	return result
}
