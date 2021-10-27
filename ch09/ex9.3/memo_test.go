// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"fmt"
	"memo"
	"sync"
	"testing"
	"time"

	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()

	// 最初はhttps://golang.orgをリクエストするがすぐにキャンセルする
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			cancel := make(chan struct{})
			close(cancel)

			go func() {
				defer wg.Done()
				start := time.Now()

				_, err := m.Get("https://golang.org", cancel)
				if err != memo.ErrCanceled {
					t.Fatal("request was not canceled")
				}

				elapsed := time.Since(start)
				if elapsed.Microseconds() > 1000 {
					t.Fatal("it takes too long to cancel")
				}
				fmt.Printf("canceled: %s\n", elapsed)
			}()
		}()
	}
	wg.Wait()

	// 一度リクエストする
	start := time.Now()
	value, err := m.Get("https://golang.org", make(chan struct{}))
	if err != nil {
		t.Fatal(err)
		return
	}
	elapsed := time.Since(start)
	if elapsed.Microseconds() < 1000 {
		t.Fatal("seem to hit the cache")
	}
	fmt.Printf("fetched: %d bytes, %s\n", len(value.([]byte)), elapsed)

	// 以後はキャッシュにヒットするので短時間で終了するはず
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()

			value, err := m.Get("https://golang.org", make(chan struct{}))
			if err != nil {
				t.Fatal(err)
				return
			}

			elapsed := time.Since(start)
			if elapsed.Microseconds() > 1000 {
				t.Fatal("doesn't seem to hit the cache")
			}
			fmt.Printf("cache hit:\t %d bytes, %s\n", len(value.([]byte)), elapsed)
		}()
	}
	wg.Wait()
}
