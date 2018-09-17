package cfg

// NonTerminalsNullable returns a list of nonterminals which
// are nullable.
func (c *Cfg) NonTerminalsNullable() []int {
	list := []int{}
	for n := range c.Prods {
		if c.inInternal(n, n, make(map[int]bool)) {
			list = append(list, n)
		}
	}
	return list
}

// IsNullable checks if a nonterminal is nullable.
func (c *Cfg) IsNullable(nt int) bool {
	return c.inInternal(nt, nt, make(map[int]bool))
}

func (c *Cfg) inInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	if c.Prods[interNt].HasEmpty() {
		return true
	}

	for _, body := range c.Prods[interNt] {
		nullable := true
		for _, comp := range body {
			if !comp.IsNonTerminal() {
				nullable = false
				break
			}
			if !c.inInternal(nt, comp.I, checked) {
				nullable = false
				break
			}
		}
		if nullable {
			return true
		}
	}
	return false
}
