package cfl

import "fmt"

// Cfl represents a context-free grammar.
type Cfl struct {
	nonTerminals []string
	terminals    []string
	ntTable      map[string]int
	tTable       map[string]int
	prods        [][][]BodyComp
}

// printCfl outputs a representation of the grammar.
func (c *Cfl) printCfl() {
	maxNL := -1
	for _, nt := range c.nonTerminals {
		if len(nt) > maxNL {
			maxNL = len(nt)
		}
	}

	for i, prod := range c.prods {
		for n, body := range prod {
			if n == 0 {
				fmt.Printf("%-[1]*s : ", maxNL, c.nonTerminals[i])
			} else {
				fmt.Printf("%-[1]*s | ", maxNL, "")
			}
			for _, cmp := range body {
				if cmp.t == BodyNonTerminal {
					fmt.Printf("%s ", c.nonTerminals[cmp.i])
				} else if cmp.t == BodyTerminal {
					fmt.Printf("`%s` ", c.terminals[cmp.i])
				} else if cmp.t == BodyEmpty {
					fmt.Printf("e ")
				} else {
					panic("unexpected body component")
				}
			}
			fmt.Printf("\n")
		}
	}
}
