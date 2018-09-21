package cfg

import "io"

// Cfg represents a context-free grammar.
type Cfg struct {
	NonTerminals []string
	Terminals    []string
	NtTable      map[string]int
	TTable       map[string]int
	Prods        []BodyList
	Firsts       []SetBodyComp
	Follows      []SetBodyComp
}

// NewCfg returns a new context-free grammer from the provided
// description.
func NewCfg(reader io.Reader) (*Cfg, error) {
	return parse(reader)
}

// NonTerminalComp returns a body component for the nonterminal
// named in the provided string.
func (c *Cfg) NonTerminalComp(nt string) BodyComp {
	return BodyComp{BodyNonTerminal, c.NtTable[nt]}
}

// TerminalComp returns a body component for the terminal
// named in the provided string.
func (c *Cfg) TerminalComp(nt string) BodyComp {
	return BodyComp{BodyTerminal, c.TTable[nt]}
}
