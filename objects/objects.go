package objects

type AlbumMetainfo struct {
	Title       string
	Artist      string
	ReleaseYear uint
	Tracks      []TrackMetainfo
	AlbumArtURL string
}

type TrackMetainfo struct {
	Title       string
	TrackNumber uint
	URL         string
}
