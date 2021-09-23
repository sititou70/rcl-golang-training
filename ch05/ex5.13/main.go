// page 159
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"ex5.13/links"
)

func main() {
	target, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	crawl := func(item string) []string {
		url, err := url.Parse(item)
		if err != nil {
			panic(err)
		}
		if url.Host != target.Host {
			return []string{}
		}

		// fetch
		fmt.Printf("fetching: %s\n", url)
		resp, err := http.Get(url.String())
		if err != nil {
			return []string{}
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return []string{}
		}

		fsPath, err := saveCrawledContent(url.Path, resp.Body)
		if err != nil {
			fmt.Printf("saveCrawledContent: %s, %v", url.String(), err)
			return []string{}
		}
		f, err := os.Open(fsPath)
		if err != nil {
			fmt.Printf("Open crawled file: %s, %v", url.String(), err)
			return []string{}
		}
		defer f.Close()
		list, err := links.Extract(f, url.String())
		if err != nil {
			fmt.Printf("Extract: %s, %v", url.String(), err)
			return []string{}
		}

		return list
	}

	breadthFirst(crawl, []string{target.String()})
}

// utils
const CRAWL_DOCUMENT_ROOT = "TEMP_crawl_document_root"

func saveCrawledContent(path string, body io.Reader) (fsPath string, err error) {
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

	io.Copy(f, body)

	return
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
