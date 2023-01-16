package objects

type AlbumMetaInfo struct {
	Title       string
	Artist      string
	ReleaseYear uint
	Tracks      []TrackMetaInfo
	AlbumArtURL string
}

type TrackMetaInfo struct {
	Title       string
	TrackNumber string
	URL         string
}
