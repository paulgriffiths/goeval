package expr

import "testing"

func TestEqualityMismatchedTypes(t *testing.T) {
	testCases := []struct {
		left, right value
	}{
		{boolValue{true}, intValue{4}},
		{boolValue{true}, realValue{4.0}},
		{boolValue{false}, stringValue{"not a bool"}},
		{intValue{4}, boolValue{false}},
		{intValue{4}, stringValue{"not an int"}},
		{realValue{4.0}, boolValue{false}},
		{realValue{4.0}, stringValue{"not a float"}},
		{stringValue{"not an int"}, intValue{4}},
		{stringValue{"not a float"}, realValue{4.0}},
		{stringValue{"not a bool"}, boolValue{false}},
	}
	for n, testCase := range testCases {
		op := equalityOp{testCase.left, testCase.right}
		_, err := op.Evaluate(nil)
		if err != TypeError {
			t.Errorf("case %d, got %v, want %v", n, err, TypeError)
		}
	}
}
