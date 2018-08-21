package expr

// SymTab implements a multi-level symbol table
type SymTab struct {
	tables []map[string]Value
}

// NewTable creates and returns a new, empty symbol table
func NewTable() *SymTab {
	return &SymTab{[]map[string]Value{make(map[string]Value)}}
}

// Push Pushes a new scope onto the symbol table stack
func (t *SymTab) Push() {
	t.tables = append(t.tables, make(map[string]Value))
}

// Pop removes the innermost scope from the symbol table stack
func (t *SymTab) Pop() {
	if len(t.tables) == 1 {
		panic("symbol table stack underflow")
	}
	t.tables[len(t.tables)-1] = nil
	t.tables = t.tables[:len(t.tables)-1]
}

// Stores a symbol with the specified key
func (t *SymTab) Store(key string, val Value) {
	t.tables[len(t.tables)-1][key] = val
}

// Retrieves the symbol for the specified key, or a dummy key and
// false if the key is not present in the symbol table.
func (t *SymTab) Retrieve(key string) (Value, bool) {
	for i := len(t.tables) - 1; i >= 0; i-- {
		if s, ok := t.tables[i][key]; ok {
			return s, true
		}
	}
	return nil, false
}
