// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	var i interface{} = []int{1, 2, 3}
	tests := []struct {
		input interface{}
		want  string
	}{
		{true, "t"},
		{false, "nil"},
		{3.14159, "3.14159"},
		{1.23 + 4.56i, "#C(1.23 4.56)"},
		{make(chan int), `#Chan(":chan int")`},
		{func(n int) int { return n + 1 }, `#Func(":func(int) int")`},
		{&i, `#Iface("[]int" (1 2 3))`},
	}

	for _, test := range tests {
		data, err := Marshal(test.input)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if string(data) != test.want {
			t.Fatalf("Marshal(%v) = %v, want %v", test.input, string(data), test.want)
		}
	}
}
