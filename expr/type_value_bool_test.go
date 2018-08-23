package expr

import "testing"

func TestBoolValueEqual(t *testing.T) {
	testCases := []struct {
		left, right value
		result      bool
	}{
		{boolValue{true}, boolValue{true}, true},
		{boolValue{true}, boolValue{false}, false},
		{boolValue{false}, boolValue{true}, false},
		{boolValue{false}, boolValue{false}, true},
	}

	for n, testCase := range testCases {
		if testCase.left.Equals(testCase.right) != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, !testCase.result,
				testCase.result)
		}
	}
}

func TestBoolValueEquality(t *testing.T) {
	testCases := []struct {
		left, right boolValue
        result bool
	}{
		{boolValue{true}, boolValue{true}, true},
		{boolValue{true}, boolValue{false}, false},
		{boolValue{false}, boolValue{true}, false},
		{boolValue{false}, boolValue{false}, true},
	}

	for n, testCase := range testCases {
		result := testCase.left.equality(testCase.right)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}

func TestBoolAnd(t *testing.T) {
	testCases := []struct {
		left, right, result boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{true}},
		{boolValue{true}, boolValue{false}, boolValue{false}},
		{boolValue{false}, boolValue{true}, boolValue{false}},
		{boolValue{false}, boolValue{false}, boolValue{false}},
	}

	for n, testCase := range testCases {
		result := testCase.left.and(testCase.right)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}

func TestBoolOr(t *testing.T) {
	testCases := []struct {
		left, right, result boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{true}},
		{boolValue{true}, boolValue{false}, boolValue{true}},
		{boolValue{false}, boolValue{true}, boolValue{true}},
		{boolValue{false}, boolValue{false}, boolValue{false}},
	}

	for n, testCase := range testCases {
		result := testCase.left.or(testCase.right)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}

func TestBoolXor(t *testing.T) {
	testCases := []struct {
		left, right, result boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{false}},
		{boolValue{true}, boolValue{false}, boolValue{true}},
		{boolValue{false}, boolValue{true}, boolValue{true}},
		{boolValue{false}, boolValue{false}, boolValue{false}},
	}

	for n, testCase := range testCases {
		result := testCase.left.xor(testCase.right)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}

func TestBoolNor(t *testing.T) {
	testCases := []struct {
		left, right, result boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{false}},
		{boolValue{true}, boolValue{false}, boolValue{false}},
		{boolValue{false}, boolValue{true}, boolValue{false}},
		{boolValue{false}, boolValue{false}, boolValue{true}},
	}

	for n, testCase := range testCases {
		result := testCase.left.nor(testCase.right)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}
