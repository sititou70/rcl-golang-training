// page 144
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		panic("too few argument")
	}

	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("words: %d, images: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	visit(n, func(n *html.Node) bool {
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style" || n.Data == "noscript") {
			return false
		}

		if n.Type == html.TextNode {
			s := strings.Split(n.Data, " ")
			words += len(s)
		}

		return true
	})

	visit(n, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}
		return true
	})

	return
}

func visit(n *html.Node, f func(n *html.Node) bool) {
	if !f(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, f)
	}
}
