package util

import (
	"fmt"
	"githubissue/github"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func PromptIssue(repository string) (*github.Issue, error) {
	var err error
	var issues *github.IssuesSearchResult
	query := ""
	for {
		prompt := promptui.Prompt{
			Label:   "Search Issue",
			Default: query,
		}
		query, err = prompt.Run()
		if err != nil {
			return &github.Issue{}, fmt.Errorf("prompt failed, %v", err)
		}
		token, _ := GetToken()
		issues, err = github.SearchIssues([]string{query}, repository, token)
		if err != nil {
			return &github.Issue{}, fmt.Errorf("search failed, %v", err)
		}

		if issues.TotalCount >= 1 {
			break
		}

		fmt.Printf("Issue not found\n")
	}

	index := 0
	if issues.TotalCount >= 2 {
		selectKeys := []string{}
		issuesMap := map[string]*github.Issue{}
		for _, issue := range issues.Items {
			issuesMap[issue.Title] = issue
			selectKeys = append(selectKeys, issue.Title)
		}
		prompt := promptui.Select{
			Label: "Select Issue",
			Items: selectKeys,
		}
		index, _, err = prompt.Run()
		if err != nil {
			return &github.Issue{}, fmt.Errorf("select failed, %v", err)
		}
	}

	color.Green("selected: %s\n", issues.Items[index].Title)
	return issues.Items[index], nil
}
