package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"ex8.10/links"
)

func crawl(url string, cancel <-chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, cancel)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	cancel := make(chan struct{})

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

		Loop:
			for {
				var link string
				select {
				case link = <-unseenLinks:
				case <-cancel:
					break Loop
				}

				foundLinks := crawl(link, cancel)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(worklist)
	}()

	sigint := make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt)
	go func() {
		<-sigint
		close(cancel)
		for range unseenLinks {
		}
	}()

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
