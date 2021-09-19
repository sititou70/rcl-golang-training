package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type IssueInput struct {
	Title string
	Body  string
}

func InputIssue(defaultIssue IssueInput) (IssueInput, error) {
	editor := os.Getenv("EDITOR")
	if len(editor) == 0 {
		editor = "vim"
	}

	defaultTitle := "write issue title here"
	if defaultIssue.Title != "" {
		defaultTitle = defaultIssue.Title
	}
	defaultBody := "write issue body here"
	if defaultIssue.Body != "" {
		defaultBody = defaultIssue.Body
	}

	template := fmt.Sprintf(`%s
----------
%s
`, defaultTitle, defaultBody)
	inputFile, err := ioutil.TempFile("", "input-issue")
	if err != nil {
		return IssueInput{}, err
	}

	inputFile.Write([]byte(template))
	cmd := exec.Command(editor, inputFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return IssueInput{}, fmt.Errorf("aborted")
	}

	input, err := os.ReadFile(inputFile.Name())
	if err != nil {
		return IssueInput{}, err
	}
	splited := strings.Split(string(input), "\n----------\n")
	if len(splited) != 2 {
		return IssueInput{}, fmt.Errorf("format invalid, use the separator '----------' only once")
	}

	os.Remove(inputFile.Name())
	return IssueInput{splited[0], splited[1]}, nil
}
