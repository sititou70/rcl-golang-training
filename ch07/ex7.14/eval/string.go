// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strings"
)

//!+Check

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprint(float64(l))
}

func (u unary) String() string {
	if _, ok := u.x.(binary); ok {
		return fmt.Sprintf("%s(%s)", string(u.op), u.x)
	} else {
		return fmt.Sprintf("%s%s", string(u.op), u.x)
	}
}

func (b binary) String() string {
	var x string
	if _, ok := b.x.(binary); ok {
		x = fmt.Sprintf("(%s)", b.x)
	} else {
		x = b.x.String()
	}
	var y string
	if _, ok := b.y.(binary); ok {
		y = fmt.Sprintf("(%s)", b.y)
	} else {
		y = b.y.String()
	}

	return fmt.Sprintf("%s %s %s", x, string(b.op), y)
}

func (c call) String() string {
	argStrs := []string{}
	for _, arg := range c.args {
		argStrs = append(argStrs, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(argStrs, ", "))
}

func (c condition) String() string {
	return fmt.Sprintf("%s ? %s : %s", c.cond, c.true, c.false)
}

//!-Check
