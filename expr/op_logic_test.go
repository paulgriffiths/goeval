package expr

import "testing"

func TestBoolEqualityOp(t *testing.T) {
	testCases := []struct {
		left, right, want boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{true}},
		{boolValue{true}, boolValue{false}, boolValue{false}},
		{boolValue{false}, boolValue{true}, boolValue{false}},
		{boolValue{false}, boolValue{false}, boolValue{true}},
	}
	for n, testCase := range testCases {
		op := equalityOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.want {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.want)
		}
	}
}

func TestAndOp(t *testing.T) {
	testCases := []struct {
		left, right, want boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{true}},
		{boolValue{true}, boolValue{false}, boolValue{false}},
		{boolValue{false}, boolValue{true}, boolValue{false}},
		{boolValue{false}, boolValue{false}, boolValue{false}},
	}
	for n, testCase := range testCases {
		op := andOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.want {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.want)
		}
	}
}

func TestOrOp(t *testing.T) {
	testCases := []struct {
		left, right, want boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{true}},
		{boolValue{true}, boolValue{false}, boolValue{true}},
		{boolValue{false}, boolValue{true}, boolValue{true}},
		{boolValue{false}, boolValue{false}, boolValue{false}},
	}
	for n, testCase := range testCases {
		op := orOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.want {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.want)
		}
	}
}

func TestXorOp(t *testing.T) {
	testCases := []struct {
		left, right, want boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{false}},
		{boolValue{true}, boolValue{false}, boolValue{true}},
		{boolValue{false}, boolValue{true}, boolValue{true}},
		{boolValue{false}, boolValue{false}, boolValue{false}},
	}
	for n, testCase := range testCases {
		op := xorOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.want {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.want)
		}
	}
}

func TestNorOp(t *testing.T) {
	testCases := []struct {
		left, right, want boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{false}},
		{boolValue{true}, boolValue{false}, boolValue{false}},
		{boolValue{false}, boolValue{true}, boolValue{false}},
		{boolValue{false}, boolValue{false}, boolValue{true}},
	}
	for n, testCase := range testCases {
		op := norOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.want {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.want)
		}
	}
}

func TestNandOp(t *testing.T) {
	testCases := []struct {
		left, right, want boolValue
	}{
		{boolValue{true}, boolValue{true}, boolValue{false}},
		{boolValue{true}, boolValue{false}, boolValue{true}},
		{boolValue{false}, boolValue{true}, boolValue{true}},
		{boolValue{false}, boolValue{false}, boolValue{true}},
	}
	for n, testCase := range testCases {
		op := nandOp{testCase.left, testCase.right}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.want {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.want)
		}
	}
}

func TestNotOp(t *testing.T) {
	testCases := []struct {
		operand, want boolValue
	}{
		{boolValue{true}, boolValue{false}},
		{boolValue{false}, boolValue{true}},
	}
	for n, testCase := range testCases {
		op := notOp{testCase.operand}
		result, err := op.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate operation: %v", n+1, err)
			continue
		}
		if result != testCase.want {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.want)
		}
	}
}
