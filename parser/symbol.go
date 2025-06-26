package parser

import (
	"fmt"
	"log"
	"strings"
)

func Symbol(input string) string {
	log.SetPrefix("parser/Symbol():")

	atIndex := strings.Index(input, "@")
	if atIndex != -1 {
		return input[atIndex+1:]
	}

	leftParenIndex := strings.Index(input, "(")
	rightParenIndex := strings.Index(input, ")")

	if leftParenIndex == -1 || rightParenIndex == -1 {
		log.Fatal(fmt.Errorf("Failed to parse symbol in line %s", input))
	}

	return input[leftParenIndex+1 : rightParenIndex]
}
