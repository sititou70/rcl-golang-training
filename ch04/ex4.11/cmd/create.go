package cmd

import (
	"githubissue/github"
	"githubissue/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func CmdCreate(cmd *cobra.Command, args []string) {
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

	inputIssue, err := util.InputIssue(util.IssueInput{})
	if err != nil {
		panic(err)
	}

	confirm := util.Confirm("Create issue?")
	if !confirm {
		color.Green("aborted")
		return
	}

	err = github.CreateIssue(repo.FullName, github.Issue{
		Title: inputIssue.Title,
		Body:  inputIssue.Body,
	}, token)
	if err != nil {
		color.Red("create issue failed")
		panic(err)
	}

	color.Green("Issue created\n")
}
