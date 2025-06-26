package parser

import (
	"fmt"
	"strings"
)

type JumpCode int

const (
	J_JGT JumpCode = iota
	J_JEQ
	J_JGE
	J_JLT
	J_JNE
	J_JLE
	J_JMP
	J_NULL
)

var jumpCodeMap = map[string]JumpCode{
	"JGT": J_JGT,
	"JEQ": J_JEQ,
	"JGE": J_JGE,
	"JLT": J_JLT,
	"JNE": J_JNE,
	"JLE": J_JLE,
	"JMP": J_JMP,
}

func Jump(input string) (JumpCode, error) {
	semiIndex := strings.Index(input, ";")
	if semiIndex == -1 {
		return J_NULL, nil
	}

	jumpChars := input[semiIndex+1:]

	code, ok := jumpCodeMap[jumpChars]
	if !ok {
		return J_NULL, fmt.Errorf("Failed to parse '%s': invalid jump '%s'", input, jumpChars)
	}

	return code, nil
}
