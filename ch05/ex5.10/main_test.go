// page 159
package main

import (
	"testing"
)

func TestToposort(t *testing.T) {
	completed := map[string]bool{}
	for _, course := range topoSort(prereqs) {
		for prereq := range prereqs[course] {
			if !completed[prereq] {
				t.Fatalf("you are trying to take %s, but have not completed %s", course, prereq)
			}
		}

		completed[course] = true
	}
}
