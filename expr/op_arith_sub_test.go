package expr

import (
	"testing"
)

var subNumberGoodCases = []struct {
	values []value
	result value
}{
	{[]value{intValue{2}}, intValue{2}},
	{[]value{realValue{3.5}}, realValue{3.5}},
	{[]value{intValue{4}, intValue{5}}, intValue{-1}},
	{[]value{intValue{6}, realValue{7.5}}, realValue{-1.5}},
	{[]value{realValue{8.5}, realValue{9}}, realValue{-0.5}},
	{[]value{realValue{10.5}, realValue{11}}, realValue{-0.5}},
	{[]value{intValue{1}, intValue{2}, intValue{3}}, intValue{-4}},
	{[]value{realValue{1.5}, intValue{2}, intValue{3}}, realValue{-3.5}},
	{[]value{intValue{1}, realValue{2.5}, intValue{3}}, realValue{-4.5}},
	{[]value{intValue{1}, intValue{2}, realValue{3.5}}, realValue{-4.5}},
	{[]value{intValue{1}, realValue{2.5}, realValue{3.5}}, realValue{-5.0}},
	{[]value{realValue{1.5}, intValue{2}, realValue{3.5}}, realValue{-4.0}},
	{[]value{realValue{1.5}, realValue{2.5}, intValue{3}}, realValue{-4.0}},
	{[]value{
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
		valResult, ok := result.(value)
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
	number   value
	variable value
	result   value
}{
	{intValue{42}, intValue{99}, intValue{-57}},
	{intValue{42}, realValue{99.5}, realValue{-57.5}},
	{realValue{42.5}, intValue{99}, realValue{-56.5}},
	{realValue{42.5}, realValue{99.5}, realValue{-57.0}},
}

func TestSuccessfulVariableSubOperation(t *testing.T) {
	for n, testCase := range subVariableGoodCases {
		table := NewTable()
		table.Store("foobar", testCase.variable)
		op := subOp{testCase.number, variableValue{"foobar"}}

		result, err := op.Evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate subtraction operation: %v", err)
			return
		}
		valResult, ok := result.(value)
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
	left  value
	right value
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
	number   value
	variable value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulVariableSubOperation(t *testing.T) {
	for n, testCase := range subVariableBadCases {
		table := NewTable()
		table.Store("foobar", testCase.variable)
		op := subOp{testCase.number, variableValue{"foobar"}}

		_, err := op.Evaluate(table)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

func TestUndefinedVariableSubOperation(t *testing.T) {
	table := NewTable()
	op := subOp{intValue{42}, variableValue{"foobar"}}

	_, err := op.Evaluate(table)
	if err != UndefinedVariableError {
		t.Errorf("got %v, want %v", err, UndefinedVariableError)
	}
}
