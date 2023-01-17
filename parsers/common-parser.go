package parsers

import "arsene/objects"

type ICommonParser interface {
	ParseAlbum() objects.AlbumMetaInfo
}
