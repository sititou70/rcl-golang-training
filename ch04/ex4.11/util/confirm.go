package util

import (
	"github.com/manifoldco/promptui"
)

func Confirm(title string) bool {
	prompt := promptui.Prompt{
		Label:     title,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return false
	}
	for _, str := range []string{"y", "Y", "yes"} {
		if result == str {
			return true
		}
	}

	return false
}
