package pp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"github.com/paulgriffiths/goeval/parsers/lexer"
	"github.com/paulgriffiths/goeval/parsers/tree"
)

// Parse parses input against a grammar and returns a parse tree,
// or nil on failure.
func (p Pp) Parse(input string) *tree.Node {
	terminals, err := lexer.Lex(p.g, input)
	if err != nil {
		return nil
	}

	node, n := p.parseNT(terminals, 0)
	if n == len(terminals) {
		return node
	}
	return nil
}

// parseComp parses a production body component.
func (p Pp) parseComp(t lexer.TerminalList, comp cfg.BodyComp) (*tree.Node, int) {
	var node *tree.Node
	numTerms := 0

	switch comp.T {
	case cfg.BodyNonTerminal:
		node, numTerms = p.parseNT(t, comp.I)
	case cfg.BodyTerminal:
		if !t.IsEmpty() && t[0].N == comp.I {
			node, numTerms = tree.NewNode(comp, t[0].S, nil), 1
		}
	case cfg.BodyEmpty:
		node = tree.NewNode(comp, "e", nil)
	}

	return node, numTerms
}

// parseNT parses a non-terminal.
func (p Pp) parseNT(t lexer.TerminalList, nt int) (*tree.Node, int) {

	// If there are no more terminals in the list, check whether
	// the current nonterminal can be followed by end-of-input.

	if t.IsEmpty() {
		body := p.table[nt][len(p.g.Terminals)]
		if body.IsEmpty() {
			return nil, 0
		}
		if body[0].IsEmptyBody() {
			em := tree.NewNode(cfg.NewBodyEmpty(), "e", nil)
			term := tree.NewNode(cfg.NewNonTerminal(nt),
				p.g.NonTerminals[nt], []*tree.Node{em})
			return term, 0
		}
		panic("unexpected terminal condition")
	}

	// Get the body for this nonterminal with the next terminal,
	// returning an error if the predictive parsing table doesn't
	// contain an entry.

	body := p.table[nt][t[0].N]
	if body.IsEmpty() {
		return nil, 0
	}

	if children, numTerms := p.parseBody(t, body[0]); children != nil {
		return tree.NewNode(
			cfg.BodyComp{cfg.BodyNonTerminal, nt},
			p.g.NonTerminals[nt],
			children,
		), numTerms
	}

	return nil, 0
}

// parseBody parses a production body.
func (p Pp) parseBody(t lexer.TerminalList, body []cfg.BodyComp) ([]*tree.Node, int) {
	var children []*tree.Node
	matchLength := 0

	for _, component := range body {
		node, numTerms := p.parseComp(t[matchLength:], component)
		if node == nil {
			return nil, 0
		}
		children = append(children, node)
		matchLength += numTerms
	}

	return children, matchLength
}
