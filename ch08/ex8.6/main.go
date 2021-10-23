package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
type Work struct {
	urls  []string
	depth int
}

type Link struct {
	url   string
	depth int
}

var (
	depth = flag.Int("depth", 0, "crawling depth")
)

func main() {
	flag.Parse()

	worklist := make(chan Work)    // lists of URLs, may have duplicates
	unseenLinks := make(chan Link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- Work{os.Args[2:], 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for unseenLink := range unseenLinks {
				foundLinks := crawl(unseenLink.url)
				go func(depth int) {
					worklist <- Work{foundLinks, depth + 1}
				}(unseenLink.depth)
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	// n: worklistを読むべき回数
	for n := 1; n > 0; n-- {
		work := <-worklist

		for _, url := range work.urls {
			if !seen[url] {
				seen[url] = true

				fmt.Println(url)

				if work.depth < *depth {
					n++
					unseenLinks <- Link{url, work.depth}
				}
			}
		}
	}
}
