package main

import (
	"reflect"
	"testing"

	"ex7.13/eval"
)

func TestEvalString(t *testing.T) {
	expr1, err := eval.Parse("((4 * (x - pow(x,3)/3 + pow(x,5)/5 - pow(x,7)/7 + pow(x,9)/9)) - (y - pow(y,3)/3 + pow(y,5)/5 - pow(y,7)/7 + pow(y,9)/9))*4")
	if err != nil {
		t.Error(err)
	}

	str1 := expr1.String()
	expr2, err := eval.Parse(str1)
	if err != nil {
		t.Error(err)
	}

	str2 := expr2.String()
	expr3, err := eval.Parse(str2)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expr1, expr2) {
		t.Fail()
	}
	if !reflect.DeepEqual(expr1, expr3) {
		t.Fail()
	}

	env := eval.Env{"x": 1 / 5.0, "y": 1 / 239.0}
	if expr1.Eval(env) != expr2.Eval(env) {
		t.Fail()
	}
	if expr1.Eval(env) != expr3.Eval(env) {
		t.Fail()
	}

	println(expr3.String())
}
