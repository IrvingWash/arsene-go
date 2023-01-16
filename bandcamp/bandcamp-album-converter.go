package bandcamp

import (
	"arsene/objects"
	"fmt"
	"log"
	"time"
)

func ConvertBandcampAlbumIntoAlbum(bcAlbum BandcampAlbumMetaInfoWithAlbumArtURL) objects.AlbumMetainfo {
	if !bcAlbum.HasAudio {
		log.Fatal("the album page has no audio")
	}

	album := objects.AlbumMetainfo{
		Artist: bcAlbum.Artist,
		Title:  bcAlbum.Current.Title,
	}

	releaseYear, err := time.Parse("02 Jan 2006 15:04:05 MST", bcAlbum.Current.ReleaseDate)

	if err != nil {
		fmt.Println("failed to parse release year, falling back to 0: ", err)
	} else {
		album.ReleaseYear = uint(releaseYear.Year())
	}

	return album
}
