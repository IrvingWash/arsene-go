package main

import (
	"arsene/bandcamp"
	"fmt"
)

func main() {
	bcp := bandcamp.NewBandcampParser("https://thecrinn.bandcamp.com/album/dreaming-saturn")

	bcAlbum := bcp.ParseAlbum()

	fmt.Println(bandcamp.ConvertBandcampAlbumIntoAlbum(*bcAlbum).ReleaseYear)
}
