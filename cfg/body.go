package cfg

// Body represents a production body.
type Body []BodyComp

// IsEmpty checks if a body contains only a single empty component.
func (b Body) IsEmpty() bool {
	return len(b) == 1 && b[0].T == BodyEmpty
}

// IsNonTerminal checks if a body contains only a single nonterminal.
func (b Body) IsNonTerminal() bool {
	return len(b) == 1 && b[0].T == BodyNonTerminal
}

// IsTerminal checks if a body contains only a single terminal.
func (b Body) IsTerminal() bool {
	return len(b) == 1 && b[0].T == BodyTerminal
}

// HasOnlyNonTerminals checks if a body contains only nonterminals.
func (b Body) HasOnlyNonTerminals() bool {
	for _, component := range b {
		if !component.IsNonTerminal() {
			return false
		}
	}
	return true
}

// HasOnlyTerminals checks if a body contains only terminals.
func (b Body) HasOnlyTerminals() bool {
	for _, component := range b {
		if !component.IsTerminal() {
			return false
		}
	}
	return true
}

// IsLast returns true if the provided index refers to the last
// element of the list.
func (b Body) IsLast(n int) bool {
	return n == len(b)-1
}
