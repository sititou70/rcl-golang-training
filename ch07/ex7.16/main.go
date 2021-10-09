package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ex7.13/eval"
)

// http://localhost:8000/eval?q={%22expr%22%3A%22((4+*+(x+-+pow(x%2C3)%2F3+%2B+pow(x%2C5)%2F5+-+pow(x%2C7)%2F7+%2B+pow(x%2C9)%2F9))+-+(y+-+pow(y%2C3)%2F3+%2B+pow(y%2C5)%2F5+-+pow(y%2C7)%2F7+%2B+pow(y%2C9)%2F9))*4%22%2C%22env%22%3A{%22x%22%3A%220.2%22%2C%22y%22%3A%220.0041841004184100415%22}}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.html")
	})
	http.HandleFunc("/eval", func(w http.ResponseWriter, r *http.Request) {
		type Query struct {
			Expr string            `json:"expr"`
			Env  map[string]string `json:"env"`
		}
		queryStr := r.URL.Query().Get("q")
		var query Query
		err := json.Unmarshal([]byte(queryStr), &query)
		if err != nil {
			fmt.Fprintf(w, "query parse error: %v", err)
			return
		}

		expr, err := eval.Parse(query.Expr)
		if err != nil {
			fmt.Fprintf(w, "expression parse error: %v", err)
			return
		}

		env := eval.Env{}
		vars := map[eval.Var]bool{}
		err = expr.Check(vars)
		if err != nil {
			fmt.Fprintf(w, "expression check error: %v", err)
			return
		}
		for key := range vars {
			val, ok := query.Env[string(key)]
			if !ok {
				fmt.Fprintf(w, "undefined variable: %v", key)
				return
			}

			float, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Fprintf(w, "unexpected value of %v=%v: %v", key, val, err)
				return
			}

			env[key] = float
		}

		fmt.Fprintf(w, "answer: %v", expr.Eval(env))
	})

	println("server ready on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
