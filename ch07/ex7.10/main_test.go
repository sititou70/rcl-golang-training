package main

import (
	"sort"
	"testing"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}

		i++
		j--
	}

	return true
}

type RuneSlice []rune

func (s RuneSlice) Len() int {
	return len(s)
}
func (s RuneSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s RuneSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		data   string
		result bool
	}{
		{"しんぶんし", true},
		{"新聞紙", false},
		{"なかきよのとおのねふりのみなめさめなみのりふねのおとのよきかな", true},
		{"長き夜の遠の睡りの皆目醒め波乗り船の音の良きかな", false},
	}

	for _, test := range tests {
		if IsPalindrome(RuneSlice([]rune(test.data))) != test.result {
			t.Fail()
		}
	}
}
