package main

import "testing"

func TestSHADiff(t *testing.T) {
	if sha256DiffBits([]byte(""), []byte("")) != loopSHA256DiffBits([]byte(""), []byte("")) {
		t.Fail()
	}
	if sha256DiffBits([]byte("asd"), []byte("asd")) != loopSHA256DiffBits([]byte("asd"), []byte("asd")) {
		t.Fail()
	}
	if sha256DiffBits([]byte(""), []byte("")) != 0 {
		t.Fail()
	}
	if sha256DiffBits([]byte("asd"), []byte("asd")) != 0 {
		t.Fail()
	}

	if sha256DiffBits([]byte("asd"), []byte("zxc")) != loopSHA256DiffBits([]byte("asd"), []byte("zxc")) {
		t.Fail()
	}
}
