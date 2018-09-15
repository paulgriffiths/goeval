package cfg

import (
	"fmt"
	"io"
)

// Cfg represents a context-free grammar.
type Cfg struct {
	nonTerminals []string
	terminals    []string
	ntTable      map[string]int
	tTable       map[string]int
	prods        [][][]BodyComp
}

// outputCfg outputs a representation of the grammar.
func (c *Cfg) outputCfg(writer io.Writer) {
	maxNL := -1
	for _, nt := range c.nonTerminals {
		if len(nt) > maxNL {
			maxNL = len(nt)
		}
	}

	for i, prod := range c.prods {
		for n, body := range prod {
			var s string
			if n == 0 {
				s = fmt.Sprintf("%-[1]*s :", maxNL, c.nonTerminals[i])
			} else {
				s = fmt.Sprintf("%-[1]*s |", maxNL, "")
			}
			writer.Write([]byte(s))

			for _, cmp := range body {
				if cmp.t == BodyNonTerminal {
					s = fmt.Sprintf(" %s", c.nonTerminals[cmp.i])
				} else if cmp.t == BodyTerminal {
					s = fmt.Sprintf(" `%s`", c.terminals[cmp.i])
				} else if cmp.t == BodyEmpty {
					s = " e"
				} else {
					panic("unexpected body component")
				}
				writer.Write([]byte(s))
			}
			writer.Write([]byte("\n"))
		}
	}
}
