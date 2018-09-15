package cfg

import (
	"github.com/paulgriffiths/goeval/lar"
	"io"
)

// parse parses a context-free grammar and creates a corresponding
// data structure.
func parse(input io.Reader) (*Cfg, error) {
	tokens, lerr := lex(input)
	if lerr != nil {
		return nil, lerr
	}

	c := firstPass(tokens)

	if perr := secondPass(c, tokens); perr != nil {
		return nil, perr
	}

	return c, nil
}

// secondPass takes a second pass through the grammar and extracts
// the productions.
func secondPass(c *Cfg, tokens []token) parseErr {
	reader := newTokenReader(tokens)

	for !reader.atEnd() {
		if perr := getNextProduction(c, &reader); perr != nil {
			return perr
		}
	}

	return nil
}

// getNextProduction extracts the next production.
func getNextProduction(c *Cfg, reader *tokenReader) parseErr {
	if !reader.match(tokenNonTerminal) {
		token := reader.lookahead()
		return parseError{parseErrMissingNonTerminal, "", token.pos}
	}

	head := c.ntTable[reader.current().s]

	if !reader.match(tokenArrow) {
		return parseError{parseErrMissingArrow, "",
			lar.FilePos{reader.current().pos.Ch + len(reader.current().s),
				reader.current().pos.Line}}
	}

	for {
		cmp, perr := getNextBody(c, reader)
		if perr != nil {
			return perr
		}

		c.prods[head] = append(c.prods[head], cmp)

		reader.match(tokenEndOfLine)
		if !reader.match(tokenAlt) {
			break
		}
	}

	return nil
}

// getNextBody extracts the next production body.
func getNextBody(c *Cfg, reader *tokenReader) ([]BodyComp, parseErr) {
	if reader.match(tokenEmpty) {
		if reader.peek(tokenNonTerminal) ||
			reader.peek(tokenTerminal) ||
			reader.peek(tokenEmpty) {
			token := reader.lookahead()
			return nil, parseError{parseErrEmptyNotAlone, "",
				token.pos}
		}
		return []BodyComp{BodyComp{BodyEmpty, 0}}, nil
	}

	cmps := []BodyComp{}

	for {
		if reader.match(tokenNonTerminal) {
			token := reader.current()
			cmps = append(cmps,
				BodyComp{BodyNonTerminal, c.ntTable[token.s]})
		} else if reader.match(tokenTerminal) {
			token := reader.current()
			cmps = append(cmps,
				BodyComp{BodyTerminal, c.tTable[token.s]})
		} else if reader.match(tokenEmpty) {
			token := reader.current()
			return nil, parseError{parseErrEmptyNotAlone, "",
				token.pos}
		} else {
			break
		}
	}

	if len(cmps) == 0 {
		token := reader.current()
		return nil, parseError{parseErrEmptyBody, "", token.pos}
	}

	return cmps, nil
}

// firstPass makes a first pass through the grammar to identify
// the terminals and non-terminals, and to set up the symbol tables.
func firstPass(tokens []token) *Cfg {
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

	c := Cfg{
		nonTerminals: nonTerminals,
		terminals:    terminals,
		ntTable:      ntTable,
		tTable:       tTable,
		prods:        make([][][]BodyComp, len(nonTerminals)),
	}
	return &c
}
