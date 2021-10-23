// See page 274
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c *net.TCPConn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

	wg.Done()
}

//!+
const TIMEOUT = time.Second * 10

func handleConn(c *net.TCPConn) {
	lastScanTime := time.Now()
	go func() {
		for {
			if time.Since(lastScanTime) > TIMEOUT {
				fmt.Fprintln(c, "client timed out in 10 seconds.")
				c.CloseWrite()
				break
			}
			time.Sleep(time.Second * 1)
		}
	}()

	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		lastScanTime = time.Now()
		go echo(c, input.Text(), 1*time.Second, &wg)
	}

	wg.Wait()
	c.CloseWrite()
}

//!-

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
