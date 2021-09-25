package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		panic("usage: go run main.go URL TAG_NAME1, TAG_NAME2, ...")
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

	nodes := ElementByTagName(doc, os.Args[2:]...)
	for _, node := range nodes {
		if node.FirstChild != nil {
			fmt.Printf("<%s %v>%s</%s>\n", node.Data, node.Attr, node.FirstChild.Data, node.Data)
		} else {
			fmt.Printf("<%s %v>\n", node.Data, node.Attr)
		}
	}
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	nodes := []*html.Node{}
	forEachNode(
		doc,
		func(n *html.Node) {
			for _, item := range name {
				if n.Type == html.ElementNode && item == n.Data {
					nodes = append(nodes, n)
				}
			}
		},
		nil,
	)

	return nodes
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
