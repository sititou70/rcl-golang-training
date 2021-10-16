package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"text/tabwriter"
)

type ClockData struct {
	Name string
	Data string
}

func main() {
	if len(os.Args) <= 1 {
		panic(fmt.Sprintf("usage: %s [NAME=HOST:PORT]...\ne.g. %s NewYork=localhost:8010 Tokyo=localhost:8020", os.Args[0], os.Args[0]))
	}

	latestClocks := make([]ClockData, len(os.Args)-1)
	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i, arg := range os.Args[1:] {
		// parse args
		s := strings.Split(arg, "=")
		if len(s) != 2 {
			panic(fmt.Sprintf("invalid argument format: %s", arg))
		}

		wg.Add(1)
		go func(i int) {
			// open connection
			conn, err := net.Dial("tcp", s[1])
			if err != nil {
				panic(err)
			}
			defer conn.Close()

			// read and output
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				mu.Lock()
				latestClocks[i] = ClockData{
					Name: s[0],
					Data: scanner.Text(),
				}
				printClocks(latestClocks)
				println()
				mu.Unlock()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func printClocks(clocks []ClockData) {
	names := []string{}
	border := []string{}
	data := []string{}
	for _, c := range clocks {
		names = append(names, c.Name)
		border = append(border, strings.Repeat("-", len(c.Name)))
		data = append(data, c.Data)
	}

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, strings.Join(names, "\t")+"\n")
	fmt.Fprintf(tw, strings.Join(border, "\t")+"\n")
	fmt.Fprintf(tw, strings.Join(data, "\t")+"\n")
	tw.Flush()
}
