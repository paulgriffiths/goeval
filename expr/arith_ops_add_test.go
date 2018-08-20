package expr

import (
	"testing"
)

var addNumberGoodCases = []struct {
	values []Value
	result Value
}{
	{[]Value{intValue{2}}, intValue{2}},
	{[]Value{realValue{3.5}}, realValue{3.5}},
	{[]Value{intValue{4}, intValue{5}}, intValue{9}},
	{[]Value{intValue{6}, realValue{7.5}}, realValue{13.5}},
	{[]Value{realValue{8.5}, realValue{9}}, realValue{17.5}},
	{[]Value{realValue{10.5}, realValue{11}}, realValue{21.5}},
	{[]Value{intValue{1}, intValue{2}, intValue{3}}, intValue{6}},
	{[]Value{realValue{1.5}, intValue{2}, intValue{3}}, realValue{6.5}},
	{[]Value{intValue{1}, realValue{2.5}, intValue{3}}, realValue{6.5}},
	{[]Value{intValue{1}, intValue{2}, realValue{3.5}}, realValue{6.5}},
	{[]Value{intValue{1}, realValue{2.5}, realValue{3.5}}, realValue{7.0}},
	{[]Value{realValue{1.5}, intValue{2}, realValue{3.5}}, realValue{7.0}},
	{[]Value{realValue{1.5}, realValue{2.5}, intValue{3}}, realValue{7.0}},
	{[]Value{
		realValue{1.5}, realValue{2.5}, realValue{3.5},
	}, realValue{7.5}},
}

func TestSuccessfulNumericAddOperation(t *testing.T) {
	for n, testCase := range addNumberGoodCases {
		var op Expr = testCase.values[0]
		for _, v := range testCase.values[1:] {
			op = addOp{op, v}
		}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate addition operation: %v",
				n, err)
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

var addVariableGoodCases = []struct {
	number   Value
	variable Value
	result   Value
}{
	{intValue{42}, intValue{99}, intValue{141}},
	{intValue{42}, realValue{99.5}, realValue{141.5}},
	{realValue{42.5}, intValue{99}, realValue{141.5}},
	{realValue{42.5}, realValue{99.5}, realValue{142.0}},
}

func TestSuccessfulVariableAddOperation(t *testing.T) {
	for n, testCase := range addVariableGoodCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := addOp{testCase.number, variableValue{"foobar"}}

		result, err := op.Evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate addition operation: %v", err)
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

var addNumberBadCases = []struct {
	left  Value
	right Value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulNumberAddOperation(t *testing.T) {
	for n, testCase := range addNumberBadCases {
		op := addOp{testCase.left, testCase.right}
		_, err := op.Evaluate(nil)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

var addVariableBadCases = []struct {
	number   Value
	variable Value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulVariableAddOperation(t *testing.T) {
	for n, testCase := range addVariableBadCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := addOp{testCase.number, variableValue{"foobar"}}

		_, err := op.Evaluate(table)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

func TestUndefinedVariableAddOperation(t *testing.T) {
	table := newTable()
	op := addOp{intValue{42}, variableValue{"foobar"}}

	_, err := op.Evaluate(table)
	if err != UnknownIdentifierError {
		t.Errorf("got %v, want %v", err, UnknownIdentifierError)
	}
}
