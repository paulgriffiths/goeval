package cfg

import (
	"fmt"
	"io"
)

// outputCfg outputs a representation of the grammar.
func (c *Cfg) outputCfg(writer io.Writer) {
	maxNL := c.MaxNonTerminalNameLength()

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
				switch {
				case cmp.IsNonTerminal():
					s = fmt.Sprintf(" %s", c.NonTerminals[cmp.I])
				case cmp.IsTerminal():
					s = fmt.Sprintf(" `%s`", c.Terminals[cmp.I])
				case cmp.IsEmpty():
					s = " e"
				default:
					panic("unexpected body component")
				}
				writer.Write([]byte(s))
			}
			writer.Write([]byte("\n"))
		}
	}
}
