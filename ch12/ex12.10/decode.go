// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 344.

// Package sexpr provides a means for converting Go objects to and
// from S-expressions.
package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

//!+Unmarshal
// Unmarshal parses S-expression data and populates the variable
// whose address is in the non-nil pointer out.
func Unmarshal(data []byte, out interface{}, typeMap TypeMap) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem(), typeMap)
	return nil
}

//!-Unmarshal

type Token interface{}
type Symbol struct {
	Name string
}
type String struct {
	Value string
}
type Int struct {
	Value int
}
type StartList struct{}
type EndList struct{}

type Decoder struct {
	lex *lexer
}

type TypeMap map[string]reflect.Type

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{
		lex: &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}},
	}
	d.lex.scan.Init(r)
	return d
}

func (d *Decoder) Decode(v interface{}, typeMap TypeMap) (err error) {
	d.lex.next()
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.lex.scan.Position, x)
		}
	}()

	read(d.lex, reflect.ValueOf(v).Elem(), typeMap)
	return nil
}

func (d *Decoder) Token() (token Token, err error) {
	d.lex.next()
	switch d.lex.token {
	case scanner.Ident:
		return Symbol{d.lex.text()}, nil
	case scanner.String:
		s, _ := strconv.Unquote(d.lex.text()) // NOTE: ignoring errors
		return String{s}, nil
	case scanner.Int:
		i, _ := strconv.Atoi(d.lex.text()) // NOTE: ignoring errors
		return Int{i}, nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("Token: unknown token type: %v", d.lex.token)
	}
}

//!+lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

//!-lexer

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
// - that v is always a variable of the appropriate type for the
//   S-expression value.  For example, v must not be a boolean,
//   interface, channel, or function, and if v is an array, the input
//   must have the correct number of elements.
// - that v in the top-level call to read has the zero value of its
//   type and doesn't need clearing.
// - that if v is a numeric variable, it is a signed integer.

//!+read
func read(lex *lexer, v reflect.Value, typeMap TypeMap) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		switch lex.text() {
		case "nil":
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		case "t":
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64) // NOTE: ignoring errors
		v.SetFloat(f)
		lex.next()
		return
	case '#':
		lex.next()
		switch lex.text() {
		case "C":
			lex.next()
			lex.next()
			real, _ := strconv.ParseFloat(lex.text(), 64)
			lex.next()
			imag, _ := strconv.ParseFloat(lex.text(), 64)
			v.SetComplex(complex(real, imag))
			lex.next()
			lex.next()
			return
		case "Iface":
			lex.next()
			lex.next()
			typeIdent, _ := strconv.Unquote(lex.text())
			t, ok := typeMap[typeIdent]
			if !ok {
				panic(fmt.Errorf("Iface: %v not in TypeMap", typeIdent))
			}
			val := reflect.New(t).Elem()
			lex.next()
			read(lex, val, typeMap)
			v.Set(val)
			lex.next()
			return
		}
	case '(':
		lex.next()
		readList(lex, v, typeMap)
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

//!-read

//!+readlist
func readList(lex *lexer, v reflect.Value, typeMap TypeMap) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i), typeMap)
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item, typeMap)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name), typeMap)
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key, typeMap)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value, typeMap)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

//!-readlist
