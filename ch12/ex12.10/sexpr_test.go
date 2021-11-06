package sexpr

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Input struct {
		True      bool
		False     bool
		Float     float64
		Complex   complex128
		Interface interface{}
	}
	input := Input{
		True:      true,
		False:     false,
		Float:     3.14159,
		Complex:   1.23 + 4.56i,
		Interface: []int{1, 2, 3},
	}

	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var actual Input
	err = Unmarshal(data, &actual, TypeMap{"[]int": reflect.TypeOf([]int{})})
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(input, actual) {
		t.Fatal("not equal")
	}
}
