package main

import (
	"flag"
	"fmt"
	"github.com/paulgriffiths/goeval/cfg"
	"os"
)

func main() {
	fileName := flag.String("f", "", "grammar file name")
	recog := flag.String("r", "", "show if the grammar recognizes "+
		"the provided string")
	showFile := flag.Bool("g", true, "show grammar representation")
	listAttribs := flag.Bool("s", true, "show grammar statistics")
	listTerms := flag.Bool("t", false, "list terminals and nonterminals")
	listCycles := flag.Bool("c", false, "list nonterminals with cycles")
	listE := flag.Bool("e", false, "list nonterminals with e-productions")
	listNull := flag.Bool("n", false, "list nonterminals which are nullable")
	listAll := flag.Bool("a", false, "list all grammar information "+
		"(equivalent to -stcen)")
	flag.Parse()

	if *fileName == "" {
		fmt.Fprintf(os.Stderr, "grammar_info: no filename provided.\n")
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file %q: %v\n",
			os.Args[1], err)
		os.Exit(1)
	}

	defer file.Close()

	grammar, err := cfg.NewCfg(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't create grammar: %v\n", err)
		os.Exit(1)
	}

	if *showFile {
		outputFile(file)
	}
	if *listAttribs || *listAll {
		outputAttribs(grammar)
	}
	if *listTerms || *listAll {
		outputTerminalsAndNonTerminals(grammar)
	}
	if *listCycles || *listAll {
		outputCycles(grammar)
	}
	if *listE || *listAll {
		outputEProductions(grammar)
	}
	if *listNull || *listAll {
		outputNullable(grammar)
	}
	if *recog != "" {
		if grammar.IsLeftRecursive() {
			fmt.Printf("Parsing currently only implemented for non-" +
				"left-recursive grammars.\n")
		} else {
			recognizeRdpString(grammar, *recog)
		}
	}
}
