// See page 216
package track

import (
	"strings"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type MultiKeySort struct {
	Track []*Track
	Keys  []string // ex: ["Title:asc", "Year:desc"]
}

func (x *MultiKeySort) Len() int {
	return len(x.Track)
}

func (x *MultiKeySort) Less(i, j int) bool {
	equalTrack := func(x, y *Track, key string) bool {
		switch key {
		case "Title":
			return x.Title == y.Title
		case "Artist":
			return x.Artist == y.Artist
		case "Album":
			return x.Album == y.Album
		case "Year":
			return x.Year == y.Year
		case "Length":
			return x.Length == y.Length
		}

		return false
	}
	lessTrack := func(x, y *Track, key string) bool {
		switch key {
		case "Title":
			return x.Title < y.Title
		case "Artist":
			return x.Artist < y.Artist
		case "Album":
			return x.Album < y.Album
		case "Year":
			return x.Year < y.Year
		case "Length":
			return x.Length < y.Length
		}

		return false
	}

	for _, item := range x.Keys {
		s := strings.Split(item, ":")
		key, mode := s[0], s[1]

		if equalTrack(x.Track[i], x.Track[j], key) {
			continue
		}
		if mode == "asc" {
			return lessTrack(x.Track[i], x.Track[j], key)
		} else {
			return lessTrack(x.Track[j], x.Track[i], key)
		}
	}

	return false
}

func (x *MultiKeySort) Swap(i, j int) {
	x.Track[i], x.Track[j] = x.Track[j], x.Track[i]
}
