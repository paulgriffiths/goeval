package symtab

// Symbol encapsulates a symbol type and associated value.
type Symbol struct {
	Type        ExprType
	valueInt    int
	valueReal   float64
	valueBool   bool
	valueString string
}

// NewInt creates a new symbol value with the specific integer value.
func NewInt(n int) Symbol {
	return Symbol{ExprTypeInt, n, 0, false, ""}
}

// NewReal creates a new symbol value with the specific real value.
func NewReal(f float64) Symbol {
	return Symbol{ExprTypeReal, 0, f, false, ""}
}

// NewBool creates a new symbol value with the specific boolean value.
func NewBool(b bool) Symbol {
	return Symbol{ExprTypeBool, 0, 0, b, ""}
}

// NewString creates a new symbol value with the specific string value.
func NewString(s string) Symbol {
	return Symbol{ExprTypeString, 0, 0, false, s}
}

// SetInt sets an existing symbol value to a new integer value.
func (sym *Symbol) SetInt(n int) {
	sym.Type = ExprTypeInt
	sym.valueInt = n
}

// SetReal sets an existing symbol value to a new real value.
func (sym *Symbol) SetReal(f float64) {
	sym.Type = ExprTypeReal
	sym.valueReal = f
}

// SetBool sets an existing symbol value to a new boolean value.
func (sym *Symbol) SetBool(b bool) {
	sym.Type = ExprTypeBool
	sym.valueBool = b
}

// SetString sets an existing symbol value to a new string value.
func (sym *Symbol) SetString(s string) {
	sym.Type = ExprTypeString
	sym.valueString = s
}

// IntValue retrieves the integer value for a symbol. The function
// panics if the symbol type is not an integer, as this is indicative
// of a programming error.
func (sym *Symbol) IntValue() int {
	if sym.Type != ExprTypeInt {
		panic("symbol type doesn't match value function")
	}
	return sym.valueInt
}

// RealValue retrieves the real value for a symbol. The function
// panics if the symbol type is not a real, as this is indicative
// of a programming error.
func (sym *Symbol) RealValue() float64 {
	if sym.Type != ExprTypeReal {
		panic("symbol type doesn't match value function")
	}
	return sym.valueReal
}

// BoolValue retrieves the boolean value for a symbol. The function
// panics if the symbol type is not a boolean, as this is indicative
// of a programming error.
func (sym *Symbol) BoolValue() bool {
	if sym.Type != ExprTypeBool {
		panic("symbol type doesn't match value function")
	}
	return sym.valueBool
}

// StringValue retrieves the string value for a symbol. The function
// panics if the symbol type is not a string, as this is indicative
// of a programming error.
func (sym *Symbol) StringValue() string {
	if sym.Type != ExprTypeString {
		panic("symbol type doesn't match value function")
	}
	return sym.valueString
}
