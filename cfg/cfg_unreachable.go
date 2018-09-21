package cfg

import (
	"github.com/paulgriffiths/gods/sets"
	"sort"
)

// Unreachable returns a list of unreachable nonterminals.
func (c *Cfg) Unreachable() []int {

	// Begin search from start symbol 𝑆, which is always reachable.

	reachable := sets.NewSetInt(0)
	visitNext := sets.NewSetInt(0)

	for {

		// Loop through all productions of all nonterminals reached
		// on the last loop iteration, and aggregate any nonterminals
		// reachable via those productions which have not previously
		// been reached.

		reached := sets.NewSetInt()
		for _, nt := range visitNext.Elements() {
			for _, body := range c.Prods[nt] {
				for _, comp := range body {
					if comp.IsNonTerminal() && !reachable.Contains(comp.I) {
						reached.Insert(comp.I)
					}
				}
			}
		}

		// If we didn't reach any nonterminals which had not already
		// been reached, then no additional loop iterations will
		// reach any more, so terminate loop.

		if reached.IsEmpty() {
			break
		}

		// Otherwise add the newly-reached nonterminals to the
		// reachable set and loop again.

		reachable.Merge(reached)
		visitNext = reached
	}

	// The set of unreachable nonterminals is the set difference
	// between the set of all nonterminals and the set of reachable
	// nonterminals.

	list := c.NonTerminalsSet().Difference(reachable).Elements()
	sort.Sort(sort.IntSlice(list))
	return list
}
