package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"ex4.14/github"
)

type templateData struct {
	RepoName   string
	Bugs       *github.IssuesSearchResult
	Milestones *github.MilestonesSearchResult
}

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("usage: %s REPO_NAME\n", os.Args[0]))
	}
	repoName := os.Args[1]

	templateHTML, _ := os.ReadFile("assets/template.html")
	var parsedTemplate = template.Must(template.New("template").Funcs(template.FuncMap{"progress": func(open, close int) string {
		return fmt.Sprint(float32(close)/float32(open+close)*100) + "%"
	}}).Parse(string(templateHTML)))

	// fetch github api
	bugs, err := github.FetchBugs(repoName)
	if err != nil {
		panic(err)
	}
	milestones, err := github.FetchMilestones(repoName)
	if err != nil {
		panic(err)
	}

	// start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = parsedTemplate.Execute(w, templateData{
			RepoName:   repoName,
			Bugs:       bugs,
			Milestones: milestones,
		})
		if err != nil {
			panic(err)
		}
	})
	println("server ready on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
