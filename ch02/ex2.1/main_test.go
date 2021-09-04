// page 47
package main

import (
	"testing"

	"ex2.1/tempconv"
)

func TestKToC(t *testing.T) {
	if tempconv.KToC(0) != -273.15 {
		t.Fail()
	}
}

func TestCToK(t *testing.T) {
	if tempconv.CToK(0) != 273.15 {
		t.Fail()
	}
}

func TestKToF(t *testing.T) {
	if tempconv.KToF(273.15) != 32 {
		t.Fail()
	}
}

func TestFToK(t *testing.T) {
	if tempconv.FToK(32) != 273.15 {
		t.Fail()
	}
}
