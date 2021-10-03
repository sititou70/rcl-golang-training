package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}

	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		*c++
	}

	return len(p), nil
}

func TestWordCounter(t *testing.T) {
	data, _ := os.ReadFile("main_test.go")
	var c WordCounter
	fmt.Fprintf(&c, "%s", data)

	if c != 108 {
		t.Fail()
	}
}

func TestLineCounter(t *testing.T) {
	data, _ := os.ReadFile("main_test.go")
	var c LineCounter
	fmt.Fprintf(&c, "%s", data)

	if c != 55 {
		t.Fail()
	}
}
