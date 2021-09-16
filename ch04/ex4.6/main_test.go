// page 104
package main

import (
	"testing"
	"unicode"
)

func normalizeSpace(s []byte) []byte {
	var norm []rune
	for _, r := range string(s) {
		if unicode.IsSpace(r) {
			norm = append(norm, ' ')
		} else {
			norm = append(norm, r)
		}
	}
	return uniq([]byte(string(norm)))
}
func uniq(s []byte) []byte {
	uniq := s[:1]
	currentString := s[0]
	for _, str := range s {
		if str != currentString {
			uniq = append(uniq, str)
			currentString = str
		}
	}
	return uniq
}

func TestUniq(t *testing.T) {
	if string(normalizeSpace([]byte("\t\n\v\f\r abc"))) != " abc" {
		t.Fail()
	}
}
