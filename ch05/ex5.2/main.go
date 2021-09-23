// page 141
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	countMap := map[string]int{}
	visit(countMap, doc)

	for k, v := range countMap {
		fmt.Printf("%-15s%5d\n", k, v)
	}
}

func visit(countMap map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		countMap[n.Data]++
	}

	visit(countMap, n.NextSibling)
	visit(countMap, n.FirstChild)
}
