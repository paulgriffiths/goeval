package lexer

// TerminalList implements a list of terminals found by the lexer.
type TerminalList []Terminal

// Len returns the length of the terminal list.
func (l TerminalList) Len() int {
	return len(l)
}

// Less returns true if element i is less than element j.
func (l TerminalList) Less(i, j int) bool {
	if l[i].S < l[j].S {
		return true
	}
	return l[i].N < l[j].N
}

// Swap swaps elements i and j
func (l TerminalList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
