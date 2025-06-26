package parser

import (
	"fmt"
	"strings"
)

func Symbol(input string) (string, error) {
	atIndex := strings.Index(input, "@")
	if atIndex != -1 {
		return input[atIndex+1:], nil
	}

	leftParenIndex := strings.Index(input, "(")
	rightParenIndex := strings.Index(input, ")")

	if leftParenIndex == -1 || rightParenIndex == -1 {
		return "", fmt.Errorf("Failed to parse symbol in line %s", input)
	}

	return input[leftParenIndex+1 : rightParenIndex], nil
}
