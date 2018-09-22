package cfg

import "github.com/paulgriffiths/gods/sets"

// First returns the set of terminals that begin strings derived
// from the provided string of components.
func (c *Cfg) First(comp ...BodyComp) SetBodyComp {
	if c.firsts == nil {
		c.calcFirsts()
	}

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
			return c.firsts[comp[0].I]
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
	nullables := c.calcNullables()
	c.firsts = make([]SetBodyComp, len(c.NonTerminals))
	lengths := make([]int, len(c.NonTerminals))
	for i := 0; i < len(c.NonTerminals); i++ {
		c.firsts[i] = NewSetBodyComp()
		lengths[i] = -1
	}

	setsChanged := true

	for setsChanged {

		// Complete one First set calculation cycle for each nonterminal.

		for n := range c.NonTerminals {
			component := NewNonTerminal(n)
			f := c.firstInternal(component, nullables,
				make(map[BodyComp]bool))
			c.firsts[n].Merge(f)
		}

		// We need to apply the rules until nothing can be added to
		// any First set, which will be the case if we've applied
		// the rules to every production and none of the First sets
		// have changed since we started.

		setsChanged = false
		for i, set := range c.firsts {
			if lengths[i] != set.Length() {
				setsChanged = true
			}
			lengths[i] = set.Length()
		}
	}
}

// firstInternal performs one complete cycle of First set
// computation rules for a given symbol.
func (c *Cfg) firstInternal(comp BodyComp, nullables sets.SetInt,
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

		// We already calculated First for this nonterminal,
		// so return the empty set and avoid an infinite loop.

		return set
	}
	checked[comp] = true

	for _, body := range c.Prods[comp.I] {
		if body.IsEmptyBody() {
			set.InsertEmpty()
			continue
		}

		for _, comp := range body {
			set.Merge(c.firstInternal(comp, nullables, checked))
			if !(comp.IsNonTerminal() && nullables.Contains(comp.I)) {
				set.DeleteEmpty()
				break
			}
		}
	}

	return set
}

// calcNullables returns the set of nonterminals which can derive 𝜀.
func (c *Cfg) calcNullables() sets.SetInt {
	nullable := sets.NewSetInt()
	newNulls := sets.NewSetInt()

	// Add to set any nonterminal 𝐴 where 𝐴 → 𝜀 is a production.

	for n, prod := range c.Prods {
		if prod.HasEmpty() {
			nullable.Insert(n)
		}
	}

	// Identify any remaining indirectly nullable nonterminals.

	for !nullable.Equals(newNulls) {
		newNulls.Merge(nullable)
		nullable.Merge(newNulls)

		for n, prod := range c.Prods {

			// If this nonterminal is already in the set, don't
			// waste time checking it again.

			if newNulls.Contains(n) {
				continue
			}

			for _, body := range prod {

				// If the production body contains a terminal, it
				// can't be nullable, so continue to the next. We
				// already identified any 𝐴 → 𝜀 productions.

				if !body.HasOnlyNonTerminals() {
					continue
				}

				// The production derives 𝜀 if and only if each
				// nonterminal in the production derives 𝜀. If the
				// production derives 𝜀, the whole nonterminal can
				// derive 𝜀 and there's no need to check further.

				derivesEmpty := true
				for _, comp := range body {
					if !newNulls.Contains(comp.I) {
						derivesEmpty = false
						break
					}
				}

				if derivesEmpty {
					newNulls.Insert(n)
					break
				}
			}
		}
	}

	return nullable
}
