package expr

import "testing"

func TestEqualityOp(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		equal       boolValue
	}{
		{intValue{3}, intValue{3}, boolValue{true}},
		{intValue{3}, intValue{-3}, boolValue{false}},
		{intValue{3}, realValue{3.0}, boolValue{true}},
		{intValue{3}, realValue{-3.0}, boolValue{false}},
		{realValue{3.0}, intValue{3}, boolValue{true}},
		{realValue{3.0}, intValue{-3}, boolValue{false}},
		{realValue{3.0}, realValue{3.0}, boolValue{true}},
		{realValue{3.0}, realValue{-3.0}, boolValue{false}},
	}
	for n, testCase := range testCases {
		op := equalityOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.equal {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.equal)
		}
	}
}

func TestNonEqualityOp(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		equal       boolValue
	}{
		{intValue{3}, intValue{3}, boolValue{false}},
		{intValue{3}, intValue{-3}, boolValue{true}},
		{intValue{3}, realValue{3.0}, boolValue{false}},
		{intValue{3}, realValue{-3.0}, boolValue{true}},
		{realValue{3.0}, intValue{3}, boolValue{false}},
		{realValue{3.0}, intValue{-3}, boolValue{true}},
		{realValue{3.0}, realValue{3.0}, boolValue{false}},
		{realValue{3.0}, realValue{-3.0}, boolValue{true}},
	}
	for n, testCase := range testCases {
		op := nonEqualityOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.equal {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.equal)
		}
	}
}

func TestLessThanOp(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		equal       boolValue
	}{
		{intValue{3}, intValue{3}, boolValue{false}},
		{intValue{3}, intValue{-3}, boolValue{false}},
		{intValue{3}, intValue{6}, boolValue{true}},
		{intValue{3}, realValue{3.0}, boolValue{false}},
		{intValue{3}, realValue{-3.0}, boolValue{false}},
		{intValue{3}, realValue{6.0}, boolValue{true}},
		{realValue{3.0}, intValue{3}, boolValue{false}},
		{realValue{3.0}, intValue{-3}, boolValue{false}},
		{realValue{3.0}, intValue{6}, boolValue{true}},
		{realValue{3.0}, realValue{3.0}, boolValue{false}},
		{realValue{3.0}, realValue{-3.0}, boolValue{false}},
		{realValue{3.0}, realValue{6.0}, boolValue{true}},
	}
	for n, testCase := range testCases {
		op := lessThanOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.equal {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.equal)
		}
	}
}

func TestGreaterThanOp(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		equal       boolValue
	}{
		{intValue{3}, intValue{3}, boolValue{false}},
		{intValue{3}, intValue{-3}, boolValue{true}},
		{intValue{3}, intValue{6}, boolValue{false}},
		{intValue{3}, realValue{3.0}, boolValue{false}},
		{intValue{3}, realValue{-3.0}, boolValue{true}},
		{intValue{3}, realValue{6.0}, boolValue{false}},
		{realValue{3.0}, intValue{3}, boolValue{false}},
		{realValue{3.0}, intValue{-3}, boolValue{true}},
		{realValue{3.0}, intValue{6}, boolValue{false}},
		{realValue{3.0}, realValue{3.0}, boolValue{false}},
		{realValue{3.0}, realValue{-3.0}, boolValue{true}},
		{realValue{3.0}, realValue{6.0}, boolValue{false}},
	}
	for n, testCase := range testCases {
		op := greaterThanOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.equal {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.equal)
		}
	}
}

func TestLessThanOrEqualOp(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		equal       boolValue
	}{
		{intValue{3}, intValue{3}, boolValue{true}},
		{intValue{3}, intValue{-3}, boolValue{false}},
		{intValue{3}, intValue{6}, boolValue{true}},
		{intValue{3}, realValue{3.0}, boolValue{true}},
		{intValue{3}, realValue{-3.0}, boolValue{false}},
		{intValue{3}, realValue{6.0}, boolValue{true}},
		{realValue{3.0}, intValue{3}, boolValue{true}},
		{realValue{3.0}, intValue{-3}, boolValue{false}},
		{realValue{3.0}, intValue{6}, boolValue{true}},
		{realValue{3.0}, realValue{3.0}, boolValue{true}},
		{realValue{3.0}, realValue{-3.0}, boolValue{false}},
		{realValue{3.0}, realValue{6.0}, boolValue{true}},
	}
	for n, testCase := range testCases {
		op := lessThanOrEqualOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.equal {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.equal)
		}
	}
}

func TestGreaterThanOrEqualOp(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		equal       boolValue
	}{
		{intValue{3}, intValue{3}, boolValue{true}},
		{intValue{3}, intValue{-3}, boolValue{true}},
		{intValue{3}, intValue{6}, boolValue{false}},
		{intValue{3}, realValue{3.0}, boolValue{true}},
		{intValue{3}, realValue{-3.0}, boolValue{true}},
		{intValue{3}, realValue{6.0}, boolValue{false}},
		{realValue{3.0}, intValue{3}, boolValue{true}},
		{realValue{3.0}, intValue{-3}, boolValue{true}},
		{realValue{3.0}, intValue{6}, boolValue{false}},
		{realValue{3.0}, realValue{3.0}, boolValue{true}},
		{realValue{3.0}, realValue{-3.0}, boolValue{true}},
		{realValue{3.0}, realValue{6.0}, boolValue{false}},
	}
	for n, testCase := range testCases {
		op := greaterThanOrEqualOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.equal {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.equal)
		}
	}
}
