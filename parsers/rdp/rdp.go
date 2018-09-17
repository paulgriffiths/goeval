package rdp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"github.com/paulgriffiths/goeval/parsers/tree"
	"regexp"
)

// Rdp implements a recursive descent parser.
type Rdp struct {
	c    *cfg.Cfg
	regs []*regexp.Regexp
}

// New creates a new recursive descent parser.
func New(c *cfg.Cfg) (*Rdp, error) {
	regs := []*regexp.Regexp{}
	for _, t := range c.Terminals {
		reg, err := regexp.Compile(t)
		if err != nil {
			return nil, err
		}
		regs = append(regs, reg)
	}
	r := Rdp{c, regs}
	return &r, nil
}

// Parse parses input against a grammar and returns a parse tree,
// or nil on failure.
func (r Rdp) Parse(input string) *tree.Node {
	node, n := r.parseNT(input, 0)
	if n == len(input) {
		return node
	}
	return nil
}

// parseComp parses a production body component.
func (r Rdp) parseComp(input string, comp cfg.BodyComp) (*tree.Node, int) {
	var node *tree.Node
	numBytes := 0

	switch comp.T {
	case cfg.BodyNonTerminal:
		node, numBytes = r.parseNT(input, comp.I)
	case cfg.BodyTerminal:
		loc := r.regs[comp.I].FindIndex([]byte(input))
		if loc != nil && loc[0] == 0 {
			numBytes = loc[1] - loc[0]
			node = tree.NewNode(comp, input[:numBytes], nil)
		}
	case cfg.BodyEmpty:
		node = tree.NewNode(comp, "e", nil)
	}

	return node, numBytes
}

// parseNT parses a non-terminal.
func (r Rdp) parseNT(s string, nt int) (*tree.Node, int) {
	for _, body := range r.c.Prods[nt] {
		if children, numBytes := r.parseBody(s, body); children != nil {
			return tree.NewNode(
				cfg.BodyComp{cfg.BodyNonTerminal, nt},
				r.c.NonTerminals[nt],
				children,
			), numBytes
		}
	}

	return nil, 0
}

// parseBody parses a production body.
func (r Rdp) parseBody(s string, body []cfg.BodyComp) ([]*tree.Node, int) {
	var children []*tree.Node
	matchLength := 0

	for _, component := range body {
		node, numBytes := r.parseComp(s[matchLength:], component)
		if node == nil {
			return nil, 0
		}
		children = append(children, node)
		matchLength += numBytes
	}

	return children, matchLength
}
