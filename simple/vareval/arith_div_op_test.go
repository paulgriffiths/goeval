package vareval

import (
	"testing"
)

var divNumberGoodCases = []struct {
	values []value
	result value
}{
	{[]value{intValue{2}}, intValue{2}},
	{[]value{realValue{3.5}}, realValue{3.5}},
	{[]value{intValue{20}, intValue{5}}, intValue{4}},
	{[]value{intValue{20}, intValue{8}}, realValue{2.5}},
	{[]value{intValue{6}, realValue{4.0}}, realValue{1.5}},
	{[]value{realValue{8.0}, intValue{2}}, realValue{4.0}},
	{[]value{realValue{7.5}, realValue{1.5}}, realValue{5.0}},
	{[]value{intValue{60}, intValue{2}, intValue{5}}, intValue{6}},
	{[]value{realValue{60.0}, intValue{2}, intValue{3}}, realValue{10.0}},
	{[]value{intValue{60}, realValue{0.5}, intValue{80}}, realValue{1.5}},
	{[]value{intValue{60}, intValue{3}, realValue{4.0}}, realValue{5.0}},
	{[]value{intValue{60}, realValue{1.5}, realValue{10.0}}, realValue{4.0}},
	{[]value{realValue{60.0}, intValue{4}, realValue{5.0}}, realValue{3.0}},
	{[]value{realValue{60.0}, realValue{1.5}, intValue{4}}, realValue{10.0}},
	{[]value{
		realValue{60.0}, realValue{2.5}, realValue{16.0},
	}, realValue{1.5}},
}

func TestSuccessfulNumericDivOperation(t *testing.T) {
	for n, testCase := range divNumberGoodCases {
		var op expr = testCase.values[0]
		for _, v := range testCase.values[1:] {
			op = divOp{op, v}
		}
		result, err := op.evaluate(nil)
		if err != nil {
			t.Errorf("couldn't evaluate division operation: %v", err)
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

var divVariableGoodCases = []struct {
	number   value
	variable value
	result   value
}{
	{intValue{120}, intValue{30}, intValue{4}},
	{intValue{120}, realValue{0.5}, realValue{240.0}},
	{realValue{120.0}, intValue{3}, realValue{40.0}},
	{realValue{120.0}, realValue{2.5}, realValue{48.0}},
}

func TestSuccessfulVariableDivOperation(t *testing.T) {
	for n, testCase := range divVariableGoodCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := divOp{testCase.number, variableValue{"foobar"}}

		result, err := op.evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate division operation: %v", err)
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

var divNumberBadCases = []struct {
	left  value
	right value
	err   error
}{
	{intValue{42}, boolValue{false}, TypeError},
	{stringValue{"commander_jameson"}, realValue{1.52}, TypeError},
	{stringValue{"cobra_mark_ii"}, boolValue{true}, TypeError},
	{boolValue{false}, boolValue{true}, TypeError},
	{intValue{42}, intValue{0}, DivideByZeroError},
	{intValue{42}, realValue{0.0}, DivideByZeroError},
	{intValue{42}, realValue{-0.0}, DivideByZeroError},
	{realValue{42.0}, intValue{0}, DivideByZeroError},
	{realValue{42.0}, realValue{0.0}, DivideByZeroError},
}

func TestUnsuccessfulNumberDivOperation(t *testing.T) {
	for n, testCase := range divNumberBadCases {
		op := divOp{testCase.left, testCase.right}
		_, err := op.evaluate(nil)
		if err != testCase.err {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, testCase.err)
		}
	}
}

var divVariableBadCases = []struct {
	number   value
	variable value
	err      error
}{
	{intValue{42}, boolValue{false}, TypeError},
	{stringValue{"commander_jameson"}, realValue{1.52}, TypeError},
	{stringValue{"cobra_mark_ii"}, boolValue{true}, TypeError},
	{boolValue{false}, boolValue{true}, TypeError},
	{intValue{42}, intValue{0}, DivideByZeroError},
	{intValue{42}, realValue{0.0}, DivideByZeroError},
	{intValue{42}, realValue{-0.0}, DivideByZeroError},
	{realValue{42.0}, intValue{0}, DivideByZeroError},
	{realValue{42.0}, realValue{0.0}, DivideByZeroError},
}

func TestUnsuccessfulVariableDivOperation(t *testing.T) {
	for n, testCase := range divVariableBadCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := divOp{testCase.number, variableValue{"foobar"}}

		_, err := op.evaluate(table)
		if err != testCase.err {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, testCase.err)
		}
	}
}

func TestUndefinedVariableDivOperation(t *testing.T) {
	table := newTable()
	op := divOp{intValue{42}, variableValue{"foobar"}}

	_, err := op.evaluate(table)
	if err != UnknownIdentifierError {
		t.Errorf("got %v, want %v", err, UnknownIdentifierError)
	}
}
