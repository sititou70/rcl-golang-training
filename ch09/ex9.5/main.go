package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	cnt := 0
	go func() {
		for {
			ch <- <-ch
			cnt++
		}
	}()
	go func() {
		for {
			ch <- <-ch
			cnt++
		}
	}()

	start := time.Now()

	ch <- 123
	time.Sleep(time.Second * 3)

	currentCnt := cnt
	elapsed := time.Since(start)
	fmt.Printf(
		"%v communications took place in %v seconds, %f communications/sec",
		currentCnt,
		elapsed,
		float64(currentCnt)/elapsed.Seconds(),
	)
}
