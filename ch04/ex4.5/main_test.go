// page 104
package main

import (
	"testing"
)

func uniq(s []string) []string {
	if len(s) == 0 {
		return s
	}

	uniq := s[:1]
	currentString := s[0]
	for _, str := range s {
		if str != currentString {
			uniq = append(uniq, str)
			currentString = str
		}
	}
	return uniq
}

func stringsEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestUniq(t *testing.T) {
	if !stringsEqual(uniq([]string{"1", "1", "2", "2", "3", "3", "3", "4", "4"}), []string{"1", "2", "3", "4"}) {
		t.Fail()
	}
	if !stringsEqual(uniq([]string{}), []string{}) {
		t.Fail()
	}
	if !stringsEqual(uniq([]string{"1"}), []string{"1"}) {
		t.Fail()
	}
}
