// page 83
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", checkAnagram(os.Args[1], os.Args[2]))
}

func checkAnagram(s1, s2 string) bool {
	runeCountMap := map[rune]int{}

	// count up by s1 runes
	for _, rune := range []rune(s1) {
		runeCountMap[rune]++
	}
	// count down by s2 runes
	for _, rune := range []rune(s2) {
		runeCountMap[rune]--
	}
	for _, cnt := range runeCountMap {
		if cnt != 0 {
			return false
		}
	}

	return true
}
