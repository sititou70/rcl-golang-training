package util

import (
	"fmt"
	"os"
)

const TokenFileName = ".github_token.secret"

func GetToken() (string, error) {
	token, err := os.ReadFile(TokenFileName)
	if err != nil {
		return "", fmt.Errorf("token load error, %v", err)
	}

	return string(token), nil
}
