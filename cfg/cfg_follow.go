package cfg

// Follow calculates the Follow set for the given nonterminal, where
// the Follow set contains the set of terminals, or the end-of-input
// marker, which can follow that nonterminal.
func (c *Cfg) Follow(n int) SetBodyComp {
	if c.follows == nil {
		c.calcFollows()
	}
	return c.follows[n]
}

// calcFollows calculates the Follow set for each nonterminal.
func (c *Cfg) calcFollows() {
	c.follows = make([]SetBodyComp, len(c.NonTerminals))
	lengths := make([]int, len(c.NonTerminals))
	for i := 0; i < len(c.NonTerminals); i++ {
		c.follows[i] = NewSetBodyComp()
		lengths[i] = -1
	}

	setsChanged := true

	// End of input can always follow the start symbol.

	c.follows[0].Insert(NewBodyInputEnd())

	for setsChanged {
		for head, prod := range c.Prods {
			for _, body := range prod {
				for i, comp := range body {

					if !comp.IsNonTerminal() {

						// We're only calculating Follow for
						// nonterminals, so skip anything else.

						continue
					}

					if !body.IsLast(i) {

						// If 𝛢→𝛼𝛣𝛽, then everything in first(𝛽)
						// is in Follow(𝛣) except 𝜀, since it's not a
						// terminal.

						first := c.First(body[i+1:]...).Copy()

						if first.ContainsEmpty() {

							// If First(𝛽) derives 𝜀, then 𝛣 can appear
							// at the end of an 𝛢 production, therefore
							// anything that follows 𝛢 can also follow 𝛣.

							c.follows[comp.I].Merge(c.follows[head])

							// 𝜀 itself can't follow 𝛣, since it's not a
							// terminal, so remove it if it's present.

							first.DeleteEmpty()
						}

						c.follows[comp.I].Merge(first)

					} else if body.IsLast(i) {

						// If 𝛢→𝛼𝛣, then 𝛣 can appear at the end of an
						// 𝛢 production, therefore anything that follows
						// 𝛢 can also follow 𝛣.

						c.follows[comp.I].Merge(c.follows[head])
					}
				}
			}
		}

		// We need to apply the rules until nothing can be added to
		// any Follow set, which will be the case if we've applied
		// the rules to every production and none of the Follow sets
		// have changed since we started.

		setsChanged = false
		for i, set := range c.follows {
			if lengths[i] != set.Length() {
				setsChanged = true
			}
			lengths[i] = set.Length()
		}
	}
}
