package cfg

// First returns the set of terminals that begin strings derived
// from the provided string of components.
func (c *Cfg) First(comp ...BodyComp) SetBodyComp {
	set := NewSetBodyComp()
	for _, component := range comp {
		f := c.firstInternal(component, make(map[BodyComp]bool))
		set = set.Union(f)
		if !f.ContainsEmpty() {
			set.DeleteEmpty()
			break
		}
	}
	return set
}

func (c *Cfg) firstInternal(comp BodyComp,
	checked map[BodyComp]bool) SetBodyComp {
	set := NewSetBodyComp()
	if checked[comp] {
		return set
	}
	checked[comp] = true

	if comp.IsTerminal() {
		set.Insert(comp)
		return set
	} else if comp.IsEmpty() {
		return set
	} else if !comp.IsNonTerminal() {
		panic("symbol passed to First neither terminal nor nonterminal")
	}

	for _, body := range c.Prods[comp.I] {
		if body.IsEmpty() {
			set.InsertEmpty()
			continue
		}
		for _, component := range body {
			set = set.Union(c.firstInternal(component, checked))
			if !(component.IsNonTerminal() && c.IsNullable(component.I)) {
				break
			}
		}
	}

	return set
}
