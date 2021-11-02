package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input  string
		sep    string
		length int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c", "b", 2},
		{"a:b:c", ":b:", 2},
		{",asdf,dsf,sdf,sa,dsf,adsf,sd,fsadsf,", ",", 10},
	}

	for _, test := range tests {
		words := strings.Split(test.input, test.sep)
		if len(words) != test.length {
			t.Errorf("Split(%v,%v) returned %d words, want %d", test.input, test.sep, len(words), test.length)
		}
	}
}
