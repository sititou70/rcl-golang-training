package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SearchIssues(query []string, repository string, apiToken string) (*IssuesSearchResult, error) {
	if repository != "" {
		query = append(query, " repo:"+repository)
	}
	q := url.QueryEscape(strings.Join(query, " "))

	req, _ := http.NewRequest("GET", APIBaseURL+"/search/issues?q="+q, nil)
	if apiToken != "" {
		req.SetBasicAuth("", apiToken)
	}
	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CreateIssue(repositoryFullName string, issue Issue, apiToken string) error {
	body, _ := json.Marshal(struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{issue.Title, issue.Body})

	req, _ := http.NewRequest("POST", APIBaseURL+"/repos/"+repositoryFullName+"/issues", bytes.NewBuffer(body))
	if apiToken != "" {
		req.SetBasicAuth("", apiToken)
	}

	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return fmt.Errorf("create query failed: %s", resp.Status)
	}

	return nil
}

func UpdateIssue(repositoryFullName string, issueNumber int, issue Issue, apiToken string) error {
	body, _ := json.Marshal(struct {
		Title string `json:"title"`
		Body  string `json:"body"`
		State string `json:"state"`
	}{issue.Title, issue.Body, issue.State})

	req, _ := http.NewRequest("PATCH", APIBaseURL+"/repos/"+repositoryFullName+"/issues/"+fmt.Sprint(issueNumber), bytes.NewBuffer(body))
	if apiToken != "" {
		req.SetBasicAuth("", apiToken)
	}

	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("update query failed: %s", resp.Status)
	}

	return nil
}
