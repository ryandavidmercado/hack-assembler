package assembler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryandavidmercado/hack-assembler/code"
	"github.com/ryandavidmercado/hack-assembler/parser"
)

func Assemble(asmIn string) (string, error) {
	asmPass1 := parser.Parser(asmIn)
	asmPass2 := parser.Parser(asmIn)
	hack := ""

	pass1Line := 0
	// pass 1: gather label symbols
	for ins := parser.Advance(asmPass1); ins != parser.NIL; ins = parser.Advance(asmPass1) {
		line := strings.TrimSpace(strings.Split(asmPass1.Text(), "//")[0])

		switch ins {
		case parser.C_INSTRUCTION, parser.A_INSTRUCTION:
			pass1Line++
		case parser.L_INSTRUCTION:
			symbol, err := parser.Symbol(line)
			if err != nil {
				return "", err
			}

			symbolMap[symbol] = strconv.Itoa(pass1Line)
		}
	}

	// pass 2: assemble
	var pass2VarRegister = 16
	for ins := parser.Advance(asmPass2); ins != parser.NIL; ins = parser.Advance(asmPass2) {
		line := strings.TrimSpace(strings.Split(asmPass2.Text(), "//")[0])

		switch ins {
		case parser.C_INSTRUCTION:
			{
				destCode, err := parser.Dest(line)
				if err != nil {
					return "", err
				}

				compCode, err := parser.Comp(line)
				if err != nil {
					return "", err
				}

				jumpCode, err := parser.Jump(line)

				if err != nil {
					return "", err
				}

				dest, comp, jump := code.Dest(destCode), code.Comp(compCode), code.Jump(jumpCode)

				hack += "111" + comp + dest + jump + "\n"
			}
		case parser.A_INSTRUCTION:
			{
				symbol, err := parser.Symbol(line)
				if err != nil {
					return "", err
				}

				decimal, err := strconv.ParseInt(symbol, 10, 16)

				if err != nil {
					value, inMap := symbolMap[symbol]
					if inMap {
						res, err := strconv.ParseInt(value, 10, 16)
						if err != nil {
							return "", err
						}

						decimal = res
					} else {
						if pass2VarRegister > 255 {
							return "", fmt.Errorf(`Can't assign index to symbol "%s": max assignments (255) reached`, symbol)
						}

						symbolMap[symbol] = strconv.Itoa(pass2VarRegister)
						decimal = int64(pass2VarRegister)
						pass2VarRegister++
					}
				}

				code := fmt.Sprintf("%015b", decimal)
				hack += "0" + code + "\n"
			}
		}
	}

	return hack, nil
}
