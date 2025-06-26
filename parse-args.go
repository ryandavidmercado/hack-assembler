package main

import (
	"errors"
	"flag"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

type args struct {
	in  string
	out string
}

func parseArgs() *args {
	log.SetPrefix("parseArgs(): ")

	inPtr := flag.String("in", "", "input file")
	var out string
	flag.StringVar(&out, "out", "", "output file")

	flag.Parse()

	checkFilename := regexp.MustCompile(`\.asm$`)

	if len(*inPtr) == 0 {
		log.Fatal("Must pass in input file")
	}

	matches := checkFilename.MatchString(*inPtr)

	if !matches {
		log.Fatal(errors.New("Input filename must have .asm extension"))
	}

	if len(out) == 0 {
		out = strings.TrimSuffix(*inPtr, filepath.Ext(*inPtr)) + ".hack"
	}

	args := args{in: *inPtr, out: out}
	return &args
}
