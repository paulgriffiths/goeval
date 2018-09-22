package cfg

// NonTerminalsNullable returns a list of nonterminals which
// are nullable.
func (c *Cfg) NonTerminalsNullable() []int {
	list := []int{}
	for n := range c.NonTerminals {
		if c.IsNullable(n) {
			list = append(list, n)
		}
	}
	return list
}

// IsNullable checks if a nonterminal is nullable.
func (c *Cfg) IsNullable(nt int) bool {
	return c.First(NewNonTerminal(nt)).ContainsEmpty()
}
