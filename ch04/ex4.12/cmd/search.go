package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

func CmdSearch(cmd *cobra.Command, args []string) {
	entries, err := os.ReadDir(CacheDir)
	if err != nil {
		fmt.Printf("can't load cache dir, run download sub-command first. %v\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		data, err := os.ReadFile(CacheDir + "/" + entry.Name())
		if err != nil {
			fmt.Printf("loading file failed. %v\n", err)
			return
		}

		var xkcdJSON XkcdJSON
		err = json.Unmarshal(data, &xkcdJSON)
		if err != nil {
			fmt.Printf("json Unmarshal failed. %v\n", err)
			return
		}

		r := regexp.MustCompile(args[0])
		if r.MatchString(xkcdJSON.Transcript) {
			printXkcdJSON(xkcdJSON)
		}
	}
}

func printXkcdJSON(xkcdJSON XkcdJSON) {
	println("\n------------------------------")
	fmt.Printf("#%d %s (%s)\n", xkcdJSON.Num, xkcdJSON.Title, xkcdJSON.Img)
	fmt.Printf("%s\n", xkcdJSON.Transcript)
}
