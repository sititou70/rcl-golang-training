package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	responce := make(chan *http.Response)
	cancel := make(chan struct{})

	wg := sync.WaitGroup{}
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, responce, cancel, &wg)
	}

	resp := <-responce
	close(cancel)
	go func() {
		wg.Wait()
		close(responce)
	}()
	for range responce {
	}

	_, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}

func fetch(url string, responce chan<- *http.Response, cancel <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintln(os.Stdout, "fetch:", err)
		return
	}
	requestCanceled := make(chan struct{})
	requestDone := make(chan struct{})
	req.Cancel = requestCanceled

	go func() {
		select {
		case <-cancel:
			close(requestCanceled)
		case <-requestDone:
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	close(requestDone)
	if err != nil {
		fmt.Fprintln(os.Stdout, "fetch:", err)
		return
	}

	responce <- resp
}
