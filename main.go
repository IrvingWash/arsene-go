package main

import (
	albumdownloader "arsene/album-downloader"
	"arsene/bandcamp"
	"arsene/utils"
)

func main() {
	url, downloadPath := utils.InputHandler()

	bcp := bandcamp.NewBandcampParser(url)

	album := bcp.ParseAlbum()

	albumDownloader := albumdownloader.NewAlbumDownloader(album, downloadPath)

	albumDownloader.DownloadAlbum()
}
