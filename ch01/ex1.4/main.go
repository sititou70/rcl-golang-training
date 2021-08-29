// page 14
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Set map[string]struct{}
type Count struct {
	count int
	files Set
}
type Counts map[string]*Count

func main() {
	counts := make(Counts)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.3: %v\n", err)
			}

			countLines(f, counts, filename)
			f.Close()
		}
	}

	for line, info := range counts {
		if info.count > 1 {
			fmt.Printf("%s\tcount: %d\tfiles: ", line, info.count)
			for filename, _ := range info.files {
				fmt.Printf("%s ", filename)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts Counts, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &Count{count: 0, files: make(Set)}
		}

		counts[input.Text()].count++
		counts[input.Text()].files[filename] = struct{}{}
	}
}
