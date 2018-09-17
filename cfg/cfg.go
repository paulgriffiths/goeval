package cfg

import "io"

// Cfg represents a context-free grammar.
type Cfg struct {
	NonTerminals []string
	Terminals    []string
	NtTable      map[string]int
	TTable       map[string]int
	Prods        []BodyList
}

// NewCfg returns a new context-free grammer from the provided
// description.
func NewCfg(reader io.Reader) (*Cfg, error) {
	newCfg, err := parse(reader)
	return newCfg, err
}
