package main

import (
	"fmt"
	"strings"
	"testing"
)

var testArgs []string = []string{"Lorem", "Ipsum", "is", "simply", "dummy", "text", "of", "the", "printing", "and", "typesetting", "industry.", "Lorem", "Ipsum", "has", "been", "the", "industry's", "standard", "dummy", "text", "ever", "since", "the", "1500s,", "when", "an", "unknown", "printer", "took", "a", "galley", "of", "type", "and", "scrambled", "it", "to", "make", "a", "type", "specimen", "book.", "It", "has", "survived", "not", "only", "five", "centuries,", "but", "also", "the", "leap", "into", "electronic", "typesetting,", "remaining", "essentially", "unchanged.", "It", "was", "popularised", "in", "the", "1960s", "with", "the", "release", "of", "Letraset", "sheets", "containing", "Lorem", "Ipsum", "passages,", "and", "more", "recently", "with", "desktop", "publishing", "software", "like", "Aldus", "PageMaker", "including", "versions", "of", "Lorem", "Ipsum."}

func loopEcho(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}
func BenchmarkLoopEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopEcho(testArgs)
	}
}

func joinEcho(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}
func BenchmarkJoinEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinEcho(testArgs)
	}
}
