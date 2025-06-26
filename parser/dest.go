package parser

import (
	"fmt"
	"log"
	"strings"
)

type DestCode int

const (
	D_M DestCode = iota
	D_D
	D_DM
	D_MD
	D_A
	D_AM
	D_AD
	D_ADM
	D_NULL
)

var destCodeMap = map[string]DestCode{
	"M":   D_M,
	"D":   D_D,
	"DM":  D_DM,
	"MD":  D_MD,
	"A":   D_A,
	"AM":  D_AM,
	"AD":  D_AD,
	"ADM": D_ADM,
}

func Dest(input string) DestCode {
	log.SetPrefix("parser/Dest():")

	eqIndex := strings.Index(input, "=")
	if eqIndex == -1 {
		return D_NULL
	}

	destChars := input[0:eqIndex]

	code, ok := destCodeMap[destChars]
	if !ok {
		log.Fatal(fmt.Errorf("Failed to parse '%s': invalid dest '%s'", input, destChars))
	}

	return code
}
