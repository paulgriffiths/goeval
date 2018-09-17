package cfg

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
