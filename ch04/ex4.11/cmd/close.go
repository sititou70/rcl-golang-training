package cmd

import (
	"githubissue/github"
	"githubissue/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func CmdClose(cmd *cobra.Command, args []string) {
	token, err := util.GetToken()
	if err != nil {
		color.Red(`API Token is required for changing issues.
		Use 'setToken' sub-command for setting token.`)
		panic(err)
	}

	repo, err := util.PromptRepository()
	if err != nil {
		panic(err)
	}

	issue, err := util.PromptIssue(repo.FullName)
	if err != nil {
		panic(err)
	}

	confirm := util.Confirm("Close issue?")
	if !confirm {
		color.Green("aborted")
		return
	}

	err = github.UpdateIssue(repo.FullName, issue.Number, github.Issue{
		Title: issue.Title,
		Body:  issue.Body,
		State: "closed",
	}, token)
	if err != nil {
		color.Red("issueclose failed")
		panic(err)
	}

	color.Green("Issue closed\n")
}
