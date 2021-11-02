package intset

import (
	"math/rand"
	"testing"
	"time"
)

type IntSetMap map[int]bool

const RAND_MAX = 10000
const ADD_NUM = 1000
const UNION_NUM = 1000

func BenchmarkIntSet(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		var x, y IntSet
		for j := 0; j < ADD_NUM; j++ {
			x.Add(rand.Intn(RAND_MAX))
		}
		for j := 0; j < ADD_NUM; j++ {
			y.Add(rand.Intn(RAND_MAX))
		}
		for j := 0; j < UNION_NUM; j++ {
			x.UnionWith(&y)
		}
	}
}
func BenchmarkIntSet32(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		var x, y IntSet32
		for j := 0; j < ADD_NUM; j++ {
			x.Add(rand.Intn(RAND_MAX))
		}
		for j := 0; j < ADD_NUM; j++ {
			y.Add(rand.Intn(RAND_MAX))
		}
		for j := 0; j < UNION_NUM; j++ {
			x.UnionWith(&y)
		}
	}
}

func BenchmarkIntSetMap(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		x, y := IntSetMap{}, IntSetMap{}
		for j := 0; j < ADD_NUM; j++ {
			x[rand.Intn(RAND_MAX)] = true
		}
		for j := 0; j < ADD_NUM; j++ {
			y[rand.Intn(RAND_MAX)] = true
		}
		for j := 0; j < UNION_NUM; j++ {
			for k := range y {
				x[k] = true
			}
		}
	}
}
