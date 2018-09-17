package cfg

// NonTerminalsWithCycles returns a list of nonterminals which
// have cycles.
func (c *Cfg) NonTerminalsWithCycles() []int {
	list := []int{}
	for n := range c.NonTerminals {
		if c.hcInternal(n, n, make(map[int]bool)) {
			list = append(list, n)
		}
	}
	return list
}

// HasCycle checks if the grammar contains a cycle.
func (c *Cfg) HasCycle() bool {
	for n := range c.NonTerminals {
		if c.hcInternal(n, n, make(map[int]bool)) {
			return true
		}
	}
	return false
}

func (c *Cfg) hcInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	for _, body := range c.Prods[interNt] {
		if body.IsNonTerminal() {
			if body[0].I == nt {
				return true
			} else if c.hcInternal(nt, body[0].I, checked) {
				return true
			}
		}
	}
	return false
}
