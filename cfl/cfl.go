package cfl

// Cfl represents a context-free language
type Cfl struct {
	nonTerminals []string
	terminals    []string
	ntTable      map[string]int
	tTable       map[string]int
}
