package cfg

import (
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
		return parseError{parseErrMissingHead, token.pos.LineOnly()}
	}

	head := c.NtTable[reader.current().s]

	if !reader.match(tokenArrow) {
		token := reader.current()
		return parseError{parseErrMissingArrow,
			token.pos.Advance(len(reader.current().s))}
	}

	for {
		cmp, perr := getNextBody(c, reader)
		if perr != nil {
			return perr
		}

		c.Prods[head] = append(c.Prods[head], cmp)

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
		token := reader.current()
		if reader.peek(tokenNonTerminal) ||
			reader.peek(tokenTerminal) ||
			reader.peek(tokenEmpty) {
			return nil, parseError{parseErrEmptyNotAlone,
				token.pos.Advance(1)}
		}
		return []BodyComp{BodyComp{BodyEmpty, 0}}, nil
	}

	cmps := []BodyComp{}

	for {
		if reader.match(tokenNonTerminal) {
			token := reader.current()
			cmps = append(cmps,
				BodyComp{BodyNonTerminal, c.NtTable[token.s]})
		} else if reader.match(tokenTerminal) {
			token := reader.current()
			cmps = append(cmps,
				BodyComp{BodyTerminal, c.TTable[token.s]})
		} else if reader.match(tokenEmpty) {
			token := reader.current()
			return nil, parseError{parseErrEmptyNotAlone,
				token.pos.Advance(1)}
		} else {
			break
		}
	}

	if len(cmps) == 0 {
		token := reader.current()
		return nil, parseError{parseErrEmptyBody,
			token.pos.Advance(1)}
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
		NonTerminals: nonTerminals,
		Terminals:    terminals,
		NtTable:      ntTable,
		TTable:       tTable,
		Prods:        make([]BodyList, len(nonTerminals)),
		firsts:       nil,
		follows:      nil,
	}
	return &c
}
