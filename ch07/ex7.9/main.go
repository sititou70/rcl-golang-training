// See page 216
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"ex7.9/track"
)

var tracks = []*track.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	templateHTML, _ := os.ReadFile("assets/template.html")
	var parsedTemplate = template.Must(template.New("template").Parse(string(templateHTML)))

	// start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		keysParam := r.URL.Query().Get("keys")
		keys := []string{}
		if keysParam != "" {
			err := json.NewDecoder(strings.NewReader(keysParam)).Decode(&keys)
			if err != nil {
				panic(err)
			}
		}

		m := track.MultiKeySort{Track: tracks, Keys: keys}
		sort.Sort(&m)

		err := parsedTemplate.Execute(w, m.Track)
		if err != nil {
			panic(err)
		}
	})
	println("server ready on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
