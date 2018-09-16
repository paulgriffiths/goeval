package rdp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"regexp"
)

// Rdp implements a recursive descent parser.
type Rdp struct {
	c    *cfg.Cfg
	regs []*regexp.Regexp
}

// New creates a new recursive descent parser.
func New(c *cfg.Cfg) *Rdp {
	regs := []*regexp.Regexp{}
	for _, t := range c.Terminals {
		reg, err := regexp.Compile(t)
		if err != nil {
			return nil
		}
		regs = append(regs, reg)
	}
	r := Rdp{c, regs}
	return &r
}

// Accepts checks if a context-free grammer accepts a string.
func (r Rdp) Accepts(input string) bool {
	match, _ := r.acceptNT(input, 0)
	return match
}

// acceptNT checks if a nonterminal accepts a prefix of a string.
func (r Rdp) acceptNT(input string, nt int) (bool, int) {
	for _, body := range r.c.Prods[nt] {
		if match, n := r.acceptBody(input, body); match {
			return true, n
		}
	}

	return false, 0
}

// acceptBody checks if a production body accepts a prefix of a string.
func (r Rdp) acceptBody(input string, body []cfg.BodyComp) (bool, int) {
	n := 0

	for _, comp := range body {
		match, c := r.acceptComp(input[n:], comp)
		if !match {
			return false, 0
		}
		n += c
	}

	return true, n
}

// acceptComp checks if a production body component accepts a
// prefix of a string.
func (r Rdp) acceptComp(input string, comp cfg.BodyComp) (bool, int) {
	match := false
	n := 0

	if comp.T == cfg.BodyNonTerminal {
		match, n = r.acceptNT(input, comp.I)
	} else if comp.T == cfg.BodyTerminal {
		loc := r.regs[comp.I].FindIndex([]byte(input))
		if loc != nil && loc[0] == 0 {
			match = true
			n = loc[1] - loc[0]
		}
	} else if comp.T == cfg.BodyEmpty {
		match = true
	}

	if !match {
		return false, 0
	}

	return true, n
}
