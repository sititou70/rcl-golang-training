package main

import (
	"math/big"
	"testing"
)

func TestBigFloatAdd(t *testing.T) {
	f1 := BigFloatComplex{
		r: *big.NewFloat(1.0),
		i: *big.NewFloat(2.0),
	}
	f2 := BigFloatComplex{
		r: *big.NewFloat(3.0),
		i: *big.NewFloat(4.0),
	}

	res := bigFloatAdd(f1, f2)
	if res.r.Cmp(big.NewFloat(4.0)) != 0 {
		t.Fail()
	}
	if res.i.Cmp(big.NewFloat(6.0)) != 0 {
		t.Fail()
	}
}
func TestBigFloatMul(t *testing.T) {
	f1 := BigFloatComplex{
		r: *big.NewFloat(1.0),
		i: *big.NewFloat(2.0),
	}
	f2 := BigFloatComplex{
		r: *big.NewFloat(3.0),
		i: *big.NewFloat(4.0),
	}

	res := bigFloatMul(f1, f2)
	if res.r.Cmp(big.NewFloat(-5.0)) != 0 {
		t.Fail()
	}
	if res.i.Cmp(big.NewFloat(10.0)) != 0 {
		t.Fail()
	}
}
func TestBigFloatAbs(t *testing.T) {
	f1 := BigFloatComplex{
		r: *big.NewFloat(3.0),
		i: *big.NewFloat(4.0),
	}

	res := bigFloatAbs(f1)
	if res.Cmp(big.NewFloat(5.0)) != 0 {
		t.Fail()
	}
}

func TestBigRatAdd(t *testing.T) {
	r1 := BigRatComplex{
		r: *new(big.Rat).SetFloat64(1),
		i: *new(big.Rat).SetFloat64(2),
	}
	r2 := BigRatComplex{
		r: *new(big.Rat).SetFloat64(3),
		i: *new(big.Rat).SetFloat64(4),
	}

	res := bigRatAdd(r1, r2)
	if res.r.Cmp(new(big.Rat).SetFloat64(4)) != 0 {
		t.Fail()
	}
	if res.i.Cmp(new(big.Rat).SetFloat64(6)) != 0 {
		t.Fail()
	}
}
func TestBigRatMul(t *testing.T) {
	r1 := BigRatComplex{
		r: *new(big.Rat).SetFloat64(1),
		i: *new(big.Rat).SetFloat64(2),
	}
	r2 := BigRatComplex{
		r: *new(big.Rat).SetFloat64(3),
		i: *new(big.Rat).SetFloat64(4),
	}

	res := bigRatMul(r1, r2)
	if res.r.Cmp(new(big.Rat).SetFloat64(-5)) != 0 {
		t.Fail()
	}
	if res.i.Cmp(new(big.Rat).SetFloat64(10)) != 0 {
		t.Fail()
	}
}
func TestBigRatAbs(t *testing.T) {
	r1 := BigRatComplex{
		r: *new(big.Rat).SetFloat64(3),
		i: *new(big.Rat).SetFloat64(4),
	}

	if bigRatAbs(r1) != 5.0 {
		t.Fail()
	}
}
