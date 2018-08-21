package expr

import (
	"math"
	"testing"
)

func TestCosOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{0.0}, realValue{1.0}},
		{realValue{60.0}, realValue{0.5}},
		{realValue{90.0}, realValue{0.0}},
	}
	for n, testCase := range testCases {
		got, err := cosOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestSinOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{0.0}, realValue{0.0}},
		{realValue{30.0}, realValue{0.5}},
		{realValue{90.0}, realValue{1.0}},
	}
	for n, testCase := range testCases {
		got, err := sinOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestTanOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{0.0}, realValue{0.0}},
		{realValue{45.0}, realValue{1.0}},
	}
	for n, testCase := range testCases {
		got, err := tanOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestAcosOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{1.0}, realValue{0.0}},
		{realValue{0.5}, realValue{60.0}},
		{realValue{0.0}, realValue{90.0}},
	}
	for n, testCase := range testCases {
		got, err := acosOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestAsinOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{0.0}, realValue{0.0}},
		{realValue{0.5}, realValue{30.0}},
		{realValue{1.0}, realValue{90.0}},
	}
	for n, testCase := range testCases {
		got, err := asinOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestAtanOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{0.0}, realValue{0.0}},
		{realValue{1.0}, realValue{45.0}},
	}
	for n, testCase := range testCases {
		got, err := atanOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestCeilOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{6.1}, realValue{7.0}},
		{realValue{6.9}, realValue{7.0}},
		{realValue{-6.1}, realValue{-6.0}},
		{realValue{-6.9}, realValue{-6.0}},
	}
	for n, testCase := range testCases {
		got, err := ceilOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).Equals(testCase.want) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestFloorOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{6.1}, realValue{6.0}},
		{realValue{6.9}, realValue{6.0}},
		{realValue{-6.1}, realValue{-7.0}},
		{realValue{-6.9}, realValue{-7.0}},
	}
	for n, testCase := range testCases {
		got, err := floorOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).Equals(testCase.want) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestRoundOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{6.1}, realValue{6.0}},
		{realValue{6.9}, realValue{7.0}},
		{realValue{-6.1}, realValue{-6.0}},
		{realValue{-6.9}, realValue{-7.0}},
	}
	for n, testCase := range testCases {
		got, err := roundOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).Equals(testCase.want) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestLogOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{0.001}, realValue{-3.0}},
		{realValue{0.01}, realValue{-2.0}},
		{realValue{0.1}, realValue{-1.0}},
		{realValue{1.0}, realValue{0.0}},
		{realValue{10.0}, realValue{1.0}},
		{realValue{100.0}, realValue{2.0}},
		{realValue{1000.0}, realValue{3.0}},
	}
	for n, testCase := range testCases {
		got, err := logOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestLnOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{math.Pow(math.E, -3.0)}, realValue{-3.0}},
		{realValue{math.Pow(math.E, -2.0)}, realValue{-2.0}},
		{realValue{math.Pow(math.E, -1.0)}, realValue{-1.0}},
		{realValue{1.0}, realValue{0.0}},
		{realValue{math.E}, realValue{1.0}},
		{realValue{math.Pow(math.E, 2)}, realValue{2.0}},
		{realValue{math.Pow(math.E, 3)}, realValue{3.0}},
	}
	for n, testCase := range testCases {
		got, err := lnOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestSqrtOp(t *testing.T) {
	testCases := []struct {
		operand, want realValue
	}{
		{realValue{0.25}, realValue{0.5}},
		{realValue{16.0}, realValue{4.0}},
		{realValue{144.0}, realValue{12.0}},
	}
	for n, testCase := range testCases {
		got, err := sqrtOp{testCase.operand}.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate function: %v", n+1, err)
			continue
		}
		if !got.(arithmeticValue).almostEquals(testCase.want, 0.0000001) {
			t.Errorf("case %d, got %v, want %v", n+1, got, testCase.want)
		}
	}
}

func TestLogOpBadOperand(t *testing.T) {
	testCases := []struct {
		operand realValue
		err     error
	}{
		{realValue{-1000.0}, DomainError},
	}
	for n, testCase := range testCases {
		_, err := logOp{testCase.operand}.Evaluate(nil)
		if err != testCase.err {
			t.Errorf("case %d, got %v, want %v", n, err, testCase.err)
			continue
		}
	}
}

func TestSqrtOpBadOperand(t *testing.T) {
	testCases := []struct {
		operand realValue
		err     error
	}{
		{realValue{-1.0}, DomainError},
	}
	for n, testCase := range testCases {
		_, err := sqrtOp{testCase.operand}.Evaluate(nil)
		if err != testCase.err {
			t.Errorf("case %d, got %v, want %v", n, err, testCase.err)
			continue
		}
	}
}
