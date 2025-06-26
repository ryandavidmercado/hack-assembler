package code

import (
	"github.com/ryandavidmercado/hack-assembler/parser"
)

var compCodeMap = map[parser.CompCode]string{
	parser.C_0:         "0101010",
	parser.C_1:         "0111111",
	parser.C_NEG_1:     "0111010",
	parser.C_D:         "0001100",
	parser.C_A:         "0110000",
	parser.C_NOT_D:     "0001101",
	parser.C_NOT_A:     "0110001",
	parser.C_NEG_D:     "0001111",
	parser.C_NEG_A:     "0110011",
	parser.C_D_PLUS_1:  "0011111",
	parser.C_A_PLUS_1:  "0110111",
	parser.C_D_MINUS_1: "0001110",
	parser.C_A_MINUS_1: "0110010",
	parser.C_D_PLUS_A:  "0000010",
	parser.C_D_MINUS_A: "0010011",
	parser.C_A_MINUS_D: "0000111",
	parser.C_D_AND_A:   "0000000",
	parser.C_D_OR_A:    "0010101",
	parser.C_M:         "1110000",
	parser.C_NOT_M:     "1110001",
	parser.C_NEG_M:     "1110011",
	parser.C_M_PLUS_1:  "1110111",
	parser.C_M_MINUS_1: "1110010",
	parser.C_D_PLUS_M:  "1000010",
	parser.C_D_MINUS_M: "1010011",
	parser.C_M_MINUS_D: "1000111",
	parser.C_D_AND_M:   "1000000",
	parser.C_D_OR_M:    "1010101",
}

func Comp(input parser.CompCode) string {
	return compCodeMap[input]
}
