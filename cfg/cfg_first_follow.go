package cfg

// First returns the set of terminals that begin strings derived
// from the provided string of components.
func (c *Cfg) First(comp ...BodyComp) SetBodyComp {

	// First(𝒶𝛽) is simply 𝒶, and 𝜀 obviously has no content so
	// return an empty set. For a single nonterminal, return the
	// precomputed set.

	if comp[0].IsTerminal() {
		return NewSetBodyComp(comp[0])
	}

	if len(comp) == 1 {
		if comp[0].IsEmpty() {
			return NewSetBodyComp()
		} else if comp[0].IsNonTerminal() {
			return c.Firsts[comp[0].I]
		}
		panic("unexpected symbol passed to First")
	}

	// For a string of 𝑋1𝑋2...𝑋n, start with the non-𝜀 symbols of
	// 𝑋1. If 𝜀 is in 𝑋1, then repeat with 𝑋2, and so on. If we
	// reach 𝑋n and 𝜀 is in 𝑋n, then 𝜀 is also in First(𝑋1𝑋2...𝑋n).

	set := NewSetBodyComp()
	for _, component := range comp {
		f := c.First(component)
		set.Merge(f)
		if !f.ContainsEmpty() {
			set.DeleteEmpty()
			break
		}
	}
	return set
}

// calcFirsts calculates the First sets for each nonterminal.
func (c *Cfg) calcFirsts() {
	c.Firsts = make([]SetBodyComp, len(c.NonTerminals))
	lengths := make([]int, len(c.NonTerminals))
	for i := 0; i < len(c.NonTerminals); i++ {
		c.Firsts[i] = NewSetBodyComp()
		lengths[i] = -1
	}

	setsChanged := true

	for setsChanged {

		// Calculate first sets for each nonterminal.

		for n := range c.NonTerminals {
			component := NewNonTerminal(n)
			f := c.firstInternal(component, make(map[BodyComp]bool))
			c.Firsts[n].Merge(f)
		}

		// We need to apply the rules until nothing can be added to
		// any follow set, which will be the case if we've applied
		// the rules to every production and none of the follow sets
		// have changed since we started.

		setsChanged = false
		for i, set := range c.Firsts {
			if lengths[i] != set.Length() {
				setsChanged = true
			}
			lengths[i] = set.Length()
		}
	}
}

// firstInternal performs one complete cycle of first set
// computation rules for a given symbol.
func (c *Cfg) firstInternal(comp BodyComp,
	checked map[BodyComp]bool) SetBodyComp {

	set := NewSetBodyComp()

	// First(𝒶) is simply 𝒶, and 𝜀 obviously has no content.

	if comp.IsTerminal() {
		set.Insert(comp)
		return set
	} else if comp.IsEmpty() {
		return set
	} else if !comp.IsNonTerminal() {
		panic("unexpected symbol passed to First")
	}

	if checked[comp] {

		// We already calculated first for this nonterminal,
		// so return the empty set and avoid an infinite loop.

		return set
	}
	checked[comp] = true

	for _, body := range c.Prods[comp.I] {
		if body.IsEmpty() {
			set.InsertEmpty()
			continue
		}

		for _, component := range body {
			set.Merge(c.firstInternal(component, checked))
			if !(component.IsNonTerminal() && c.IsNullable(component.I)) {
				break
			}
		}
	}

	return set
}

// calcFollows calculates the Follow set for each nonterminal, where
// the Follow set contains the set of terminals, or the end-of-input
// marker, which can follow that nonterminal.
func (c *Cfg) calcFollows() {
	c.Follows = make([]SetBodyComp, len(c.NonTerminals))
	lengths := make([]int, len(c.NonTerminals))
	for i := 0; i < len(c.NonTerminals); i++ {
		c.Follows[i] = NewSetBodyComp()
		lengths[i] = -1
	}

	setsChanged := true

	// End of input can always follow the start symbol.

	c.Follows[0].Insert(NewBodyInputEnd())

	for setsChanged {
		for head, prod := range c.Prods {
			for _, body := range prod {
				for i, comp := range body {

					if !comp.IsNonTerminal() {

						// We're only calculating follow for
						// nonterminals, so skip anything else.

						continue
					}

					if !body.IsLast(i) {

						// If 𝛢→𝛼𝛣𝛽, then everything in first(𝛽)
						// is in follow(𝛣) except 𝜀, since it's not a
						// terminal.

						first := c.First(body[i+1:]...).Copy()

						if first.ContainsEmpty() {

							// If first(𝛽) derives 𝜀, then 𝛣 can appear
							// at the end of an 𝛢 production, therefore
							// anything that follows 𝛢 can also follow 𝛣.

							c.Follows[comp.I].Merge(c.Follows[head])

							// 𝜀 itself can't follow 𝛣, since it's not a
							// terminal, so remove it if it's present.

							first.DeleteEmpty()
						}

						c.Follows[comp.I].Merge(first)

					} else if body.IsLast(i) {

						// If 𝛢→𝛼𝛣, then 𝛣 can appear at the end of an
						// 𝛢 production, therefore anything that follows
						// 𝛢 can also follow 𝛣.

						c.Follows[comp.I].Merge(c.Follows[head])
					}
				}
			}
		}

		// We need to apply the rules until nothing can be added to
		// any follow set, which will be the case if we've applied
		// the rules to every production and none of the follow sets
		// have changed since we started.

		setsChanged = false
		for i, set := range c.Follows {
			if lengths[i] != set.Length() {
				setsChanged = true
			}
			lengths[i] = set.Length()
		}
	}
}
