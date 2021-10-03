package main

import (
	"io"
	"testing"

	"golang.org/x/net/html"
)

type Reader struct {
	str         *string
	readedBytes int
}

func (r *Reader) Read(p []byte) (int, error) {
	if r.readedBytes >= len(*r.str) {
		return 0, io.EOF
	}

	copied := copy(p, []byte(*r.str))
	r.readedBytes += copied
	return copied, nil
}

func NewReader(s string) *Reader {
	r := Reader{str: &s}
	return &r
}

func TestNewReader(t *testing.T) {
	r := NewReader(`
	<html>
		<head>
			<title>title</title>
		</head>
		<body>
			<p>hello</p>
		</body>
	</html>`)
	doc, err := html.Parse(r)

	if err != nil {
		t.Fatal(err)
	}
	if doc.FirstChild.LastChild.FirstChild.NextSibling.FirstChild.Data != "hello" {
		t.Fail()
	}
}
