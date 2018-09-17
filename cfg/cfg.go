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
	Prods        []BodyList
}

// NewCfg returns a new context-free grammer from the provided
// description.
func NewCfg(reader io.Reader) (*Cfg, error) {
	newCfg, err := parse(reader)
	return newCfg, err
}

// MaxNonTerminalNameLength returns the length of the longest
// nonterminal name.
func (c *Cfg) MaxNonTerminalNameLength() int {
	maxNL := -1
	for _, nt := range c.NonTerminals {
		if len(nt) > maxNL {
			maxNL = len(nt)
		}
	}
	return maxNL
}

// NumProductions returns the number of productions in the grammar.
func (c *Cfg) NumProductions() int {
	n := 0
	for _, body := range c.Prods {
		n += len(body)
	}
	return n
}

// outputCfg outputs a representation of the grammar.
func (c *Cfg) outputCfg(writer io.Writer) {
	maxNL := c.MaxNonTerminalNameLength()

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
				switch {
				case cmp.IsNonTerminal():
					s = fmt.Sprintf(" %s", c.NonTerminals[cmp.I])
				case cmp.IsTerminal():
					s = fmt.Sprintf(" `%s`", c.Terminals[cmp.I])
				case cmp.IsEmpty():
					s = " e"
				default:
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
		if body[0].IsNonTerminal() && body[0].I == nt {
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
		if body[0].IsNonTerminal() {
			if body[0].I == nt {
				return true
			} else if c.lrInternal(nt, body[0].I, checked) {
				return true
			}
		}
	}
	return false
}

// NonTerminalsWithCycles returns a list of nonterminals which
// have cycles.
func (c *Cfg) NonTerminalsWithCycles() []int {
	list := []int{}
	for n := range c.NonTerminals {
		if c.hcInternal(n, n, make(map[int]bool)) {
			list = append(list, n)
		}
	}
	return list
}

// HasCycle checks if the grammar contains a cycle.
func (c *Cfg) HasCycle() bool {
	for n := range c.NonTerminals {
		if c.hcInternal(n, n, make(map[int]bool)) {
			return true
		}
	}
	return false
}

func (c *Cfg) hcInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	for _, body := range c.Prods[interNt] {
		if body.IsNonTerminal() {
			if body[0].I == nt {
				return true
			} else if c.hcInternal(nt, body[0].I, checked) {
				return true
			}
		}
	}
	return false
}

// NonTerminalsNullable returns a list of nonterminals which
// are nullable.
func (c *Cfg) NonTerminalsNullable() []int {
	list := []int{}
	for n := range c.Prods {
		if c.inInternal(n, n, make(map[int]bool)) {
			list = append(list, n)
		}
	}
	return list
}

// IsNullable checks if a nonterminal is nullable.
func (c *Cfg) IsNullable(nt int) bool {
	return c.inInternal(nt, nt, make(map[int]bool))
}

func (c *Cfg) inInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	if c.Prods[interNt].HasEmpty() {
		return true
	}

	for _, body := range c.Prods[interNt] {
		nullable := true
		for _, comp := range body {
			if !comp.IsNonTerminal() {
				nullable = false
				break
			}
			if !c.inInternal(nt, comp.I, checked) {
				nullable = false
				break
			}
		}
		if nullable {
			return true
		}
	}
	return false
}

// HasEProduction checks if the grammar has an e-production.
func (c *Cfg) HasEProduction() bool {
	for _, prod := range c.Prods {
		if prod.HasEmpty() {
			return true
		}
	}
	return false
}

// NonTerminalsWithEProductions returns a list of nonterminals which
// have e-productions.
func (c *Cfg) NonTerminalsWithEProductions() []int {
	list := []int{}
	for n, prod := range c.Prods {
		if prod.HasEmpty() {
			list = append(list, n)
		}
	}
	return list
}
