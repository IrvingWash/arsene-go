package bandcamp

import (
	"arsene/objects"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type BandcampParser struct {
	albumURL     string
	albumPageSRC string
}

func NewBandcampParser(albumURL string) *BandcampParser {
	return &BandcampParser{
		albumURL: albumURL,
	}
}

func (bcp *BandcampParser) ParseAlbum() objects.AlbumMetaInfo {
	bcp.albumPageSRC = bcp.albumPageSourceCode()

	albumJSON := bcp.albumMetaInfoJSON()

	bcAlbum := bcp.albumMetainfo(albumJSON)
	albumArtURL := bcp.albumArtURL()

	return convertBandcampAlbumIntoAlbum(bcAlbum, albumArtURL)
}

func (bcp *BandcampParser) albumMetainfo(albumJSON string) BandcampAlbumMetaInfo {
	bcAlbum := &BandcampAlbumMetaInfo{}

	if err := json.Unmarshal([]byte(albumJSON), bcAlbum); err != nil {
		log.Fatal("failed to parse json into Bandcamp Album MetaInfo: ", err)
	}

	return *bcAlbum
}

func (bcp *BandcampParser) albumPageSourceCode() string {
	resp, err := http.Get(bcp.albumURL)

	if err != nil {
		log.Fatal("failed to get album page source code: ", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("failed to read album page response: ", err)
	}

	return string(body)
}

func (bcp *BandcampParser) albumMetaInfoJSON() string {
	const albumDataAttrCodeStart = "data-tralbum=\""
	const albumDataAttrCodeEnd = "}\""

	albumJSONStartID := strings.Index(bcp.albumPageSRC, albumDataAttrCodeStart)

	if albumJSONStartID == -1 {
		log.Fatal("failed to find album JSON in page source code")
	}

	albumJSONTemp := bcp.albumPageSRC[albumJSONStartID+len(albumDataAttrCodeStart):]

	albumJSONEndID := strings.Index(albumJSONTemp, albumDataAttrCodeEnd) + 1

	albumJSONTemp = albumJSONTemp[:albumJSONEndID]

	return strings.ReplaceAll(albumJSONTemp, "&quot;", "\"")
}

func (bcp *BandcampParser) albumArtURL() string {
	const urlCodeStart = "src=\"https://"
	const urlCodeEnd = "jpg\""

	albumArtCodeChunk := bcp.albumPageSRC[strings.Index(bcp.albumPageSRC, "<div id=\"tralbumArt\">"):]

	urlCodeStartID := strings.Index(albumArtCodeChunk, urlCodeStart) + len(urlCodeStart)

	urlCodeChunkTemp := albumArtCodeChunk[urlCodeStartID:]

	urlCodeEndID := strings.Index(urlCodeChunkTemp, urlCodeEnd) + len(urlCodeEnd) - 1

	return "http://" + urlCodeChunkTemp[:urlCodeEndID]
}
