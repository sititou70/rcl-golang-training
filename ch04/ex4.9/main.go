package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	for s.Scan() {
		counts[string(s.Bytes())]++
	}

	printMap(counts)
}

type countMapEntry struct {
	name  string
	count int
}

func printMap(m map[string]int) {
	s := make([]countMapEntry, len(m))
	for name, count := range m {
		s = append(s, countMapEntry{name, count})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].count > s[j].count
	})

	for _, ent := range s {
		if ent.count != 0 {
			fmt.Printf("%-30s\t%d\n", ent.name, ent.count)
		}
	}
}
