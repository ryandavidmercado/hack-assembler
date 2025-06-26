package parser

import (
	"bufio"
	"strings"
)

func Parser(asm string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(asm))
}
