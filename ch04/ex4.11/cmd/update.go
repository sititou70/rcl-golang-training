package cmd

import (
	"githubissue/github"
	"githubissue/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func CmdUpdate(cmd *cobra.Command, args []string) {
	token, err := util.GetToken()
	if err != nil {
		color.Red(`API Token is required for updating issues.
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

	inputIssue, err := util.InputIssue(util.IssueInput{
		Title: issue.Title,
		Body:  issue.Body,
	})
	if err != nil {
		panic(err)
	}

	confirm := util.Confirm("Update issue?")
	if !confirm {
		color.Green("aborted")
		return
	}

	err = github.UpdateIssue(repo.FullName, issue.Number, github.Issue{
		Title: inputIssue.Title,
		Body:  inputIssue.Body,
		State: issue.State,
	}, token)
	if err != nil {
		color.Red("issue update failed")
		panic(err)
	}

	color.Green("Issue updated\n")
}
