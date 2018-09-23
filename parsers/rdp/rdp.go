package rdp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"github.com/paulgriffiths/goeval/parsers/lexer"
	"github.com/paulgriffiths/goeval/parsers/tree"
)

// Rdp implements a recursive descent parser.
type Rdp struct {
	c     *cfg.Cfg
	lexer *lexer.Lexer
}

// New creates a new recursive descent parser.
func New(c *cfg.Cfg) (*Rdp, error) {
	l, err := lexer.New(c)
	if err != nil {
		return nil, err
	}
	r := Rdp{c, l}
	return &r, nil
}

// Parse parses input against a grammar and returns a parse tree,
// or nil on failure.
func (r Rdp) Parse(input string) *tree.Node {
	terminals, err := r.lexer.Lex(input)
	if err != nil {
		return nil
	}

	node, n := r.parseNT(terminals, 0)
	if n == len(terminals) {
		return node
	}
	return nil
}

// parseComp parses a production body component.
func (r Rdp) parseComp(t lexer.TerminalList, comp cfg.BodyComp) (*tree.Node, int) {
	var node *tree.Node
	numTerms := 0

	switch comp.T {
	case cfg.BodyNonTerminal:
		node, numTerms = r.parseNT(t, comp.I)
	case cfg.BodyTerminal:
		if len(t) > 0 && comp.I == t[0].N {
			node, numTerms = tree.NewNode(comp, t[0].S, nil), 1
		}
	case cfg.BodyEmpty:
		node = tree.NewNode(comp, "e", nil)
	}

	return node, numTerms
}

// parseNT parses a non-terminal.
func (r Rdp) parseNT(t lexer.TerminalList, nt int) (*tree.Node, int) {
	for _, body := range r.c.Prods[nt] {
		if children, numTerms := r.parseBody(t, body); children != nil {
			return tree.NewNode(
				cfg.BodyComp{cfg.BodyNonTerminal, nt},
				r.c.NonTerminals[nt],
				children,
			), numTerms
		}
	}

	return nil, 0
}

// parseBody parses a production body.
func (r Rdp) parseBody(t lexer.TerminalList, body []cfg.BodyComp) ([]*tree.Node, int) {
	var children []*tree.Node
	matchLength := 0

	for _, component := range body {
		node, numTerms := r.parseComp(t[matchLength:], component)
		if node == nil {
			return nil, 0
		}
		children = append(children, node)
		matchLength += numTerms
	}

	return children, matchLength
}
