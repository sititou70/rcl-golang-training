// See page 216
package main

import (
	"sort"
	"strings"
	"testing"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// test data
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// multiKeySort
type multiKeySort struct {
	t    []*Track
	keys []string // ex: ["Title:asc", "Year:desc"]
}

func (x *multiKeySort) Len() int {
	return len(x.t)
}

func (x *multiKeySort) Less(i, j int) bool {
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

	for _, item := range x.keys {
		s := strings.Split(item, ":")
		key, mode := s[0], s[1]

		if equalTrack(x.t[i], x.t[j], key) {
			continue
		}
		if mode == "asc" {
			return lessTrack(x.t[i], x.t[j], key)
		} else {
			return lessTrack(x.t[j], x.t[i], key)
		}
	}

	return false
}

func (x *multiKeySort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

// test
func TestMultiKeySort(t *testing.T) {
	m := multiKeySort{tracks[:], []string{"Title:asc", "Year:desc"}}
	sort.Sort(&m)

	if m.t[0].Title != "Go" {
		t.Fail()
	}
	if m.t[0].Year != 2012 {
		t.Fail()
	}

	if m.t[1].Title != "Go" {
		t.Fail()
	}
	if m.t[1].Year != 1992 {
		t.Fail()
	}

	if m.t[2].Title != "Go Ahead" {
		t.Fail()
	}

	if m.t[3].Title != "Ready 2 Go" {
		t.Fail()
	}
}
