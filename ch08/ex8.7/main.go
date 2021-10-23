package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"ex8.7/links"
)

func main() {
	rootTarget := os.Args[1]
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- []string{rootTarget} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for unseenLink := range unseenLinks {
				foundLinks := crawl(unseenLink, rootTarget)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	// n: worklistを読むべき回数
	for n := 1; n > 0; n-- {
		work := <-worklist

		for _, link := range work {
			if !seen[link] {
				seen[link] = true

				if isSameHost(link, rootTarget) {
					n++
					unseenLinks <- link
				}
			}
		}
	}
}

// utils
func isSameHost(url1, url2 string) bool {
	parsed1, err := url.Parse(url1)
	if err != nil {
		return false
	}
	parsed2, err := url.Parse(url2)
	if err != nil {
		return false
	}

	return parsed1.Host == parsed2.Host
}

func crawl(targetUrl string, rootUrl string) []string {
	target, err := url.Parse(targetUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("fetching: %s\n", target.String())
	resp, err := http.Get(target.String())
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return []string{}
	}

	fsPath, err := saveCrawledContent(target.Path, resp.Body, rootUrl)
	if err != nil {
		fmt.Printf("saveCrawledContent: %s, %v", target.String(), err)
		return []string{}
	}
	f, err := os.Open(fsPath)
	if err != nil {
		fmt.Printf("Open crawled file: %s, %v", target.String(), err)
		return []string{}
	}
	defer f.Close()
	list, err := links.Extract(f, target.String())
	if err != nil {
		fmt.Printf("Extract: %s, %v", target.String(), err)
		return []string{}
	}

	return list
}

const CRAWL_DOCUMENT_ROOT = "TEMP_crawl_document_root"

func saveCrawledContent(path string, body io.Reader, rootUrl string) (fsPath string, err error) {
	fsPath = CRAWL_DOCUMENT_ROOT + "/" + path
	if len(fsPath) != 0 && fsPath[len(fsPath)-1] == '/' || !strings.Contains(filepath.Base(path), ".") {
		fsPath = filepath.Join(fsPath, "index.html")
	}

	err = os.MkdirAll(filepath.Dir(fsPath), 0775)
	if err != nil {
		return
	}

	f, err := os.Create(fsPath)
	if err != nil {
		return
	}
	defer f.Close()

	parsedRoot, err := url.Parse(rootUrl)
	if err != nil {
		return
	}
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return
	}
	replacedBody := strings.ReplaceAll(string(bodyBytes), parsedRoot.Host, "/")

	_, err = f.Write([]byte(replacedBody))
	if err != nil {
		return
	}

	return
}
