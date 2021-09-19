// page 112
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func main() {
	categories := map[string]int{}
	scripts := map[string]int{}
	properties := map[string]int{}
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// countup
		//// categories
		for name, rangeTable := range unicode.Categories {
			if unicode.In(r, rangeTable) {
				categories[name]++
			}
		}
		//// scripts
		for name, rangeTable := range unicode.Scripts {
			if unicode.In(r, rangeTable) {
				scripts[name]++
			}
		}
		//// properties
		for name, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				properties[name]++
			}
		}
		//// utflen
		utflen[n]++
	}

	// print result
	printTitle("categories count")
	printMap(categories)
	printTitle("scripts count")
	printMap(scripts)
	printTitle("properties count")
	printMap(properties)
	printTitle("len count")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func printTitle(title string) {
	fmt.Printf("\n##### %s #####\n", title)
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
