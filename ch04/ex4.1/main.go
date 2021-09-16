package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Printf("diffs: %v\n", sha256DiffBits([]byte("x"), []byte("X")))
	fmt.Printf("diffs: %v\n", loopSHA256DiffBits([]byte("x"), []byte("X")))
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func sha256DiffBits(c1, c2 []byte) int {
	h1 := sha256.Sum256(c1)
	h2 := sha256.Sum256(c2)

	cnt := 0
	for i := 0; i < len(h1); i++ {
		cnt += int(pc[h1[i]^h2[i]])
	}

	return cnt
}

func loopSHA256DiffBits(c1, c2 []byte) int {
	h1 := sha256.Sum256(c1)
	h2 := sha256.Sum256(c2)

	cnt := 0
	for i := 0; i < len(h1); i++ {
		for j := 0; j < 8; j++ {
			bit1 := (h1[i] >> j) & 1
			bit2 := (h2[i] >> j) & 1
			if bit1 != bit2 {
				cnt++
			}
		}
	}

	return cnt
}
