package cfg

import (
	"github.com/paulgriffiths/gods/sets"
	"sort"
)

// Unreachable returns a list of unreachable nonterminals.
func (c *Cfg) Unreachable() []int {
	set := sets.NewSetInt(0)
	nextSet := sets.NewSetInt(0)

	for {
		tempSet := sets.NewSetInt()
		for _, nt := range nextSet.Elements() {
			for _, body := range c.Prods[nt] {
				for _, comp := range body {
					if comp.IsNonTerminal() && !set.Contains(comp.I) {
						tempSet.Insert(comp.I)
					}
				}
			}
		}
		if tempSet.IsEmpty() {
			break
		}
		set.Merge(tempSet)
		nextSet = tempSet
	}

	uSet := sets.NewSetInt(intRange(len(c.NonTerminals))...).Difference(set)
	list := []int{}
	for _, element := range uSet.Elements() {
		list = append(list, element)
	}
	sort.Sort(sort.IntSlice(list))
	return list
}
