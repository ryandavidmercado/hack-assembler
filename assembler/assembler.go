package assembler

import (
	"fmt"
	"github.com/ryandavidmercado/hack-assembler/code"
	"github.com/ryandavidmercado/hack-assembler/parser"
	"log"
	"strconv"
	"strings"
)

func Assemble(asmIn string) string {
	asmPass1 := parser.Parser(asmIn)
	asmPass2 := parser.Parser(asmIn)
	hack := ""

	pass1Line := 0
	// pass 1: gather label symbols
	for ins := parser.Advance(asmPass1); ins != parser.NIL; ins = parser.Advance(asmPass1) {
		line := strings.TrimSpace(asmPass1.Text())

		switch ins {
		case parser.C_INSTRUCTION, parser.A_INSTRUCTION:
			pass1Line++
		case parser.L_INSTRUCTION:
			symbol := parser.Symbol(line)
			symbolMap[symbol] = strconv.Itoa(pass1Line)
		}
	}

	// pass 2: assemble
	var pass2VarRegister = 16
	for ins := parser.Advance(asmPass2); ins != parser.NIL; ins = parser.Advance(asmPass2) {
		line := strings.TrimSpace(asmPass2.Text())

		switch ins {
		case parser.C_INSTRUCTION:
			{
				destCode, compCode, jumpCode := parser.Dest(line), parser.Comp(line), parser.Jump(line)
				dest, comp, jump := code.Dest(destCode), code.Comp(compCode), code.Jump(jumpCode)

				hack += "111" + comp + dest + jump + "\n"
			}
		case parser.A_INSTRUCTION:
			{
				symbol := parser.Symbol(line)
				decimal, err := strconv.ParseInt(symbol, 10, 16)

				if err != nil {
					value, inMap := symbolMap[symbol]
					if inMap {
						res, err := strconv.ParseInt(value, 10, 16)
						if err != nil {
							log.SetPrefix("assember/Assemble():")
							log.Fatalf("bad value in symbolMap[symbol]: %s", value)
						}

						decimal = res
					} else {
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

	return hack
}
