package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

type Reader struct {
	reader      io.Reader
	readedBytes int
	limitBytes  int
}

func (r *Reader) Read(p []byte) (int, error) {
	if r.readedBytes >= r.limitBytes {
		return 0, io.EOF
	}

	readBytes, err := r.reader.Read(p)
	if err != nil {
		return readBytes, err
	}

	limitedReadBytes := min(readBytes, r.limitBytes-r.readedBytes)
	r.readedBytes += limitedReadBytes
	return limitedReadBytes, nil
}
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	c := Reader{reader: r, limitBytes: int(n)}
	return &c
}

func TestLimitReader1(t *testing.T) {
	data, _ := os.ReadFile("main_test.go")

	r := LimitReader(bytes.NewReader(data), 100)

	readedData, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if len(readedData) != 100 {
		t.Fail()
	}
	if !reflect.DeepEqual(readedData, data[0:100]) {
		t.Fail()
	}
}

func TestLimitReader2(t *testing.T) {
	data, _ := os.ReadFile("main_test.go")

	r := LimitReader(bytes.NewReader(data), 9999)

	readedData, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if len(readedData) != len(data) {
		t.Fail()
	}
	if !reflect.DeepEqual(readedData, data) {
		t.Fail()
	}
}
