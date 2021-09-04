// page 48
package main

import (
	"testing"

	"ex2.2/lengthconv"
	"ex2.2/tempconv"
	"ex2.2/weightconv"
)

func TestCToF(t *testing.T) {
	if tempconv.CToF(0) != 32 {
		t.Fail()
	}
}

func TestFToC(t *testing.T) {
	if tempconv.FToC(32) != 0 {
		t.Fail()
	}
}

func TestMToF(t *testing.T) {
	if lengthconv.MToF(0.3048) != 1 {
		t.Fail()
	}
}

func TestFToM(t *testing.T) {
	if lengthconv.FToM(1) != 0.3048 {
		t.Fail()
	}
}

func TestKToP(t *testing.T) {
	if weightconv.KToP(0.45359237) != 1 {
		t.Fail()
	}
}

func TestPToK(t *testing.T) {
	if weightconv.PToK(1) != 0.45359237 {
		t.Fail()
	}
}
