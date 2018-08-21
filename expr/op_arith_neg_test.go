package expr

import (
	"testing"
)

var negNumberGoodCases = []struct {
	value  Value
	result Value
}{
	{intValue{2}, intValue{-2}},
	{intValue{-2}, intValue{2}},
}

func TestSuccessfulNumericNegOperation(t *testing.T) {
	for n, testCase := range negNumberGoodCases {
		op := negOp{testCase.value}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("couldn't evaluate exponentiation operation: %v", err)
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

var negNumberBadCases = []struct {
	operand Value
	err     error
}{
	{boolValue{false}, TypeError},
}

func TestUnsuccessfulNumberNegOperation(t *testing.T) {
	for n, testCase := range negNumberBadCases {
		op := negOp{testCase.operand}
		_, err := op.Evaluate(nil)
		if err != testCase.err {
			t.Errorf("case %d, got %v, want %v", n+1,
				err, testCase.err)
		}
	}
}
