package main

import (
	"testing"
)

func noReturn() (str string) {
	defer func() {
		str = recover().(string)
	}()
	panic("non-zero value")
}

func TestNoReturn(t *testing.T) {
	if noReturn() != "non-zero value" {
		t.Fail()
	}
}
