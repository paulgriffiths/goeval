package expr

import "testing"

func TestIntValueEqual(t *testing.T) {
	testCases := []struct {
		left, right Value
		result      bool
	}{
		{intValue{3}, intValue{3}, true},
		{intValue{3}, intValue{4}, false},
		{intValue{3}, intValue{-3}, false},
		{intValue{3}, intValue{0}, false},
		{intValue{3}, realValue{3}, false},
		{intValue{3}, boolValue{true}, false},
		{intValue{3}, stringValue{"three"}, false},
		{intValue{3}, variableValue{"foobar"}, false},
	}

	for n, testCase := range testCases {
		if testCase.left.Equals(testCase.right) != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, !testCase.result,
				testCase.result)
		}
	}
}

func TestIntValueAlmostEqual(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		result      bool
	}{
		{intValue{3}, intValue{3}, true},
		{intValue{3}, intValue{4}, false},
		{intValue{3}, intValue{-3}, false},
		{intValue{3}, intValue{0}, false},
		{intValue{3}, realValue{3}, false},
	}

	for n, testCase := range testCases {
		result := testCase.left.almostEquals(testCase.right, 0)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}

func TestIntValueAdd(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{intValue{3}.add(intValue{7}), intValue{10}},
		{intValue{3}.add(intValue{-7}), intValue{-4}},
		{intValue{-3}.add(intValue{7}), intValue{4}},
		{intValue{-3}.add(intValue{-7}), intValue{-10}},
		{intValue{3}.add(realValue{7.0}), realValue{10.0}},
		{intValue{3}.add(realValue{-7.0}), realValue{-4.0}},
		{intValue{-3}.add(realValue{7.0}), realValue{4.0}},
		{intValue{-3}.add(realValue{-7.0}), realValue{-10.0}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}

func TestIntValueSub(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{intValue{3}.sub(intValue{7}), intValue{-4}},
		{intValue{3}.sub(intValue{-7}), intValue{10}},
		{intValue{-3}.sub(intValue{7}), intValue{-10}},
		{intValue{-3}.sub(intValue{-7}), intValue{4}},
		{intValue{3}.sub(realValue{7.0}), realValue{-4.0}},
		{intValue{3}.sub(realValue{-7.0}), realValue{10.0}},
		{intValue{-3}.sub(realValue{7.0}), realValue{-10.0}},
		{intValue{-3}.sub(realValue{-7.0}), realValue{4.0}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}

func TestIntValueMul(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{intValue{3}.mul(intValue{7}), intValue{21}},
		{intValue{3}.mul(intValue{-7}), intValue{-21}},
		{intValue{-3}.mul(intValue{7}), intValue{-21}},
		{intValue{-3}.mul(intValue{-7}), intValue{21}},
		{intValue{3}.mul(realValue{7.0}), realValue{21.0}},
		{intValue{3}.mul(realValue{-7.0}), realValue{-21.0}},
		{intValue{-3}.mul(realValue{7.0}), realValue{-21.0}},
		{intValue{-3}.mul(realValue{-7.0}), realValue{21.0}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}

func TestIntValueDiv(t *testing.T) {
	testCases := []struct {
		dend, dvsr, quot arithmeticValue
	}{
		{intValue{12}, intValue{4}, intValue{3}},
		{intValue{12}, intValue{-4}, intValue{-3}},
		{intValue{-12}, intValue{4}, intValue{-3}},
		{intValue{-12}, intValue{-4}, intValue{3}},
		{intValue{12}, intValue{8}, realValue{1.5}},
		{intValue{12}, intValue{-8}, realValue{-1.5}},
		{intValue{-12}, intValue{8}, realValue{-1.5}},
		{intValue{-12}, intValue{-8}, realValue{1.5}},
		{intValue{12}, realValue{4.0}, realValue{3.0}},
		{intValue{12}, realValue{-4.0}, realValue{-3.0}},
		{intValue{-12}, realValue{4.0}, realValue{-3.0}},
		{intValue{-12}, realValue{-4.0}, realValue{3.0}},
	}

	for n, testCase := range testCases {
		quot, err := testCase.dend.div(testCase.dvsr)
		if err != nil {
			t.Errorf("case %d, couldn't calculate quotient: %v", n+1, err)
			continue
		}
		if !quot.Equals(testCase.quot) {
			t.Errorf("case %d, got %v, want %v", n+1, quot, testCase.quot)
		}
	}
}

func TestIntValueDivByZero(t *testing.T) {
	testCases := []struct {
		dend, dvsr arithmeticValue
	}{
		{intValue{5}, intValue{0}},
		{intValue{5}, realValue{0.0}},
		{intValue{5}, realValue{-0.0}},
	}

	for n, testCase := range testCases {
		_, err := testCase.dend.div(testCase.dvsr)
		if err != DivideByZeroError {
			t.Errorf("case %d, got %v, want: %v", n+1,
				err, DivideByZeroError)
		}
	}
}

func TestIntValuePow(t *testing.T) {
	testCases := []struct {
		base, exp, prod arithmeticValue
	}{
		{intValue{2}, intValue{0}, intValue{1}},
		{intValue{2}, intValue{1}, intValue{2}},
		{intValue{2}, intValue{3}, intValue{8}},
		{intValue{2}, realValue{3.0}, realValue{8.0}},
		{intValue{16}, realValue{1.0}, realValue{16.0}},
		{intValue{16}, realValue{0.5}, realValue{4.0}},
		{intValue{16}, realValue{0.25}, realValue{2.0}},
		{intValue{16}, realValue{0.0}, realValue{1.0}},
	}

	for n, testCase := range testCases {
		prod, err := testCase.base.pow(testCase.exp)
		if err != nil {
			t.Errorf("case %d, couldn't calculate product: %v", n+1, err)
			continue
		}
		if !prod.Equals(testCase.prod) {
			t.Errorf("case %d, got %v, want %v", n+1, prod, testCase.prod)
		}
	}
}

func TestIntValuePowBadDomain(t *testing.T) {
	testCases := []struct {
		base, exp arithmeticValue
	}{
		{intValue{-1}, realValue{0.5}},
		{intValue{-1}, realValue{1.5}},
	}

	for n, testCase := range testCases {
		_, err := testCase.base.pow(testCase.exp)
		if err != DomainError {
			t.Errorf("case %d, got %v, want: %v", n+1,
				err, DivideByZeroError)
		}
	}
}

func TestIntValueNegate(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{intValue{3}.negate(), intValue{-3}},
		{intValue{-3}.negate(), intValue{3}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}

func TestIntValueEquality(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		result      bool
	}{
		{intValue{3}, intValue{3}, true},
		{intValue{3}, intValue{-3}, false},
		{intValue{3}, intValue{2}, false},
		{intValue{3}, intValue{-2}, false},
	}

	for n, testCase := range testCases {
		result := testCase.left.equality(testCase.right)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}
