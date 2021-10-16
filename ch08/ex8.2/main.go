package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"ex8.2/ftp"
)

const PORT = "2121"
const DEBUG = true

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ftp server is ready on %s\n", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go ftp.HandleFTP(conn)
	}
}
