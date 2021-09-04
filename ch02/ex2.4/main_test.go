package popcount

import "testing"

func TestOrigPopCount(t *testing.T) {
	if OrigPopCount(0x1234567890ABCDEF) != 32 {
		t.Fail()
	}
}

func TestShiftPopCount(t *testing.T) {
	if ShiftPopCount(0x1234567890ABCDEF) != 32 {
		t.Fail()
	}
}

// 0.2682 ns/op
func BenchmarkOrigPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OrigPopCount(0x1234567890ABCDEF)
	}
}

// 17.44 ns/op
func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(0x1234567890ABCDEF)
	}
}
