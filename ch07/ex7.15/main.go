package main

import (
	"fmt"
	"os"
	"strconv"

	"ex7.15/eval"
	"github.com/manifoldco/promptui"
)

func main() {
	expr, err := eval.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	env := eval.Env{}
	vars := map[eval.Var]bool{}
	expr.Check(vars)
	for name := range vars {
		validate := func(input string) error {
			env[name], err = strconv.ParseFloat(input, 64)
			return err
		}
		prompt := promptui.Prompt{
			Label:    "value of " + string(name) + " ?",
			Validate: validate,
		}
		prompt.Run()
	}

	fmt.Printf("\nanswer: %v\n", expr.Eval(env))
}
