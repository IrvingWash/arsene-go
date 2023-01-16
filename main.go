package main

import (
	"arsene/bandcamp"
	"fmt"
)

func main() {
	bcp := bandcamp.NewBandcampParser("https://thecrinn.bandcamp.com/album/dreaming-saturn")

	album := bcp.ParseAlbum()

	fmt.Println(album.ReleaseYear)
}
