package parser

import "strings"

type InstructionType int

const (
	A_INSTRUCTION InstructionType = iota
	C_INSTRUCTION
	L_INSTRUCTION
	NIL
)

func instructionType(line string) InstructionType {
	if strings.HasPrefix(line, "@") {
		return A_INSTRUCTION
	} else if strings.HasPrefix(line, "(") {
		return L_INSTRUCTION
	} else {
		return C_INSTRUCTION
	}
}
