package main

import (
	albumdownloader "arsene/album-downloader"
	"arsene/parsers"
	"arsene/utils"
	"log"
)

func main() {
	url, downloadPath := utils.InputHandler()

	parser, err := parsers.ParserFactory(url)

	if err != nil {
		log.Fatal(err)
	}

	album := parser.ParseAlbum()

	albumDownloader := albumdownloader.NewAlbumDownloader(album, downloadPath)

	albumDownloader.DownloadAlbum()
}
