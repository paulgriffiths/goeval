package cfg

// BodyCompType represents the type of a BodyComp.
type BodyCompType int

const (
	// BodyNonTerminal represents a nonterminal in a production body.
	BodyNonTerminal BodyCompType = iota
	// BodyTerminal represents a terminal in a production body.
	BodyTerminal
	// BodyEmpty represents an empty production body.
	BodyEmpty
	// BodyInputEnd represents the end of input.
	BodyInputEnd
)

// BodyComp represents a terminal, nonterminal, or empty
// body in a production body string.
type BodyComp struct {
	T BodyCompType
	I int
}

// IsNonTerminal checks if a body component is a nonterminal.
func (c BodyComp) IsNonTerminal() bool {
	return c.T == BodyNonTerminal
}

// IsTerminal checks if a body component is a terminal.
func (c BodyComp) IsTerminal() bool {
	return c.T == BodyTerminal
}

// IsEmpty checks if a body component is a empty component.
func (c BodyComp) IsEmpty() bool {
	return c.T == BodyEmpty
}

// IsInputEnd checks if a body component is the end of input marker.
func (c BodyComp) IsInputEnd() bool {
	return c.T == BodyInputEnd
}
