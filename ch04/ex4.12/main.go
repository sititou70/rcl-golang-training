package main

import (
	"os"
	"xkcd/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var cmdDownload = &cobra.Command{
		Use:   "download",
		Short: "download comic information",
		Long:  `download comic information`,
		Run:   cmd.CmdDownload,
	}
	var cmdSearch = &cobra.Command{
		Use:   "search REGEXP",
		Short: "search comic",
		Long:  `search comic transcript`,
		Run:   cmd.CmdSearch,
		Args:  cobra.MinimumNArgs(1),
	}

	var rootCmd = &cobra.Command{Use: os.Args[0]}
	rootCmd.AddCommand(cmdDownload, cmdSearch)
	rootCmd.Execute()
}
