package intset

import "testing"

type IntSetMap map[int]bool

func TestIntset(t *testing.T) {
	var sx, sy IntSet
	sx.Add(1)
	sx.Add(144)
	sx.Add(9)
	sy.Add(9)
	sy.Add(42)
	sx.UnionWith(&sy)

	mx, my := IntSetMap{}, IntSetMap{}
	mx[1] = true
	mx[144] = true
	mx[9] = true
	my[9] = true
	my[42] = true
	for k := range my {
		mx[k] = true
	}

	for i := 0; i < 256; i++ {
		if sx.Has(i) != mx[i] {
			t.Errorf("sx.Has(%v) == %v, but mx[%v] == %v", i, sx.Has(i), i, mx[i])
		}
	}
}
