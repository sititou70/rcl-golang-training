package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	prettierHTML(doc, os.Stdout)

	return nil
}

func prettierHTML(n *html.Node, w io.Writer) {
	depth := 0

	forEachNode(n,
		func(n *html.Node) {
			if n.Type == html.ElementNode {
				attrs := ""
				for _, attr := range n.Attr {
					if attr.Val != "" {
						attrs += fmt.Sprintf(" %s='%s'", attr.Key, attr.Val)
					} else {
						attrs += fmt.Sprintf(" %s", attr.Key)
					}
				}

				closeSlash := ""
				if n.FirstChild == nil {
					closeSlash = " /"
				}

				fmt.Fprintf(w, "%*s<%s%s%s>\n", depth*2, "", n.Data, attrs, closeSlash)
				depth++
			}

			if n.Type == html.TextNode {
				fmt.Fprintf(w, "%*s%s\n", (depth)*2, "", n.Data)
			}

			if n.Type == html.CommentNode {
				fmt.Fprintf(w, "%*s<!-- %s -->\n", (depth)*2, "", n.Data)
			}
		},
		func(n *html.Node) {
			if n.Type == html.ElementNode {
				depth--
				if n.FirstChild != nil {
					fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
				}
			}
		},
	)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
