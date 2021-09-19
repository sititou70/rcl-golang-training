package github

import "time"

const APIBaseURL = "https://api.github.com"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

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

type Repository struct {
	Name     string
	FullName string `json:"full_name"`
}

type RepositorySearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Repository
}
