package main

import (
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 1 {
		panic("too few arguments")
	}

	expanded := expand(os.Args[1], func(s string) string {
		return os.Getenv(s)
	})
	println(expanded)
}

func expand(s string, f func(string) string) string {
	return regexp.MustCompile(`\$[^ ]+`).ReplaceAllStringFunc(s, func(s string) string {
		key := s[1:]
		return f(key)
	})
}
