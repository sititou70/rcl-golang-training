package cmd

import (
	"fmt"
	"githubissue/util"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func CmdSetToken(cmd *cobra.Command, args []string) {
	fmt.Printf(`Enter your GitHub token for the CLI tool.
Token will wrote to the '%s' file.
To create a token: https://github.com/settings/tokens
`, util.TokenFileName)
	prompt := promptui.Prompt{
		Label: "GitHub Token",
		Mask:  '*',
	}

	token, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	confirm := util.Confirm("Set token?")
	if !confirm {
		color.Green("aborted")
		return
	}

	err = os.WriteFile(util.TokenFileName, []byte(token), 0600)
	if err != nil {
		panic(err)
	}

	color.Green("Successful writing\n")
}
