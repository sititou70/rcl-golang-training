package main

import (
	"reflect"
	"testing"
	"unsafe"
)

type SeenSet map[unsafe.Pointer]bool

func isCycled(v interface{}) bool {
	return walk(reflect.ValueOf(v), SeenSet{})
}

// return cycled(true) or uncycled(false)
func walk(v reflect.Value, seen SeenSet) bool {
	if v.CanAddr() {
		if seen[unsafe.Pointer(v.UnsafeAddr())] {
			return true
		}
		seen[unsafe.Pointer(v.UnsafeAddr())] = true
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return walk(v.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if walk(v.Index(i), seen) {
				return true
			}
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if walk(v.Field(i), seen) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if walk(k, seen) {
				return true
			}
			if walk(v.MapIndex(k), seen) {
				return true
			}
		}
	}

	return false
}

func TestIsCycled(t *testing.T) {
	type Array [1]*Array
	a1 := Array{}
	a1[0] = &a1
	a2 := Array{}
	a3 := Array{}
	a2[0] = &a3
	a3[0] = &a2
	a4 := Array{&a2}

	type Slice []Slice
	s1, s2 := make(Slice, 1), make(Slice, 1)
	s1[0] = s2
	s2[0] = s1
	s3 := Slice{s1}
	s4 := Slice{}
	s5 := Slice{s4}

	type Map map[string]*Map
	m1 := Map{}
	m1["m1"] = &m1
	type Map2 map[*Map2]bool
	m2 := Map2{}
	m2[&m2] = true

	type Link struct {
		next *Link
	}
	l1, l2, l3 := Link{}, Link{}, Link{}
	l2.next, l3.next, l1.next = &l1, &l2, &l3

	tests := []struct {
		title  string
		input  interface{}
		result bool
	}{
		{"a1", a1, true},
		{"a2", a2, true},
		{"a3", a3, true},
		{"a4", a4, true},
		{"s1", s1, true},
		{"s2", s2, true},
		{"s3", s3, true},
		{"s4", s4, false},
		{"s5", s5, false},
		//{"m1", m1, true},
		{"m2", m2, true},
		{"l1", l1, true},
		{"l2", l2, true},
		{"l3", l3, true},
	}

	for _, test := range tests {
		if isCycled(test.input) != test.result {
			t.Errorf("isCycled(%s) != %t",
				test.title, test.result)
		}
	}
}
