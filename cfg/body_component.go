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
)

// BodyComp represents a terminal, nonterminal, or empty
// body in a production body string.
type BodyComp struct {
	T BodyCompType
	I int
}
