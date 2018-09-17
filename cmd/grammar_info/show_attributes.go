package main

import (
	"fmt"
	"github.com/paulgriffiths/goeval/cfg"
)

func outputAttribs(grammar *cfg.Cfg) {
	if grammar.IsLeftRecursive() {
		fmt.Printf("The grammar is left-recursive.\n")
	} else {
		fmt.Printf("The grammar is not left-recursive.\n")
	}

	fmt.Printf("The grammar has %s, %s, and %s\n",
		plural(len(grammar.NonTerminals), "nonterminal", "nonterminals"),
		plural(len(grammar.Terminals), "terminal", "terminals"),
		plural(grammar.NumProductions(), "production", "productions"),
	)
}

func outputTerminalsAndNonTerminals(grammar *cfg.Cfg) {
	nnt := len(grammar.NonTerminals)
	nt := len(grammar.Terminals)

	fmt.Printf("The %s ",
		plural(nnt, "nonterminal is", "nonterminals are"))
	printCommaList(intRange(nnt), grammar.NonTerminals, "")
	fmt.Printf(".\n")

	fmt.Printf("The %s ", plural(nt, "terminal is", "terminals are"))
	printCommaList(intRange(nt), grammar.Terminals, "`")
	fmt.Printf(".\n")
}

func outputCycles(grammar *cfg.Cfg) {
	outputNonTerminalList(grammar, grammar.NonTerminalsWithCycles(),
		"has a cycle", "have cycles")
}

func outputEProductions(grammar *cfg.Cfg) {
	outputNonTerminalList(grammar, grammar.NonTerminalsWithEProductions(),
		"has an e-production", "have e-productions")
}

func outputNullable(grammar *cfg.Cfg) {
	outputNonTerminalList(grammar, grammar.NonTerminalsNullable(),
		"is nullable", "are nullable")
}

func outputNonTerminalList(grammar *cfg.Cfg, list []int,
	singular, plural string) {
	if len(list) == 0 {
		fmt.Printf("No nonterminals %s.\n", plural)
		return
	}

	printCommaList(list, grammar.NonTerminals, "")
	if len(list) == 1 {
		fmt.Printf(" %s.\n", singular)
	} else {
		fmt.Printf(" %s.\n", plural)
	}
}
