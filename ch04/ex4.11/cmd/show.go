package cmd

import (
	"fmt"
	"githubissue/util"

	"github.com/spf13/cobra"
)

func CmdShow(cmd *cobra.Command, args []string) {
	repo, err := util.PromptRepository()
	if err != nil {
		panic(err)
	}

	issue, err := util.PromptIssue(repo.FullName)
	if err != nil {
		panic(err)
	}

	println()
	fmt.Printf("Title:     %s #%d (%s)\n", issue.Title, issue.Number, issue.HTMLURL)
	fmt.Printf("State:     %s\n", issue.State)
	fmt.Printf("Author:    %s (%s)\n", issue.User.Login, issue.User.HTMLURL)
	fmt.Printf("CreatedAt: %s\n", issue.CreatedAt)
	fmt.Printf("----------\n")
	fmt.Printf("%s\n", issue.Body)
}
