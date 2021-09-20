package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

func CmdDownload(cmd *cobra.Command, args []string) {
	EnsureCacheDir()

	concurrencyCh := make(chan struct{}, 10)
	top, err := fetchXkcdJSON("")
	if err != nil {
		fmt.Printf("CmdDownload: fetch top json failed: %v\n", err)
	}

	// download
	var wg sync.WaitGroup
	for i := 1; i <= top.Num; i++ {
		wg.Add(1)
		concurrencyCh <- struct{}{}

		go func(i int) {
			defer func() {
				wg.Done()
				<-concurrencyCh
			}()

			_, err := loadXkcdJSONCache(strconv.Itoa(i))
			if err == nil {
				fmt.Printf("cache exist: %d\n", i)
				return
			}

			fmt.Printf("Downloading: %d\n", i)
			res, err := fetchXkcdJSON(strconv.Itoa(i))
			if err != nil {
				fmt.Printf("download: fetch failed: %d, %v\n", i, err)
			}

			json, _ := json.Marshal(res)
			err = os.WriteFile(GetCachePath(strconv.Itoa(res.Num)), json, 0644)
			if err != nil {
				fmt.Printf("download: write file failed: %d, %v\n", i, err)
			}
		}(i)
	}
	wg.Wait()

	println("download done.")
}

func loadXkcdJSONCache(id string) (XkcdJSON, error) {
	cache, err := os.ReadFile(GetCachePath(id))
	if err != nil {
		return XkcdJSON{}, err
	}

	var xkcdJSON XkcdJSON
	err = json.Unmarshal(cache, &xkcdJSON)
	if err != nil {
		return XkcdJSON{}, err
	}

	return xkcdJSON, nil
}

func fetchXkcdJSON(id string) (XkcdJSON, error) {
	var xkcdJSON XkcdJSON
	cache, err := os.ReadFile(GetCachePath(id))
	if err == nil {
		err = json.Unmarshal(cache, &xkcdJSON)
		if err == nil {
			return xkcdJSON, nil
		}
	}

	// fetch
	if id != "" {
		id = "/" + id
	}
	resp, err := http.Get(domain + id + "/info.0.json")
	if err != nil {
		fmt.Printf("fetchXkcdJSON: %s, %v\n", id, err)
		return XkcdJSON{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return XkcdJSON{}, err
	}

	err = json.Unmarshal(body, &xkcdJSON)
	if err != nil {
		return XkcdJSON{}, err
	}

	return xkcdJSON, nil
}

func GetCachePath(id string) string {
	return CacheDir + "/" + id + "_.json"
}

func EnsureCacheDir() {
	os.Mkdir(CacheDir, 0775)
}
