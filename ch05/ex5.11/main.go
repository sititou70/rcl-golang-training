// page 159
package main

import (
	"fmt"
	"sort"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	// ex5.11 added
	"linear algebra": {"calculus"},
}

func main() {
	result, err := topoSort(prereqs)
	if err != nil {
		panic(err)
	}

	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string, path []string) error
	visitAll = func(items []string, path []string) error {
		for _, item := range items {
			// check circular reference
			for i, pathItem := range path {
				if pathItem == item {
					return fmt.Errorf("circular reference found: %v", strings.Join(append(path[i:], item), " -[refers to]-> "))
				}
			}

			// visit next node
			if !seen[item] {
				seen[item] = true
				err := visitAll(m[item], append(path, item))
				if err != nil {
					return err
				}
				order = append(order, item)
			}
		}

		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	err := visitAll(keys, nil)

	return order, err
}
