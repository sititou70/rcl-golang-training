package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	// read stdin
	buf := []byte{}
	_, err := os.Stdin.Read(buf)
	if err != nil {
		panic(err)
	}

	// print
	format := "%x\n"
	mode := "256"
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}
	switch mode {
	case "384":
		fmt.Printf(format, sha512.Sum384(buf))
	case "512":
		fmt.Printf(format, sha512.Sum512(buf))
	default:
		fmt.Printf(format, sha256.Sum256(buf))
	}
}
