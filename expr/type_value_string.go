package expr

type stringValue struct {
	value string
}

func NewString(value string) Expr {
	return stringValue{value}
}

func (s stringValue) equality(other stringValue) boolValue {
	return boolValue{s.value == other.value}
}

func (s stringValue) Evaluate(_ *SymTab) (Expr, error) {
	return s, nil
}

func (s stringValue) Equals(other value) bool {
	if !IsString(other) {
		return false
	}
	return s.value == other.(stringValue).value
}

func (s stringValue) String() string {
	return s.value
}
