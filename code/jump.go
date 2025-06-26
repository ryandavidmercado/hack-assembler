package code

import (
	"github.com/ryandavidmercado/hack-assembler/parser"
)

var jumpCodeMap = map[parser.JumpCode]string{
	parser.J_NULL: "000",
	parser.J_JGT:  "001",
	parser.J_JEQ:  "010",
	parser.J_JGE:  "011",
	parser.J_JLT:  "100",
	parser.J_JNE:  "101",
	parser.J_JLE:  "110",
	parser.J_JMP:  "111",
}

func Jump(input parser.JumpCode) string {
	return jumpCodeMap[input]
}
