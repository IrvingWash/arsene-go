package main

import (
	albumdownloader "arsene/album-downloader"
	"arsene/bandcamp"
)

func main() {
	bcp := bandcamp.NewBandcampParser("https://thecrinn.bandcamp.com/album/dreaming-saturn")

	album := bcp.ParseAlbum()

	albumDownloader := albumdownloader.NewAlbumDownloader(album, "~/Downloads")

	albumDownloader.DownloadAlbum()
}
