package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	Name    string
	WriteCh chan<- string // an outgoing message channel
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan string) // all incoming client messages
	TIMEOUT  = time.Minute * 5
)

func broadcaster() {
	clients := make(map[*client]bool) // all connected clients

	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.WriteCh <- msg
			}

		case cli := <-entering:
			if len(clients) != 0 {
				names := []string{}
				for c := range clients {
					names = append(names, c.Name)
				}
				cli.WriteCh <- "Here are: " + strings.Join(names, ", ")
			} else {
				cli.WriteCh <- "Nobody here"
			}

			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.WriteCh)
		}
	}
}

func handleConn(conn net.Conn) {
	name := conn.RemoteAddr().String()
	writeCh := make(chan string)
	client := client{
		Name:    name,
		WriteCh: writeCh,
	}
	go clientWriter(conn, writeCh)

	// timeout routine
	lastMessage := time.Now()
	go func() {
		for {
			if time.Since(lastMessage) > TIMEOUT {
				fmt.Fprintln(conn, "Your session has timed out,", TIMEOUT)
				conn.Close()
				break
			}
			time.Sleep(time.Second)
		}
	}()

	// entering process
	client.WriteCh <- "You are " + name
	messages <- name + " has arrived"
	entering <- &client

	// read messages
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- name + ": " + input.Text()
		lastMessage = time.Now()
	}
	// NOTE: ignoring potential errors from input.Err()

	// leaving process
	leaving <- &client
	messages <- name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
