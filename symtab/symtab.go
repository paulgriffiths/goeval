package symtab

// SymTab implements a multi-level symbol table
type SymTab struct {
    tables []map[string]Symbol
    top int
}

// NewTable creates and returns a new, empty symbol table
func NewTable() *SymTab {
    var s SymTab
    s.tables = []map[string]Symbol{make(map[string]Symbol)}
    s.top = 0
    return &s
}

// Push pushes a new scope onto the symbol table stack
func (t *SymTab) Push() {
    t.tables = append(t.tables, make(map[string]Symbol))
    t.top++
}

// Pop removes the innermost scope from the symbol table stack
func (t *SymTab) Pop() {
    if t.top < 1 {
        panic("symbol table stack underflow")
    }
    t.tables[t.top] = nil
    t.top--
}

// Stores a symbol with the specified key
func (t *SymTab) Store(key string, val Symbol) {
    t.tables[t.top][key] = val
}

// Retrieves the symbol for the specified key, or a dummy key and
// false if the key is not present in the symbol table.
func (t *SymTab) Retrieve(key string) (Symbol, bool) {
    for n := t.top; n >= 0; n-- {
        s, ok := t.tables[n][key]
        if ok {
            return s, true
        }
    }
    return NewInt(0), false
}
