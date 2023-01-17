package parsers

import (
	"arsene/parsers/bandcamp"
	"fmt"
	"strings"
)

func ParserFactory(albumURL string) (ICommonParser, error) {
	if strings.Contains(albumURL, "bandcamp") {
		return bandcamp.NewBandcampParser(albumURL), nil
	}

	return nil, fmt.Errorf("Unsupported url")
}
