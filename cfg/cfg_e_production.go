package cfg

// HasEProduction checks if the grammar has an e-production.
func (c *Cfg) HasEProduction() bool {
	for _, prod := range c.Prods {
		if prod.HasEmpty() {
			return true
		}
	}
	return false
}

// NonTerminalsWithEProductions returns a list of nonterminals which
// have e-productions.
func (c *Cfg) NonTerminalsWithEProductions() []int {
	list := []int{}
	for n, prod := range c.Prods {
		if prod.HasEmpty() {
			list = append(list, n)
		}
	}
	return list
}
