package util

import (
	"fmt"
	"githubissue/github"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func PromptRepository() (*github.Repository, error) {
	var err error
	var repos *github.RepositorySearchResult
	query := ""
	for {
		prompt := promptui.Prompt{
			Label:   "Search Repository",
			Default: query,
		}
		query, err = prompt.Run()
		if err != nil {
			return &github.Repository{}, fmt.Errorf("prompt failed, %v", err)
		}
		token, _ := GetToken()
		repos, err = github.SearchRepositories([]string{query}, token)
		if err != nil {
			return &github.Repository{}, fmt.Errorf("search failed, %v", err)
		}

		if repos.TotalCount >= 1 {
			break
		}

		fmt.Printf("Repository not found\n")
	}

	index := 0
	if repos.TotalCount >= 2 {
		selectKeys := []string{}
		reposMap := map[string]*github.Repository{}
		for _, repo := range repos.Items {
			reposMap[repo.FullName] = repo
			selectKeys = append(selectKeys, repo.FullName)
		}
		prompt := promptui.Select{
			Label: "Select Repository",
			Items: selectKeys,
		}
		index, _, err = prompt.Run()
		if err != nil {
			return &github.Repository{}, fmt.Errorf("select failed, %v", err)
		}
	}

	color.Green("selected: %s\n", repos.Items[index].FullName)
	return repos.Items[index], nil
}
