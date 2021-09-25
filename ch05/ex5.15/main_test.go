// page 163
package main

import (
	"testing"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func min(values ...int) int {
	min := MaxInt
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min
}
func min2(value int, values ...int) int {
	min := value
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min
}

func max(values ...int) int {
	max := MinInt
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return max
}
func max2(value int, values ...int) int {
	max := value
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return max
}

func TestMin(t *testing.T) {
	if min(1, 2, 3, 4) != 1 {
		t.Fail()
	}
	if min(-3, 1, 23, 56) != -3 {
		t.Fail()
	}
	if min() != MaxInt {
		t.Fail()
	}
	if min2(1, 2, 3, 4) != 1 {
		t.Fail()
	}
	if min2(-3, 1, 23, 56) != -3 {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	if max(1, 2, 3, 4) != 4 {
		t.Fail()
	}
	if max(-3, 1, 23, 56) != 56 {
		t.Fail()
	}
	if max() != MinInt {
		t.Fail()
	}
	if max2(1, 2, 3, 4) != 4 {
		t.Fail()
	}
	if max2(-3, 1, 23, 56) != 56 {
		t.Fail()
	}
}
