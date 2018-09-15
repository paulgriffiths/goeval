package cfl

import (
	"io"
)

func parse(input io.Reader) (*Cfl, error) {
	tokens, err := lex(input)
	if err != nil {
		return nil, err
	}

	nonTerminals := []string{}
	terminals := []string{}
	ntTable := make(map[string]int)
	tTable := make(map[string]int)

	for _, token := range tokens {
		switch token.t {
		case tokenNonTerminal:
			if _, ok := ntTable[token.s]; !ok {
				ntTable[token.s] = len(nonTerminals)
				nonTerminals = append(nonTerminals, token.s)
			}
		case tokenTerminal:
			if _, ok := tTable[token.s]; !ok {
				tTable[token.s] = len(terminals)
				terminals = append(terminals, token.s)
			}
		}
	}

	c := Cfl{nonTerminals: nonTerminals, terminals: terminals,
		ntTable: ntTable, tTable: tTable}
	return &c, nil
}
