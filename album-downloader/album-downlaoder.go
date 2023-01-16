package albumdownloader

import (
	"arsene/objects"
	"arsene/utils"
	"fmt"
	"log"
	"os"
)

type AlbumDownloader struct {
	album             objects.AlbumMetaInfo
	downloadDirectory string
	path              string
}

func NewAlbumDownloader(album objects.AlbumMetaInfo, downloadDir string) *AlbumDownloader {
	return &AlbumDownloader{
		album:             album,
		downloadDirectory: downloadDir,
	}
}

func (ad *AlbumDownloader) DownloadAlbum() {
	ad.path = ad.makePath()
	ad.createDownloadDir()

	ad.downloadAlbumArt()

	fmt.Printf("Started downloading to %s\n", ad.path)

	for _, track := range ad.album.Tracks {
		trackPath := fmt.Sprintf("%s/%s", ad.path, makeTrackName(track))

		utils.DownloadFile(track.URL, trackPath)
	}

	fmt.Printf("Finished downloading")
}

func (ad *AlbumDownloader) makePath() string {
	return fmt.Sprintf(
		"%s/%s/%d - %s",
		ad.downloadDirectory,
		ad.album.Artist,
		ad.album.ReleaseYear,
		ad.album.Title,
	)
}

func (ad *AlbumDownloader) createDownloadDir() {
	if err := os.MkdirAll(ad.path, os.ModePerm); err != nil {
		log.Fatal("failed to create directory: ", err)
	}
}

func (ad *AlbumDownloader) downloadAlbumArt() {
	coverPath := fmt.Sprintf("%s/cover.jpg", ad.path)

	utils.DownloadFile(ad.album.AlbumArtURL, coverPath)
}

func makeTrackName(track objects.TrackMetaInfo) string {
	return fmt.Sprintf("%s. %s.mp3", track.TrackNumber, track.Title)
}
