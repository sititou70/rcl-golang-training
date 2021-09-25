package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByTagName(t *testing.T) {
	f, err := os.Open("test-data.html")
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	title := ElementByTagName(doc, "title")
	headers := ElementByTagName(doc, "h1", "h2", "h3", "h4", "h5", "h6")
	img := ElementByTagName(doc, "img")

	if len(title) != 1 {
		t.Fail()
	}
	if len(headers) != 4 {
		t.Fail()
	}
	if len(img) != 3 {
		t.Fail()
	}
}
