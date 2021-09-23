package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		panic(fmt.Errorf("usage: ./%s URL ELEMENT_ID", os.Args[0]))
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	node := ElementByID(doc, os.Args[2])
	if node != nil {
		fmt.Printf("<%s %v />", node.Data, node.Attr)
	} else {
		fmt.Printf("element not found")
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node

	forEachNode(doc,
		func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				for _, attr := range n.Attr {
					if attr.Key == "id" && attr.Val == id {
						node = n
						return false
					}
				}
			}

			return true
		},
		nil,
	)

	return node
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && !pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && !post(n) {
		return
	}
}
