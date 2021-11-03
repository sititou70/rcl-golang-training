// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 333.

// Package display provides a means to display structured data.
package display

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//!+Display

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	b := strings.Builder{}
	display(&b, name, reflect.ValueOf(x))
	fmt.Print(b.String())
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(b *strings.Builder, path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(b, "%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(b, fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(b, fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			kb := &strings.Builder{}
			display(kb, "key", key)

			display(b, fmt.Sprintf("%s[%s]", path, strings.Trim(kb.String(), "\n")), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprintf(b, "%s = nil\n", path)
		} else {
			display(b, fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(b, "%s = nil\n", path)
		} else {
			fmt.Fprintf(b, "%s.type = %s\n", path, v.Elem().Type())
			display(b, path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Fprintf(b, "%s = %s\n", path, formatAtom(v))
	}
}

//!-display
