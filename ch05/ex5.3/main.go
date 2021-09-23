// page 141
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	text := visit(nil, doc)
	fmt.Printf("%v\n", strings.Join(text, ""))
}

func visit(text []string, n *html.Node) []string {
	if n == nil {
		return text
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style" || n.Data == "noscript") {
		return text
	}

	if n.Type == html.TextNode {
		text = append(text, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = visit(text, c)
	}
	return text
}
