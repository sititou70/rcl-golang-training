// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bzip

import (
	"bytes"
	"compress/bzip2" // reader
	"io"
	"sync"
	"testing"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := NewWriter(&compressed)

	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		uncompressed.WriteString("hello")
	}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			w.Write([]byte("hello"))
			wg.Done()
		}()
	}
	wg.Wait()
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 50; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}
