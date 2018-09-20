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
		panic("unexpected symbol passed to First")
	}

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

// Follow returns an array of sets, where each set contains the set
// of terminals, or the end-of-input marker, which can follow the
// nonterminal corresponding to the element of the array.
func (c *Cfg) Follow() []SetBodyComp {
	followSets := make([]SetBodyComp, len(c.NonTerminals))
	lengths := make([]int, len(c.NonTerminals))
	for i := 0; i < len(c.NonTerminals); i++ {
		followSets[i] = NewSetBodyComp()
		lengths[i] = -1
	}

	setsChanged := true
	followSets[0].Insert(NewBodyInputEnd())

	for setsChanged {
		for head, prod := range c.Prods {
			for _, body := range prod {
				for i, comp := range body {
					if !comp.IsNonTerminal() {
						continue
					}

					firstContainsEmpty := false
					if !body.IsLast(i) {
						first := c.First(body[i+1:]...)
						if first.ContainsEmpty() {
							firstContainsEmpty = true
							first.DeleteEmpty()
						}
						followSets[comp.I].Merge(first)
					}

					if body.IsLast(i) || firstContainsEmpty {
						followSets[comp.I].Merge(followSets[head])
					}
				}
			}
		}

		setsChanged = false
		for i, set := range followSets {
			if lengths[i] != set.Length() {
				setsChanged = true
			}
			lengths[i] = set.Length()
		}
	}

	return followSets
}
