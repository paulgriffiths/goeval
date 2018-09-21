package cfg

// Unproductive returns a list of unproductive nonterminals.
func (c *Cfg) Unproductive() []int {
	list := []int{}

	// A nonterminal 𝐴 is unproductive if First(𝐴) yields the empty set.

	for i := range c.NonTerminals {
		if c.First(NewNonTerminal(i)).IsEmpty() {
			list = append(list, i)
		}
	}

	return list
}
