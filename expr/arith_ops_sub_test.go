package expr

import (
	"testing"
)

var subNumberGoodCases = []struct {
	values []Value
	result Value
}{
	{[]Value{intValue{2}}, intValue{2}},
	{[]Value{realValue{3.5}}, realValue{3.5}},
	{[]Value{intValue{4}, intValue{5}}, intValue{-1}},
	{[]Value{intValue{6}, realValue{7.5}}, realValue{-1.5}},
	{[]Value{realValue{8.5}, realValue{9}}, realValue{-0.5}},
	{[]Value{realValue{10.5}, realValue{11}}, realValue{-0.5}},
	{[]Value{intValue{1}, intValue{2}, intValue{3}}, intValue{-4}},
	{[]Value{realValue{1.5}, intValue{2}, intValue{3}}, realValue{-3.5}},
	{[]Value{intValue{1}, realValue{2.5}, intValue{3}}, realValue{-4.5}},
	{[]Value{intValue{1}, intValue{2}, realValue{3.5}}, realValue{-4.5}},
	{[]Value{intValue{1}, realValue{2.5}, realValue{3.5}}, realValue{-5.0}},
	{[]Value{realValue{1.5}, intValue{2}, realValue{3.5}}, realValue{-4.0}},
	{[]Value{realValue{1.5}, realValue{2.5}, intValue{3}}, realValue{-4.0}},
	{[]Value{
		realValue{1.5}, realValue{2.5}, realValue{3.5},
	}, realValue{-4.5}},
}

func TestSuccessfulNumericSubOperation(t *testing.T) {
	for n, testCase := range subNumberGoodCases {
		var op Expr = testCase.values[0]
		for _, v := range testCase.values[1:] {
			op = subOp{op, v}
		}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("couldn't evaluate subtraction operation: %v", err)
			return
		}
		valResult, ok := result.(Value)
		if !ok {
			t.Errorf("couldn't convert result to value")
			return
		}
		if !valResult.Equals(testCase.result) {
			t.Errorf("case %d, got %v, want %v", n+1,
				valResult, testCase.result)
		}

	}
}

var subVariableGoodCases = []struct {
	number   Value
	variable Value
	result   Value
}{
	{intValue{42}, intValue{99}, intValue{-57}},
	{intValue{42}, realValue{99.5}, realValue{-57.5}},
	{realValue{42.5}, intValue{99}, realValue{-56.5}},
	{realValue{42.5}, realValue{99.5}, realValue{-57.0}},
}

func TestSuccessfulVariableSubOperation(t *testing.T) {
	for n, testCase := range subVariableGoodCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := subOp{testCase.number, variableValue{"foobar"}}

		result, err := op.Evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate subtraction operation: %v", err)
			return
		}
		valResult, ok := result.(Value)
		if !ok {
			t.Errorf("couldn't convert result to value")
			return
		}
		if !valResult.Equals(testCase.result) {
			t.Errorf("case %d, got %v, want %v", n+1,
				valResult, testCase.result)
		}
	}
}

var subNumberBadCases = []struct {
	left  Value
	right Value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulNumberSubOperation(t *testing.T) {
	for n, testCase := range subNumberBadCases {
		op := subOp{testCase.left, testCase.right}
		_, err := op.Evaluate(nil)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

var subVariableBadCases = []struct {
	number   Value
	variable Value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulVariableSubOperation(t *testing.T) {
	for n, testCase := range subVariableBadCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := subOp{testCase.number, variableValue{"foobar"}}

		_, err := op.Evaluate(table)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

func TestUndefinedVariableSubOperation(t *testing.T) {
	table := newTable()
	op := subOp{intValue{42}, variableValue{"foobar"}}

	_, err := op.Evaluate(table)
	if err != UnknownIdentifierError {
		t.Errorf("got %v, want %v", err, UnknownIdentifierError)
	}
}
