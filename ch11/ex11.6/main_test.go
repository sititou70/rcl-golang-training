package main

import "testing"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func OrigPopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func ShiftPopCount(x uint64) int {
	var cnt int

	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			cnt++
		}
		x = x >> 1
	}

	return cnt
}
func ClearPopCount(x uint64) int {
	var cnt int

	for x != 0 {
		x = x & (x - 1)
		cnt++
	}

	return cnt
}
func BenchmarkOrig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OrigPopCount(0x1234567890ABCDEF)
	}
}
func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(0x1234567890ABCDEF)
	}
}
func BenchmarkClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(0x1234567890ABCDEF)
	}
}
