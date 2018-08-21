package expr

import "testing"

func TestRealValueEqual(t *testing.T) {
	testCases := []struct {
		left, right Value
		result      bool
	}{
		{realValue{3.0}, realValue{3.0}, true},
		{realValue{3.0}, realValue{4.0}, false},
		{realValue{3.0}, realValue{-3.0}, false},
		{realValue{3.0}, realValue{0.0}, false},
		{realValue{3.0}, intValue{3}, false},
		{realValue{3.0}, boolValue{true}, false},
		{realValue{3.0}, stringValue{"three"}, false},
		{realValue{3.0}, variableValue{"foobar"}, false},
	}

	for n, testCase := range testCases {
		if testCase.left.Equals(testCase.right) != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, !testCase.result,
				testCase.result)
		}
	}
}

func TestRealValueAlmostEqual(t *testing.T) {
	testCases := []struct {
		left, right arithmeticValue
		eps         float64
		result      bool
	}{
		{realValue{1.0}, realValue{1.0}, 0.0, true},
		{realValue{1.0}, realValue{1.9}, 1.0, true},
		{realValue{1.0}, realValue{1.9}, 0.0, false},
		{realValue{1.01}, realValue{1.09}, 0.1, true},
		{realValue{1.01}, realValue{1.09}, 0.01, false},
		{realValue{1.01}, realValue{1.09}, 0.0, false},
		{realValue{1.001}, realValue{1.009}, 0.1, true},
		{realValue{1.001}, realValue{1.009}, 0.01, true},
		{realValue{1.001}, realValue{1.009}, 0.001, false},
		{realValue{1.001}, realValue{1.009}, 0.0, false},
		{realValue{1.0001}, realValue{1.0009}, 0.1, true},
		{realValue{1.0001}, realValue{1.0009}, 0.01, true},
		{realValue{1.0001}, realValue{1.0009}, 0.001, true},
		{realValue{1.0001}, realValue{1.0009}, 0.0001, false},
		{realValue{1.0001}, realValue{1.0009}, 0.0, false},
	}

	for n, testCase := range testCases {
		result := testCase.left.almostEquals(testCase.right, testCase.eps)
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1, result,
				testCase.result)
		}
	}
}

func TestRealValueAdd(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{realValue{3.0}.add(intValue{7}), realValue{10.0}},
		{realValue{3.0}.add(intValue{-7}), realValue{-4.0}},
		{realValue{-3.0}.add(intValue{7}), realValue{4.0}},
		{realValue{-3.0}.add(intValue{-7}), realValue{-10.0}},
		{realValue{3.0}.add(realValue{7.0}), realValue{10.0}},
		{realValue{3.0}.add(realValue{-7.0}), realValue{-4.0}},
		{realValue{-3.0}.add(realValue{7.0}), realValue{4.0}},
		{realValue{-3.0}.add(realValue{-7.0}), realValue{-10.0}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}

func TestRealValueSub(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{realValue{3.0}.sub(intValue{7}), realValue{-4.0}},
		{realValue{3.0}.sub(intValue{-7}), realValue{10.0}},
		{realValue{-3.0}.sub(intValue{7}), realValue{-10.0}},
		{realValue{-3.0}.sub(intValue{-7}), realValue{4.0}},
		{realValue{3.0}.sub(realValue{7.0}), realValue{-4.0}},
		{realValue{3.0}.sub(realValue{-7.0}), realValue{10.0}},
		{realValue{-3.0}.sub(realValue{7.0}), realValue{-10.0}},
		{realValue{-3.0}.sub(realValue{-7.0}), realValue{4.0}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}

func TestRealValueMul(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{realValue{3.0}.mul(intValue{7}), realValue{21.0}},
		{realValue{3.0}.mul(intValue{-7}), realValue{-21.0}},
		{realValue{-3.0}.mul(intValue{7}), realValue{-21.0}},
		{realValue{-3.0}.mul(intValue{-7}), realValue{21.0}},
		{realValue{3.0}.mul(realValue{7.0}), realValue{21.0}},
		{realValue{3.0}.mul(realValue{-7.0}), realValue{-21.0}},
		{realValue{-3.0}.mul(realValue{7.0}), realValue{-21.0}},
		{realValue{-3.0}.mul(realValue{-7.0}), realValue{21.0}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}

func TestRealValueDiv(t *testing.T) {
	testCases := []struct {
		dend, dvsr, quot arithmeticValue
	}{
		{realValue{12.0}, intValue{8}, realValue{1.5}},
		{realValue{12.0}, intValue{-8}, realValue{-1.5}},
		{realValue{-12.0}, intValue{8}, realValue{-1.5}},
		{realValue{-12.0}, intValue{-8}, realValue{1.5}},
		{realValue{12.0}, realValue{4.0}, realValue{3.0}},
		{realValue{12.0}, realValue{-4.0}, realValue{-3.0}},
		{realValue{-12.0}, realValue{4.0}, realValue{-3.0}},
		{realValue{-12.0}, realValue{-4.0}, realValue{3.0}},
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

func TestRealValueDivByZero(t *testing.T) {
	testCases := []struct {
		dend, dvsr arithmeticValue
	}{
		{realValue{5.0}, intValue{0}},
		{realValue{5.0}, realValue{0.0}},
		{realValue{5.0}, realValue{-0.0}},
	}

	for n, testCase := range testCases {
		_, err := testCase.dend.div(testCase.dvsr)
		if err != DivideByZeroError {
			t.Errorf("case %d, got %v, want: %v", n+1,
				err, DivideByZeroError)
		}
	}
}

func TestRealValuePow(t *testing.T) {
	testCases := []struct {
		base, exp, prod arithmeticValue
	}{
		{realValue{2.0}, intValue{3}, realValue{8.0}},
		{realValue{2.0}, realValue{3.0}, realValue{8.0}},
		{realValue{16.0}, realValue{1.0}, realValue{16.0}},
		{realValue{16.0}, realValue{0.5}, realValue{4.0}},
		{realValue{16.0}, realValue{0.25}, realValue{2.0}},
		{realValue{16.0}, realValue{0.0}, realValue{1.0}},
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

func TestRealValuePowBadDomain(t *testing.T) {
	testCases := []struct {
		base, exp arithmeticValue
	}{
		{realValue{-1}, realValue{0.5}},
		{realValue{-1}, realValue{1.5}},
	}

	for n, testCase := range testCases {
		_, err := testCase.base.pow(testCase.exp)
		if err != DomainError {
			t.Errorf("case %d, got %v, want: %v", n+1,
				err, DivideByZeroError)
		}
	}
}

func TestRealValueNegate(t *testing.T) {
	testCases := []struct {
		result, check arithmeticValue
	}{
		{realValue{3.0}.negate(), realValue{-3.0}},
		{realValue{-3.0}.negate(), realValue{3.0}},
	}

	for n, testCase := range testCases {
		if !testCase.result.Equals(testCase.check) {
			t.Errorf("case %d, got %v, want %v", n+1, testCase.result,
				testCase.check)
		}
	}
}
