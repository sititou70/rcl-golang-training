// page 169
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const FETCH_PREFIX = "TEMP_"

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	filename = path.Base(resp.Request.URL.Path)
	if filename == "/" {
		filename = "index.html"
	}
	filename = FETCH_PREFIX + filename
	f, err := os.Create(filename)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	return
}

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
