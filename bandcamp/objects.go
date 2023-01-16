package bandcamp

type BandcampAlbum struct {
	Current   BandcampCurrent     `json:"current"`
	HasAudio  bool                `json:"hasAudio"`
	Artist    string              `json:"artist"`
	TrackInfo []BandcampTrackInfo `json:"trackinfo"`
}

type BandcampCurrent struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}

type BandcampTrackInfo struct {
	File        BandcampFile `json:"file"`
	Title       string       `json:"title"`
	TrackNumber uint         `json:"track_num"`
	HasLyrics   bool         `json:"has_lyrics"`
}

type BandcampFile struct {
	MP3_128 string `json:"mp3-128"`
}
