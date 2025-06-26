package main

import (
	"github.com/ryandavidmercado/hack-assembler/assembler"
	"log"
	"os"
)

func main() {
	args := parseArgs()

	asm, err := os.ReadFile(args.in)
	if err != nil {
		log.Fatal(err)
	}

	hack := assembler.Assemble(string(asm))

	writeErr := os.WriteFile(args.out, []byte(hack), 0644)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
