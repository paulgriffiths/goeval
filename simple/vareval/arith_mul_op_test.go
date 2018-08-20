package vareval

import (
	"testing"
)

var mulNumberGoodCases = []struct {
	values []value
	result value
}{
	{[]value{intValue{2}}, intValue{2}},
	{[]value{realValue{3.5}}, realValue{3.5}},
	{[]value{intValue{4}, intValue{5}}, intValue{20}},
	{[]value{intValue{6}, realValue{7.5}}, realValue{45.0}},
	{[]value{realValue{8.5}, realValue{9}}, realValue{76.5}},
	{[]value{realValue{10.5}, realValue{11}}, realValue{115.5}},
	{[]value{intValue{1}, intValue{2}, intValue{3}}, intValue{6}},
	{[]value{realValue{1.5}, intValue{2}, intValue{3}}, realValue{9.0}},
	{[]value{intValue{1}, realValue{2.5}, intValue{3}}, realValue{7.5}},
	{[]value{intValue{1}, intValue{2}, realValue{3.5}}, realValue{7.0}},
	{[]value{intValue{1}, realValue{2.5}, realValue{3.5}}, realValue{8.75}},
	{[]value{realValue{1.5}, intValue{2}, realValue{3.5}}, realValue{10.5}},
	{[]value{realValue{1.5}, realValue{2.5}, intValue{3}}, realValue{11.25}},
	{[]value{
		realValue{1.5}, realValue{2.5}, realValue{3.5},
	}, realValue{13.125}},
}

func TestSuccessfulNumericMulOperation(t *testing.T) {
	for n, testCase := range mulNumberGoodCases {
		var op expr = testCase.values[0]
		for _, v := range testCase.values[1:] {
			op = mulOp{op, v}
		}
		result, err := op.evaluate(nil)
		if err != nil {
			t.Errorf("couldn't evaluate multiplication operation: %v", err)
			return
		}
		valResult, ok := result.(value)
		if !ok {
			t.Errorf("couldn't convert result to value")
			return
		}
		if !valResult.equals(testCase.result) {
			t.Errorf("case %d, got %v, want %v", n+1,
				valResult, testCase.result)
		}

	}
}

var mulVariableGoodCases = []struct {
	number   value
	variable value
	result   value
}{
	{intValue{42}, intValue{99}, intValue{4158}},
	{intValue{42}, realValue{99.5}, realValue{4179}},
	{realValue{42.5}, intValue{99}, realValue{4207.5}},
	{realValue{42.5}, realValue{99.5}, realValue{4228.75}},
}

func TestSuccessfulVariableMulOperation(t *testing.T) {
	for n, testCase := range mulVariableGoodCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := mulOp{testCase.number, variableValue{"foobar"}}

		result, err := op.evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate multiplication operation: %v", err)
			return
		}
		valResult, ok := result.(value)
		if !ok {
			t.Errorf("couldn't convert result to value")
			return
		}
		if !valResult.equals(testCase.result) {
			t.Errorf("case %d, got %v, want %v", n+1,
				valResult, testCase.result)
		}
	}
}

var mulNumberBadCases = []struct {
	left  value
	right value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulNumberMulOperation(t *testing.T) {
	for n, testCase := range mulNumberBadCases {
		op := mulOp{testCase.left, testCase.right}
		_, err := op.evaluate(nil)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

var mulVariableBadCases = []struct {
	number   value
	variable value
}{
	{intValue{42}, boolValue{false}},
	{stringValue{"commander_jameson"}, realValue{1.52}},
	{stringValue{"cobra_mark_ii"}, boolValue{true}},
	{boolValue{false}, boolValue{true}},
}

func TestUnsuccessfulVariableMulOperation(t *testing.T) {
	for n, testCase := range mulVariableBadCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := mulOp{testCase.number, variableValue{"foobar"}}

		_, err := op.evaluate(table)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, TypeError)
		}
	}
}

func TestUndefinedVariableMulOperation(t *testing.T) {
	table := newTable()
	op := mulOp{intValue{42}, variableValue{"foobar"}}

	_, err := op.evaluate(table)
	if err != UnknownIdentifierError {
		t.Errorf("got %v, want %v", err, UnknownIdentifierError)
	}
}
