// page 83
package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	if comma("") != "" {
		t.Fail()
	}
	if comma("1") != "1" {
		t.Fail()
	}
	if comma("123") != "123" {
		t.Fail()
	}
	if comma("1234") != "1,234" {
		t.Fail()
	}
	if comma("12345") != "12,345" {
		t.Fail()
	}
	if comma("123456") != "123,456" {
		t.Fail()
	}
	if comma("1234567") != "1,234,567" {
		t.Fail()
	}
}
