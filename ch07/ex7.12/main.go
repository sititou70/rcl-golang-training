// page 226
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	println("server launch on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

var mu sync.Mutex

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	templateHTML, _ := os.ReadFile("assets/template.html")
	var parsedTemplate = template.Must(template.New("template").Parse(string(templateHTML)))

	type Entry struct {
		Name  string
		Price int
	}
	data := []Entry{}
	mu.Lock()
	for item, price := range db {
		data = append(data, Entry{item, price})
	}
	mu.Unlock()

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	mu.Lock()
	price, ok := db[item]
	mu.Unlock()

	if ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item parameter is required\n")
	}
	priceStr := req.URL.Query().Get("price")
	if priceStr == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price parameter is required\n")
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price parameter must be number\n")
	}

	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s is already exist\n", item)
	}

	mu.Lock()
	db[item] = price
	mu.Unlock()
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item parameter is required\n")
	}
	priceStr := req.URL.Query().Get("price")
	if priceStr == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price parameter is required\n")
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price parameter must be number\n")
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s is not found\n", item)
	}

	mu.Lock()
	db[item] = price
	mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item parameter is required\n")
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s is not found\n", item)
	}

	mu.Lock()
	delete(db, item)
	mu.Unlock()
}
