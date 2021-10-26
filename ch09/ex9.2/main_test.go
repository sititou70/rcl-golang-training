package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	if PopCount(0x1234567890ABCDEF) != 32 {
		t.Fail()
	}
}
