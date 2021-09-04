package popcount

import "testing"

func TestOrigPopCount(t *testing.T) {
	if OrigPopCount(0x1234567890ABCDEF) != 32 {
		t.Fail()
	}
}

func TestClearPopCount(t *testing.T) {
	if ClearPopCount(0x1234567890ABCDEF) != 32 {
		t.Fail()
	}
}

// 0.2682 ns/op
func BenchmarkOrigPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OrigPopCount(0x1234567890ABCDEF)
	}
}

// 11.94 ns/op
func BenchmarkClearPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(0x1234567890ABCDEF)
	}
}
