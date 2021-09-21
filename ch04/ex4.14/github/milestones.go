package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func FetchMilestones(repository string) (*MilestonesSearchResult, error) {
	req, _ := http.NewRequest("GET", APIBaseURL+"/repos/"+repository+"/milestones", nil)
	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result MilestonesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
