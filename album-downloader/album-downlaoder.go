package albumdownloader

import (
	"arsene/objects"
	"fmt"
	"io"
	"log"
	"net/http"
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

	for _, track := range ad.album.Tracks {
		out, err := os.Create(fmt.Sprintf("%s/%s", ad.path, makeTrackName(track)))

		if err != nil {
			log.Fatal("failed to create file: ", err)
		}

		defer out.Close()

		resp, err := http.Get(track.URL)

		if err != nil {
			log.Fatal("failed to download track: ", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatal("status is not ok: ", resp.Status)
		}

		_, err = io.Copy(out, resp.Body)

		if err != nil {
			log.Fatal("failed to write data: ", err)
		}
	}
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
	out, err := os.Create(fmt.Sprintf("%s/cover.jpg", ad.path))

	if err != nil {
		log.Fatal("failed to create file: ", err)
	}

	defer out.Close()

	resp, err := http.Get(ad.album.AlbumArtURL)

	if err != nil {
		log.Fatal("failed to download album art: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("status is not ok: ", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal("failed to write data: ", err)
	}
}

func makeTrackName(track objects.TrackMetaInfo) string {
	return fmt.Sprintf("%s. %s.mp3", track.TrackNumber, track.Title)
}
