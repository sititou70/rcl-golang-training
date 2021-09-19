package main

import (
	"githubissue/cmd"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cmdSetToken = &cobra.Command{
		Use:   "setToken",
		Short: "set GitHub token",
		Long:  `set GitHub token to secret file.`,
		Run:   cmd.CmdSetToken,
	}
	var cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "create issue",
		Long:  `create GitHub issue.`,
		Run:   cmd.CmdCreate,
	}
	var cmdShow = &cobra.Command{
		Use:   "show",
		Short: "show issue",
		Long:  `show GitHub issue.`,
		Run:   cmd.CmdShow,
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "update issue",
		Long:  `update GitHub issue.`,
		Run:   cmd.CmdUpdate,
	}
	var cmdClose = &cobra.Command{
		Use:   "close",
		Short: "close issue",
		Long:  `close GitHub issue.`,
		Run:   cmd.CmdClose,
	}

	var rootCmd = &cobra.Command{Use: os.Args[0]}
	rootCmd.AddCommand(cmdSetToken, cmdCreate, cmdShow, cmdUpdate, cmdClose)
	rootCmd.Execute()
}
