package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func FetchBugs(repository string) (*IssuesSearchResult, error) {
	q := url.QueryEscape("bug repo:" + repository)

	req, _ := http.NewRequest("GET", APIBaseURL+"/search/issues?q="+q, nil)
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
