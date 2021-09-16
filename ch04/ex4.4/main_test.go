// page 104
package main

import (
	"bytes"
	"testing"
)

func rotateLeft(s []byte, num int) []byte {
	index := num % len(s)
	rotated := s[index:]
	rotated = append(rotated, s[:index]...)
	return rotated
}

func TestRotateLeft(t *testing.T) {
	if !bytes.Equal(rotateLeft([]byte{1, 2, 3, 4, 5}, 0), []byte{1, 2, 3, 4, 5}) {
		t.Fail()
	}
	if !bytes.Equal(rotateLeft([]byte{1, 2, 3, 4, 5}, 1), []byte{2, 3, 4, 5, 1}) {
		t.Fail()
	}
	if !bytes.Equal(rotateLeft([]byte{1, 2, 3, 4, 5}, 2), []byte{3, 4, 5, 1, 2}) {
		t.Fail()
	}
	if !bytes.Equal(rotateLeft([]byte{1, 2, 3, 4, 5}, 3), []byte{4, 5, 1, 2, 3}) {
		t.Fail()
	}
	if !bytes.Equal(rotateLeft([]byte{1, 2, 3, 4, 5}, 4), []byte{5, 1, 2, 3, 4}) {
		t.Fail()
	}
	if !bytes.Equal(rotateLeft([]byte{1, 2, 3, 4, 5}, 5), []byte{1, 2, 3, 4, 5}) {
		t.Fail()
	}
	if !bytes.Equal(rotateLeft([]byte{1, 2, 3, 4, 5}, 6), []byte{2, 3, 4, 5, 1}) {
		t.Fail()
	}
}
