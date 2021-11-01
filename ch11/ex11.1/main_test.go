package main

import (
	"os"
	"testing"
)

func TestCharcount(t *testing.T) {
	file, err := os.Open("main.go")
	if err != nil {
		t.Errorf("open main.go: %v\n", err)
		return
	}
	defer file.Close()

	counts, utflen, invalid, err := charcount(file)
	if err != nil {
		t.Errorf("charcount: %v\n", err)
		return
	}

	tests := []struct {
		title  string
		actual int
		want   int
	}{
		{"len(counts)", len(counts), 72},
		{"counts['a']", counts['a'], 46},
		{"utflen[1]", utflen[1], 1363},
		{"utflen[2]", utflen[2], 1},
		{"utflen[3]", utflen[3], 0},
		{"utflen[4]", utflen[4], 0},
		{"invalid", invalid, 0},
	}
	for _, test := range tests {
		if test.actual != test.want {
			t.Errorf("%s == %v, want %v\n", test.title, test.actual, test.want)
		}
	}
}
