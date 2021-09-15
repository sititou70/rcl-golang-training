// page 83
package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	r := regexp.MustCompile(`[+-]?([0-9]+)(\.[0-9]+)?`)
	index := r.FindStringSubmatchIndex(s)[2:4]
	return s[:index[0]] + commaInt(s[index[0]:index[1]]) + s[index[1]:]
}

func commaInt(s string) string {
	var ans bytes.Buffer
	length := len(s)
	commaNum := (length - 1) / 3

	upperDigits := length - commaNum*3
	ans.Write([]byte(s[0:upperDigits]))

	for i := 0; i < commaNum; i++ {
		extractBeginIndex := upperDigits + i*3
		ans.Write([]byte("," + s[extractBeginIndex:extractBeginIndex+3]))
	}

	return ans.String()
}
