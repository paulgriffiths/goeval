package pp

import (
	"github.com/paulgriffiths/goeval/cfg"
)

type ppTable [][][]cfg.Body

func makePPTable(grammar *cfg.Cfg) ppTable {
	numTerms := len(grammar.Terminals) + 1 // +1 for end of input marker
	numNonTerms := len(grammar.NonTerminals)

	newTable := make([][][]cfg.Body, numNonTerms)
	for i := 0; i < numNonTerms; i++ {
		newTable[i] = make([][]cfg.Body, numTerms)
		for j := 0; j < numTerms; j++ {
			newTable[i][j] = []cfg.Body{}
		}
	}

	return newTable
}

type Pp struct {
	g     *cfg.Cfg
	table ppTable
}

func NewPp(grammar *cfg.Cfg) *Pp {
	table := makePPTable(grammar)
	newParser := Pp{grammar, table}
	buildTable(grammar, table)
	return &newParser
}

func buildTable(g *cfg.Cfg, m ppTable) {

	// Loop through all grammar productions 𝛢 → 𝛼

	for nt, prod := range g.Prods {
		for _, body := range prod {
			first := g.First(body...)

			// For each terminal 𝑎 in First(𝛼), add 𝛢 → 𝛼 to m[𝛢,𝑎]

			for _, c := range first.Elements() {
				if c.IsTerminal() {
					m[nt][c.I] = append(m[nt][c.I], body)
				}
			}

			// If First(𝛼) contains 𝜀 (or if 𝛼 = 𝜀) then for each
			// terminal 𝑏 in Follow(𝛢), add 𝛢 → 𝛼 to m[𝛢,b]. If
			// First(𝛼) contains 𝜀 (or if 𝛼 = 𝜀) and Follow(𝛢)
			// contains the end-of-input marker, then add 𝛢 → 𝛼
			// to m[𝛢,$], too.

			follow := g.Follow(nt)
			if first.ContainsEmpty() || body.IsEmpty() {
				for _, c := range follow.Elements() {
					if c.IsTerminal() {
						m[nt][c.I] = append(m[nt][c.I], body)
					}
				}

				if follow.ContainsEndOfInput() {
					eoi := len(g.Terminals)
					m[nt][eoi] = append(m[nt][eoi], body)
				}
			}
		}
	}
}
