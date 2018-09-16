package cfg

import (
	"fmt"
	"io"
)

// Cfg represents a context-free grammar.
type Cfg struct {
	NonTerminals []string
	Terminals    []string
	NtTable      map[string]int
	TTable       map[string]int
	Prods        [][][]BodyComp
}

// outputCfg outputs a representation of the grammar.
func (c *Cfg) outputCfg(writer io.Writer) {
	maxNL := -1
	for _, nt := range c.NonTerminals {
		if len(nt) > maxNL {
			maxNL = len(nt)
		}
	}

	for i, prod := range c.Prods {
		for n, body := range prod {
			var s string
			if n == 0 {
				s = fmt.Sprintf("%-[1]*s :", maxNL, c.NonTerminals[i])
			} else {
				s = fmt.Sprintf("%-[1]*s |", maxNL, "")
			}
			writer.Write([]byte(s))

			for _, cmp := range body {
				if cmp.t == BodyNonTerminal {
					s = fmt.Sprintf(" %s", c.NonTerminals[cmp.i])
				} else if cmp.t == BodyTerminal {
					s = fmt.Sprintf(" `%s`", c.Terminals[cmp.i])
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
