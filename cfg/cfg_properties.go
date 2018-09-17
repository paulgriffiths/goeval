package cfg

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
