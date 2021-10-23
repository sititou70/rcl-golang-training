// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 138.
//!+Extract

// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"io"
	"net/url"
	"path"

	"golang.org/x/net/html"
)

func Extract(body io.Reader, currentPath string) ([]string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %v", err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := url.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}

				if link.Host == "" {
					base, _ := url.Parse(currentPath)

					if len(link.Path) != 0 && link.Path[0] == '/' {
						base.Path = path.Join(link.Path)
					} else {
						base.Path = path.Join(base.Path, link.Path)
					}

					links = append(links, base.String())
				}

				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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
