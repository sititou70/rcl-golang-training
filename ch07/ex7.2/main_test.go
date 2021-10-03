package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

type ByteCounter struct {
	writer io.Writer
	count  int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return n, err
	}

	c.count += int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{writer: w}
	return &c, &c.count
}

func TestCountingWriter(t *testing.T) {
	data, _ := os.ReadFile("main_test.go")

	buf := new(bytes.Buffer)
	w, c := CountingWriter(buf)
	fmt.Fprintf(w, "%s", data)

	if !reflect.DeepEqual(buf.Bytes(), data) {
		t.Fail()
	}
	if *c != int64(len(data)) {
		t.Fail()
	}
}
