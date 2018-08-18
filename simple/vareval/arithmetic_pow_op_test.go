package vareval

import (
	"testing"
)

var powNumberGoodCases = []struct {
	values []value
	result value
}{
	{[]value{intValue{2}}, intValue{2}},
	{[]value{realValue{3.5}}, realValue{3.5}},
	{[]value{intValue{2}, intValue{3}}, realValue{8.0}},
	{[]value{intValue{2}, realValue{3.0}}, realValue{8.0}},
	{[]value{realValue{2.0}, intValue{3}}, realValue{8.0}},
	{[]value{realValue{16.0}, realValue{0.5}}, realValue{4.0}},
	{[]value{intValue{2}, intValue{3}, intValue{4}}, realValue{4096.0}},
}

func TestSuccessfulNumericPowOperation(t *testing.T) {
	for n, testCase := range powNumberGoodCases {
		var op expr = testCase.values[0]
		for _, v := range testCase.values[1:] {
			op = powOp{op, v}
		}
		result, err := op.evaluate(nil)
		if err != nil {
			t.Errorf("couldn't evaluate exponentiation operation: %v", err)
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

var powVariableGoodCases = []struct {
	number   value
	variable value
	result   value
}{
	{intValue{2}, intValue{3}, realValue{8.0}},
	{intValue{2}, realValue{3.0}, realValue{8.0}},
	{realValue{2.0}, intValue{3}, realValue{8.0}},
	{realValue{16.0}, realValue{0.5}, realValue{4.0}},
}

func TestSuccessfulVariablePowOperation(t *testing.T) {
	for n, testCase := range powVariableGoodCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := powOp{testCase.number, variableValue{"foobar"}}

		result, err := op.evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate exponentiation operation: %v", err)
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

var powNumberBadCases = []struct {
	left  value
	right value
	err   error
}{
	{intValue{42}, boolValue{false}, TypeError},
	{stringValue{"commander_jameson"}, realValue{1.52}, TypeError},
	{stringValue{"cobra_mark_ii"}, boolValue{true}, TypeError},
	{boolValue{false}, boolValue{true}, TypeError},
	{realValue{-42.0}, realValue{0.5}, DomainError},
	{realValue{-42.0}, realValue{1.5}, DomainError},
}

func TestUnsuccessfulNumberPowOperation(t *testing.T) {
	for n, testCase := range powNumberBadCases {
		op := powOp{testCase.left, testCase.right}
		_, err := op.evaluate(nil)
		if err != testCase.err {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, testCase.err)
		}
	}
}

var powVariableBadCases = []struct {
	number   value
	variable value
	err      error
}{
	{intValue{42}, boolValue{false}, TypeError},
	{stringValue{"commander_jameson"}, realValue{1.52}, TypeError},
	{stringValue{"cobra_mark_ii"}, boolValue{true}, TypeError},
	{boolValue{false}, boolValue{true}, TypeError},
	{realValue{-42.0}, realValue{0.5}, DomainError},
	{realValue{-42.0}, realValue{1.5}, DomainError},
}

func TestUnsuccessfulVariablePowOperation(t *testing.T) {
	for n, testCase := range powVariableBadCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := powOp{testCase.number, variableValue{"foobar"}}

		_, err := op.evaluate(table)
		if err != testCase.err {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, testCase.err)
		}
	}
}

func TestUndefinedVariablePowOperation(t *testing.T) {
	table := newTable()
	op := powOp{intValue{42}, variableValue{"foobar"}}

	_, err := op.evaluate(table)
	if err != UnknownIdentifierError {
		t.Errorf("got %v, want %v", err, UnknownIdentifierError)
	}
}
