package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	issueGroup := [3][]*Issue{}
	for _, item := range result.Items {
		if item.CreatedAt.Sub(time.Now().AddDate(0, -1, 0)) > 0 {
			issueGroup[0] = append(issueGroup[0], item)
		}
		if item.CreatedAt.Sub(time.Now().AddDate(0, -1, 0)) > 0 {
			issueGroup[1] = append(issueGroup[1], item)
		}
		if item.CreatedAt.Sub(time.Now().AddDate(-1, 0, 0)) < 0 {
			issueGroup[2] = append(issueGroup[2], item)
		}
	}

	println("\n### within a month")
	for _, item := range issueGroup[0] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	println("\n### within a year")
	for _, item := range issueGroup[1] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	println("\n### older than a year")
	for _, item := range issueGroup[2] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
