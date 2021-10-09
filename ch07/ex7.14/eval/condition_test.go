package eval

import "testing"

func TestCondition(t *testing.T) {
	// sqrt(2) > 2 ? 1 : 0
	expr :=
		condition{
			binary{
				'-',
				call{
					"sqrt",
					[]Expr{literal(2)},
				},
				literal(2),
			},
			literal(1),
			literal(0),
		}

	if expr.Eval(Env{}) != 1 {
		t.Fail()
	}
}
