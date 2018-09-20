package cfg

// Unproductive returns a list of unproductive nonterminals.
func (c *Cfg) Unproductive() []int {
	list := []int{}

	for i := 0; i < len(c.NonTerminals); i++ {
		first := c.First(NewNonTerminal(i))
		if first.Length() == 0 {
			list = append(list, i)
		}
	}
	return list
}
