package popcount

import "testing"

func TestOrigPopCount(t *testing.T) {
	if OrigPopCount(0x1234567890ABCDEF) != 32 {
		t.Fail()
	}
}

func TestLoopPopCount(t *testing.T) {
	if LoopPopCount(0x1234567890ABCDEF) != 32 {
		t.Fail()
	}
}

// 0.2682 ns/op
func BenchmarkOrigPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OrigPopCount(0x1234567890ABCDEF)
	}
}

// 4.385 ns/op
func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(0x1234567890ABCDEF)
	}
}
