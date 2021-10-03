// See page 209
package main

import (
	"flag"
	"fmt"

	"ex7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
