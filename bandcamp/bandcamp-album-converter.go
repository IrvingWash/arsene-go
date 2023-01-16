package bandcamp

import (
	"fmt"
	"log"
	"time"

	"arsene/objects"
)

func convertBandcampAlbumIntoAlbum(bcAlbum BandcampAlbumMetaInfo, albumArtURL string) objects.AlbumMetaInfo {
	if !bcAlbum.HasAudio {
		log.Fatal("the album page has no audio")
	}

	album := objects.AlbumMetaInfo{
		Artist:      bcAlbum.Artist,
		Title:       bcAlbum.Current.Title,
		AlbumArtURL: albumArtURL,
		Tracks:      convertBandcampTracksIntoTracks(bcAlbum.TrackInfo),
	}

	releaseYear, err := time.Parse("02 Jan 2006 15:04:05 MST", bcAlbum.Current.ReleaseDate)

	if err != nil {
		fmt.Println("failed to parse release year, falling back to 0: ", err)
	} else {
		album.ReleaseYear = uint(releaseYear.Year())
	}

	return album
}

func convertBandcampTracksIntoTracks(bcTracks []BandcampTrackInfo) []objects.TrackMetaInfo {
	var tracks []objects.TrackMetaInfo

	for _, track := range bcTracks {
		var trackNumber string

		if track.Track_Num > 9 {
			trackNumber = fmt.Sprint(track.Track_Num)
		} else {
			trackNumber = fmt.Sprintf("0%d", track.Track_Num)
		}

		tracks = append(tracks, objects.TrackMetaInfo{
			Title:       track.Title,
			TrackNumber: trackNumber,
			URL:         track.File.MP3_128,
		})
	}

	return tracks
}
