// page 200
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var root *tree
	for i := 0; i < 50; i++ {
		root = add(root, rand.Int()%50)
	}
	println(root.String())
}

func (t *tree) String() string {
	var stringifyNode func(t *tree, b *strings.Builder, firstPrefix string, prefix string)
	stringifyNode = func(t *tree, b *strings.Builder, firstPrefix string, prefix string) {
		if t == nil {
			return
		}
		fmt.Fprintf(b, "%s%d\n", firstPrefix, t.value)

		if t.right != nil {
			stringifyNode(t.left, b, prefix+"├─L─", prefix+"│   ")
		} else {
			stringifyNode(t.left, b, prefix+"└─L─", prefix+"    ")
		}
		stringifyNode(t.right, b, prefix+"└─R─", prefix+"    ")
	}

	var b strings.Builder
	stringifyNode(t, &b, "", "")

	return b.String()
}

// tree
type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
