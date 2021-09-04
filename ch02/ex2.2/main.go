// page 25
package main

import (
	"fmt"
	"os"
	"strconv"

	"ex2.2/lengthconv"
	"ex2.2/tempconv"
	"ex2.2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		ft := lengthconv.Feet(t)
		m := lengthconv.Metre(t)
		p := weightconv.Pound(t)
		k := weightconv.Kilogram(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n",
			ft, lengthconv.FToM(ft), m, lengthconv.MToF(m))
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}
