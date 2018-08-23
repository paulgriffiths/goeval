package expr

import "testing"

func TestExprStringValues(t *testing.T) {
	testCases := []struct {
		exp    Expr
		result string
	}{
		{NewInt(3), "3"},
		{NewReal(2.5), "2.5"},
		{NewBool(true), "true"},
		{NewString("castle"), "\"castle\""},
		{NewVariable("foobar"), "foobar"},
		{NewAdd(NewInt(3), NewInt(4)), "(3)+(4)"},
		{NewSub(NewInt(3), NewInt(4)), "(3)-(4)"},
		{NewMul(NewInt(3), NewInt(4)), "(3)*(4)"},
		{NewDiv(NewInt(3), NewInt(4)), "(3)/(4)"},
		{NewPow(NewInt(3), NewInt(4)), "(3)^(4)"},
		{NewNeg(NewInt(3)), "-(3)"},
		{NewMul(NewAdd(NewInt(3), NewInt(4)), NewInt(5)), "((3)+(4))*(5)"},
		{NewEquality(NewInt(3), NewInt(4)), "(3)==(4)"},
		{NewNonEquality(NewInt(3), NewInt(4)), "(3)!=(4)"},
		{NewLessThan(NewInt(3), NewInt(4)), "(3)<(4)"},
		{NewLessThanOrEqual(NewInt(3), NewInt(4)), "(3)<=(4)"},
		{NewGreaterThan(NewInt(3), NewInt(4)), "(3)>(4)"},
		{NewGreaterThanOrEqual(NewInt(3), NewInt(4)), "(3)>=(4)"},
		{NewAnd(NewBool(true), NewBool(false)), "and(true,false)"},
		{NewOr(NewBool(true), NewBool(false)), "or(true,false)"},
		{NewXor(NewBool(true), NewBool(false)), "xor(true,false)"},
		{NewNand(NewBool(true), NewBool(false)), "nand(true,false)"},
		{NewNor(NewBool(true), NewBool(false)), "nor(true,false)"},
		{NewNot(NewBool(true)), "not(true)"},
		{NewCos(NewInt(3)), "cos(3)"},
		{NewSin(NewInt(3)), "sin(3)"},
		{NewTan(NewInt(3)), "tan(3)"},
		{NewAcos(NewInt(3)), "arccos(3)"},
		{NewAsin(NewInt(3)), "arcsin(3)"},
		{NewAtan(NewInt(3)), "arctan(3)"},
		{NewCeil(NewInt(3)), "ceil(3)"},
		{NewFloor(NewInt(3)), "floor(3)"},
		{NewRound(NewInt(3)), "round(3)"},
		{NewSqrt(NewInt(3)), "sqrt(3)"},
		{NewLog(NewInt(3)), "log(3)"},
		{NewLn(NewInt(3)), "ln(3)"},
	}

	for n, testCase := range testCases {
		if testCase.exp.String() != testCase.result {
			t.Errorf("case %d, got %v, want %v", n+1,
				testCase.exp.String(), testCase.result)
		}
	}
}
