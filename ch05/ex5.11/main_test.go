// page 159
package main

import (
	"testing"
)

// test data
var validPrereqs = map[string][]string{
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
}
var invalidPrereqs1 = map[string][]string{
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

	// circular reference
	"linear algebra": {"calculus"},
}
var invalidPrereqs2 = map[string][]string{
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

	// circular reference
	"intro to programming": {"operating systems"},
}

// test funcs
func TestValidToposort(t *testing.T) {
	completed := map[string]bool{}
	result, err := topoSort(validPrereqs)
	if err != nil {
		t.Fatal(err)
	}
	for _, course := range result {
		for _, prereq := range validPrereqs[course] {
			if !completed[prereq] {
				t.Fatalf("you are trying to take %s, but have not completed %s", course, prereq)
			}
		}

		completed[course] = true
	}
}
func TestInvalidToposort1(t *testing.T) {
	_, err := topoSort(invalidPrereqs1)
	if err == nil {
		t.Fatal(err)
	}
}

func TestInvalidToposort2(t *testing.T) {
	_, err := topoSort(invalidPrereqs2)
	if err == nil {
		t.Fatal(err)
	}
}
