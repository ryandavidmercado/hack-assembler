package parser

import (
	"fmt"
	"strings"
)

type CompCode int

const (
	C_0 CompCode = iota
	C_1
	C_NEG_1
	C_D
	C_A
	C_NOT_D
	C_NOT_A
	C_NEG_D
	C_NEG_A
	C_D_PLUS_1
	C_A_PLUS_1
	C_D_MINUS_1
	C_A_MINUS_1
	C_D_PLUS_A
	C_D_MINUS_A
	C_A_MINUS_D
	C_D_AND_A
	C_D_OR_A
	C_M
	C_NOT_M
	C_NEG_M
	C_M_PLUS_1
	C_M_MINUS_1
	C_D_PLUS_M
	C_D_MINUS_M
	C_M_MINUS_D
	C_D_AND_M
	C_D_OR_M
)

var compCodeMap = map[string]CompCode{
	"0":   C_0,
	"1":   C_1,
	"-1":  C_NEG_1,
	"D":   C_D,
	"A":   C_A,
	"!D":  C_NOT_D,
	"!A":  C_NOT_A,
	"-D":  C_NEG_D,
	"-A":  C_NEG_A,
	"D+1": C_D_PLUS_1,
	"A+1": C_A_PLUS_1,
	"D-1": C_D_MINUS_1,
	"A-1": C_A_MINUS_1,
	"D+A": C_D_PLUS_A,
	"D-A": C_D_MINUS_A,
	"A-D": C_A_MINUS_D,
	"D&A": C_D_AND_A,
	"D|A": C_D_OR_A,
	"M":   C_M,
	"!M":  C_NOT_M,
	"-M":  C_NEG_M,
	"M+1": C_M_PLUS_1,
	"M-1": C_M_MINUS_1,
	"D+M": C_D_PLUS_M,
	"D-M": C_D_MINUS_M,
	"M-D": C_M_MINUS_D,
	"D&M": C_D_AND_M,
	"D|M": C_D_OR_M,
}

func Comp(input string) (CompCode, error) {
	compChars := input

	eqIndex := strings.Index(compChars, "=")
	if eqIndex != -1 {
		compChars = compChars[eqIndex+1:]
	} // everything after =

	semiIndex := strings.Index(compChars, ";")
	if semiIndex != -1 {
		compChars = compChars[0:semiIndex]
	} // everything before ;

	code, ok := compCodeMap[compChars]
	if !ok {
		return C_0, fmt.Errorf("Failed to parse %s: invalid comp '%s'", input, compChars)
	}

	return code, nil
}
