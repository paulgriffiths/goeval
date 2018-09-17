package main

import (
	"bytes"
	"fmt"
	"github.com/paulgriffiths/goeval/cfg"
	"github.com/paulgriffiths/goeval/parsers/rdp"
	"os"
)

func recognizeRdpString(grammar *cfg.Cfg, input string) {
	parser, err := rdp.New(grammar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create parser: %v\n", err)
		return
	}

	tree := parser.Parse(input)
	if tree != nil {
		fmt.Printf("Grammar recognizes string '%s'.\n", input)
	} else {
		fmt.Printf("Grammar does not recognize string '%s'.\n", input)
	}
}

func parseRdpString(grammar *cfg.Cfg, input string) {
	parser, err := rdp.New(grammar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create parser: %v\n", err)
		return
	}

	tree := parser.Parse(input)
	if tree == nil {
		fmt.Printf("Grammar does not recognize string '%s'.\n", input)
		return
	}

	outBuffer := bytes.NewBuffer(nil)
	tree.WriteBracketed(outBuffer, "`")
	output := string(outBuffer.Bytes())

	fmt.Printf("Parse tree for string '%s': %s\n", input, output)
}
