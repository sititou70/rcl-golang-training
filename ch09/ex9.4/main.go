package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("usage: ./%s goroutine_num", os.Args[0]))
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	inputCh := make(chan int)
	prevOutCh := inputCh
	for i := 0; i < num; i++ {
		outCh := make(chan int)
		go func(in, out chan int) {
			out <- <-in
		}(prevOutCh, outCh)
		prevOutCh = outCh
	}

	start := time.Now()
	inputCh <- 123
	<-prevOutCh
	fmt.Printf("done: %v\n", time.Since(start))
}
