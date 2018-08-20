package vareval

import "testing"

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
		if !testCase.result.equals(testCase.check) {
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
		if !testCase.result.equals(testCase.check) {
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
		if !testCase.result.equals(testCase.check) {
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
		if !quot.equals(testCase.quot) {
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
		{intValue{2}, intValue{3}, realValue{8.0}},
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
		if !prod.equals(testCase.prod) {
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