package code

import (
	"github.com/ryandavidmercado/hack-assembler/parser"
)

var destCodeMap = map[parser.DestCode]string{
	parser.D_NULL: "000",
	parser.D_M:    "001",
	parser.D_D:    "010",
	parser.D_DM:   "011",
	parser.D_MD:   "011",
	parser.D_A:    "100",
	parser.D_AM:   "101",
	parser.D_AD:   "110",
	parser.D_ADM:  "111",
}

func Dest(input parser.DestCode) string {
	return destCodeMap[input]
}
