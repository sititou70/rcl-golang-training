package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SearchRepositories(query []string, apiToken string) (*RepositorySearchResult, error) {
	q := url.QueryEscape(strings.Join(query, " "))
	req, _ := http.NewRequest("GET", APIBaseURL+"/search/repositories?q="+q, nil)
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

	var result RepositorySearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
