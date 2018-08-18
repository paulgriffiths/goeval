package vareval

import (
	"testing"
)

var numberCases = []struct {
	values []value
	sum    value
}{
	{[]value{intValue{2}}, intValue{2}},
	{[]value{realValue{3.5}}, realValue{3.5}},
	{[]value{intValue{4}, intValue{5}}, intValue{9}},
	{[]value{intValue{6}, realValue{7.5}}, realValue{13.5}},
	{[]value{realValue{8.5}, realValue{9}}, realValue{17.5}},
	{[]value{realValue{10.5}, realValue{11}}, realValue{21.5}},
	{[]value{intValue{1}, intValue{2}, intValue{3}}, intValue{6}},
	{[]value{realValue{1.5}, intValue{2}, intValue{3}}, realValue{6.5}},
	{[]value{intValue{1}, realValue{2.5}, intValue{3}}, realValue{6.5}},
	{[]value{intValue{1}, intValue{2}, realValue{3.5}}, realValue{6.5}},
	{[]value{intValue{1}, realValue{2.5}, realValue{3.5}}, realValue{7.0}},
	{[]value{realValue{1.5}, intValue{2}, realValue{3.5}}, realValue{7.0}},
	{[]value{realValue{1.5}, realValue{2.5}, intValue{3}}, realValue{7.0}},
	{[]value{
		realValue{1.5}, realValue{2.5}, realValue{3.5},
	}, realValue{7.5}},
}

func TestSuccessfulNumericAddOperation(t *testing.T) {
	for n, testCase := range numberCases {
		var op expr = testCase.values[0]
		for _, v := range testCase.values[1:] {
			op = addOp{op, v}
		}
		sum, err := op.evaluate(nil)
		if err != nil {
			t.Errorf("couldn't evaluate addition operation: %v", err)
			return
		}
		ntsum, ok := sum.(value)
		if !ok {
			t.Errorf("couldn't convert sum to numberValue")
			return
		}
		if !ntsum.equals(testCase.sum) {
			t.Errorf("case %d, got %v, want %v", n+1, ntsum, testCase.sum)
		}

	}
}

var variableCases = []struct {
	number   value
	variable value
	sum      value
}{
	{intValue{42}, intValue{99}, intValue{141}},
	{intValue{42}, realValue{99.5}, realValue{141.5}},
	{realValue{42.5}, intValue{99}, realValue{141.5}},
	{realValue{42.5}, realValue{99.5}, realValue{142.0}},
}

func TestSuccessfulVariableAddOperation(t *testing.T) {
	for n, testCase := range variableCases {
		table := newTable()
		table.store("foobar", testCase.variable)
		op := addOp{testCase.number, variableValue{"foobar"}}

		sum, err := op.evaluate(table)
		if err != nil {
			t.Errorf("couldn't evaluate addition operation: %v", err)
			return
		}
		ntsum, ok := sum.(value)
		if !ok {
			t.Errorf("couldn't convert sum to numberValue")
			return
		}
		if !ntsum.equals(testCase.sum) {
			t.Errorf("case %d, got %v, want %v", n+1, ntsum, testCase.sum)
		}
	}
}
