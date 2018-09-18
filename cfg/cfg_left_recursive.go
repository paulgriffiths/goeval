package cfg

// NonTerminalsImmediatelyLeftRecursive returns a list of nonterminals
// which are immediately left-recursive.
func (c *Cfg) NonTerminalsImmediatelyLeftRecursive() []int {
	list := []int{}
	for n := range c.NonTerminals {
		for _, body := range c.Prods[n] {
			if body[0].IsNonTerminal() && body[0].I == n {
				list = append(list, n)
				break
			}
		}
	}
	return list
}

// NonTerminalsLeftRecursive returns a list of nonterminals which
// are left-recursive.
func (c *Cfg) NonTerminalsLeftRecursive() []int {
	list := []int{}
	for n := range c.NonTerminals {
		if c.lrInternal(n, n, make(map[int]bool)) {
			list = append(list, n)
		}
	}
	return list
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
