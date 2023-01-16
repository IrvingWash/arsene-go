package utils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(url, path string) {
	out, err := os.Create(path)

	if err != nil {
		log.Fatal("failed to create file: ", err)
	}

	defer out.Close()

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("failed to download file: ", err)
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
