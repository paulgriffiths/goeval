package expr

import (
	"testing"
)

var mulNumberGoodCases = []struct {
	values []Value
	result Value
}{
	{[]Value{intValue{2}}, intValue{2}},
	{[]Value{realValue{3.5}}, realValue{3.5}},
	{[]Value{intValue{4}, intValue{5}}, intValue{20}},
	{[]Value{intValue{6}, realValue{7.5}}, realValue{45.0}},
	{[]Value{realValue{8.5}, realValue{9}}, realValue{76.5}},
	{[]Value{realValue{10.5}, realValue{11}}, realValue{115.5}},
	{[]Value{intValue{1}, intValue{2}, intValue{3}}, intValue{6}},
	{[]Value{realValue{1.5}, intValue{2}, intValue{3}}, realValue{9.0}},
	{[]Value{intValue{1}, realValue{2.5}, intValue{3}}, realValue{7.5}},
	{[]Value{intValue{1}, intValue{2}, realValue{3.5}}, realValue{7.0}},
	{[]Value{intValue{1}, realValue{2.5}, realValue{3.5}}, realValue{8.75}},
	{[]Value{realValue{1.5}, intValue{2}, realValue{3.5}}, realValue{10.5}},
	{[]Value{realValue{1.5}, realValue{2.5}, intValue{3}}, realValue{11.25}},
	{[]Value{
		realValue{1.5}, realValue{2.5}, realValue{3.5},
	}, realValue{13.125}},
}

func TestSuccessfulNumericMulOperation(t *testing.T) {
	for n, testCase := range mulNumberGoodCases {
		var op Expr = testCase.values[0]
		for _, v := range testCase.values[1:] {
			op = mulOp{op, v}
		}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("couldn't evaluate multiplication operation: %v", err)
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

var mulVariableGoodCases = []struct {
	number   Value
	variable Value
	result   Value
}{
	{intValue{42}, intValue{99}, intValue{4158}},
	{intValue{42}, realValue{99.5}, realValue{4179}},
	{realValue{42.5}, intValue{99}, realValue{4207.5}},
	{realValue{42.5}, realValue{99.5}, realValue{4228.75}},
}

func TestSuccessfulVariableMulOperation(t *testing.T) {
	for n, testCase := range mulVariableGoodCases {
		table := NewTable()
		table.Store("foobar", testCase.variable)
		op := mulOp{testCase.number, variableValue{"foobar"}}

		result, err := op.Evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate multiplication operation: %v", err)
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

var mulNumberBadCases = []struct {
	left  Value
	right Value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulNumberMulOperation(t *testing.T) {
	for n, testCase := range mulNumberBadCases {
		op := mulOp{testCase.left, testCase.right}
		_, err := op.Evaluate(nil)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

var mulVariableBadCases = []struct {
	number   Value
	variable Value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulVariableMulOperation(t *testing.T) {
	for n, testCase := range mulVariableBadCases {
		table := NewTable()
		table.Store("foobar", testCase.variable)
		op := mulOp{testCase.number, variableValue{"foobar"}}

		_, err := op.Evaluate(table)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

func TestUndefinedVariableMulOperation(t *testing.T) {
	table := NewTable()
	op := mulOp{intValue{42}, variableValue{"foobar"}}

	_, err := op.Evaluate(table)
	if err != UnknownIdentifierError {
		t.Errorf("got %v, want %v", err, UnknownIdentifierError)
	}
}
