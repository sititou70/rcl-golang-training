package main

import (
	"bytes"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettierHTML(t *testing.T) {
	f, err := os.Open("test-data.html")
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		panic(err)
	}
	buffer := bytes.Buffer{}
	prettierHTML(doc, &buffer)

	_, err = html.Parse(&buffer)
	if err != nil {
		t.Errorf("parsing pretterd string: %v", err)
	}
}
