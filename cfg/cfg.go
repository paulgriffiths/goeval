package cfg

import (
	"fmt"
	"io"
)

// Cfg represents a context-free grammar.
type Cfg struct {
	NonTerminals []string
	Terminals    []string
	NtTable      map[string]int
	TTable       map[string]int
	Prods        [][][]BodyComp
}

// NewCfg returns a new context-free grammer from the provided
// description.
func NewCfg(reader io.Reader) (*Cfg, error) {
	newCfg, err := parse(reader)
	return newCfg, err
}

// outputCfg outputs a representation of the grammar.
func (c *Cfg) outputCfg(writer io.Writer) {
	maxNL := -1
	for _, nt := range c.NonTerminals {
		if len(nt) > maxNL {
			maxNL = len(nt)
		}
	}

	for i, prod := range c.Prods {
		for n, body := range prod {
			var s string
			if n == 0 {
				s = fmt.Sprintf("%-[1]*s :", maxNL, c.NonTerminals[i])
			} else {
				s = fmt.Sprintf("%-[1]*s |", maxNL, "")
			}
			writer.Write([]byte(s))

			for _, cmp := range body {
				if cmp.T == BodyNonTerminal {
					s = fmt.Sprintf(" %s", c.NonTerminals[cmp.I])
				} else if cmp.T == BodyTerminal {
					s = fmt.Sprintf(" `%s`", c.Terminals[cmp.I])
				} else if cmp.T == BodyEmpty {
					s = " e"
				} else {
					panic("unexpected body component")
				}
				writer.Write([]byte(s))
			}
			writer.Write([]byte("\n"))
		}
	}
}

// IsImmediateLeftRecursive checks if the specified nonterminal
// is immediately left-recursive.
func (c *Cfg) IsImmediateLeftRecursive(nt int) bool {
	for _, body := range c.Prods[nt] {
		if body[0].T == BodyNonTerminal && body[0].I == nt {
			return true
		}
	}
	return false
}

// IsLeftRecursive checks if the grammar is left-recursive.
func (c *Cfg) IsLeftRecursive() bool {
	for n := range c.NonTerminals {
		if c.lrInternal(n, n, make(map[int]bool)) {
			return true
		}
	}
	return false
}

func (c *Cfg) lrInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	for _, body := range c.Prods[interNt] {
		if body[0].T == BodyNonTerminal {
			if body[0].I == nt {
				return true
			} else if c.lrInternal(nt, body[0].I, checked) {
				return true
			}
		}
	}
	return false
}

// HasCycle checks if the grammar contains a cycle.
func (c *Cfg) HasCycle() bool {
	for n := range c.NonTerminals {
		if c.hsInternal(n, n, make(map[int]bool)) {
			return true
		}
	}
	return false
}

func (c *Cfg) hsInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	for _, body := range c.Prods[interNt] {
		if len(body) == 1 && body[0].T == BodyNonTerminal {
			if body[0].I == nt {
				return true
			} else if c.hsInternal(nt, body[0].I, checked) {
				return true
			}
		}
	}
	return false
}

// HasEProduction checks if the grammar has an e-production.
func (c *Cfg) HasEProduction() bool {
	for _, prod := range c.Prods {
		for _, body := range prod {
			for _, comp := range body {
				if comp.T == BodyEmpty {
					return true
				}
			}
		}
	}
	return false
}
