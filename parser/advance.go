package parser

import (
	"bufio"
	"strings"
)

func Advance(scanner *bufio.Scanner) InstructionType {
	inComment := false

	for line := scanner.Scan(); line; line = scanner.Scan() {
		trimmed := strings.TrimSpace(scanner.Text())

		/* Handle multi-line comments */
		if !inComment && strings.HasPrefix(trimmed, "/*") {
			if !strings.HasSuffix(trimmed, "*/") {
				inComment = true
			}

			continue
		}

		if inComment {
			if strings.HasSuffix(trimmed, "*/") {
				inComment = false
			}

			continue
		}
		/* -------------------------- */

		/* Handle single-line comments and empty lines */
		if strings.HasPrefix(trimmed, "//") || len(trimmed) == 0 {
			continue
		}
		/* ------------------------------------------- */

		instructionType := instructionType(trimmed)

		if instructionType != NIL {
			return instructionType
		}
	}

	return NIL
}
