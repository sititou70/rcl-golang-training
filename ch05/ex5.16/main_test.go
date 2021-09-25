// page 163
package main

import (
	"strings"
	"testing"
)

func join(sep string, strs ...string) string {
	b := strings.Builder{}

	for i, str := range strs {
		if i != 0 {
			b.Write([]byte(sep))
		}
		b.Write([]byte(str))
	}

	return b.String()
}

func TestJoin(t *testing.T) {
	if join(" | ", "cat foo.txt", "sort", "uniq", "less") != "cat foo.txt | sort | uniq | less" {
		t.Fail()
	}
	if join(" | ", "cat foo.txt") != "cat foo.txt" {
		t.Fail()
	}
}
