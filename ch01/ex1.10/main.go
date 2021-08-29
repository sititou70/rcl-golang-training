// page 21
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	secs := time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// get body sise
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error: read body %s: %v", url, err)
		return
	}
	resp.Body.Close() // don't leak resources

	// check cache
	cacheStatus := "new"
	const cachePrefix = "TEMP_"
	cachePath := "./" + cachePrefix + sanitiseForFSPath(url)
	cache, err := ioutil.ReadFile(cachePath)
	if err == nil {
		if reflect.DeepEqual(body, cache) {
			cacheStatus = "no change"
		} else {
			cacheStatus = "change"
		}
	}
	err = ioutil.WriteFile(cachePath, body, 0644)
	if err != nil {
		ch <- fmt.Sprintf("Error: write cache %s: %v", cachePath, err)
		return
	}

	ch <- fmt.Sprintf("cache: %s\ttime: %.2fs\tbody size: %7d\t%s", cacheStatus, secs, len(body), url)
}

func sanitiseForFSPath(str string) string {
	return strings.ReplaceAll(str, "/", "_")
}
