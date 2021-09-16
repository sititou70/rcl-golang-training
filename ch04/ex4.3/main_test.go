// page 104
package main

import "testing"

func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestReverse(t *testing.T) {
	test := [5]int{1, 2, 3, 4, 5}
	reverse(&test)
	if test != [5]int{5, 4, 3, 2, 1} {
		t.Fail()
	}
}
