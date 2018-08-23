package expr_test

import (
	"github.com/paulgriffiths/goeval/expr"
	"testing"
)

func TestEvaluateVariablesSuccess(t *testing.T) {
	table := expr.NewTable()
	table.Store("n", expr.NewInt(1))
	table.Store("r", expr.NewReal(1.0))
	table.Store("z", expr.NewReal(0.0))
	table.Store("b", expr.NewBool(true))
	table.Store("s", expr.NewString("citadel"))

	n := expr.NewInt(1)
	r := expr.NewReal(1.0)
	b := expr.NewBool(true)
	s := expr.NewString("citadel")
	vn := expr.NewVariable("n")
	vr := expr.NewVariable("r")
	vz := expr.NewVariable("z")
	vb := expr.NewVariable("b")
	vs := expr.NewVariable("s")

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAdd(n, vn), expr.NewInt(2)},
		{expr.NewAdd(vn, n), expr.NewInt(2)},
		{expr.NewAdd(vn, vn), expr.NewInt(2)},
		{expr.NewAdd(r, vr), expr.NewReal(2.0)},
		{expr.NewAdd(vr, r), expr.NewReal(2.0)},
		{expr.NewAdd(vr, vr), expr.NewReal(2.0)},
		{expr.NewSub(n, vn), expr.NewInt(0)},
		{expr.NewSub(vn, n), expr.NewInt(0)},
		{expr.NewSub(vn, vn), expr.NewInt(0)},
		{expr.NewSub(r, vr), expr.NewReal(0.0)},
		{expr.NewSub(vr, r), expr.NewReal(0.0)},
		{expr.NewSub(vr, vr), expr.NewReal(0.0)},
		{expr.NewMul(n, vn), expr.NewInt(1)},
		{expr.NewMul(vn, n), expr.NewInt(1)},
		{expr.NewMul(vn, vn), expr.NewInt(1)},
		{expr.NewMul(r, vr), expr.NewReal(1.0)},
		{expr.NewMul(vr, r), expr.NewReal(1.0)},
		{expr.NewMul(vr, vr), expr.NewReal(1.0)},
		{expr.NewDiv(n, vn), expr.NewInt(1)},
		{expr.NewDiv(vn, n), expr.NewInt(1)},
		{expr.NewDiv(vn, vn), expr.NewInt(1)},
		{expr.NewDiv(r, vr), expr.NewReal(1.0)},
		{expr.NewDiv(vr, r), expr.NewReal(1.0)},
		{expr.NewDiv(vr, vr), expr.NewReal(1.0)},
		{expr.NewPow(n, vn), expr.NewInt(1)},
		{expr.NewPow(vn, n), expr.NewInt(1)},
		{expr.NewPow(vn, vn), expr.NewInt(1)},
		{expr.NewPow(r, vr), expr.NewReal(1.0)},
		{expr.NewPow(vr, r), expr.NewReal(1.0)},
		{expr.NewPow(vr, vr), expr.NewReal(1.0)},
		{expr.NewNeg(vn), expr.NewInt(-1)},
		{expr.NewNeg(vr), expr.NewReal(-1.0)},
		{expr.NewCos(vz), expr.NewReal(1.0)},
		{expr.NewSin(vz), expr.NewReal(0.0)},
		{expr.NewTan(vz), expr.NewReal(0.0)},
		{expr.NewAcos(vr), expr.NewReal(0.0)},
		{expr.NewAsin(vr), expr.NewReal(90.0)},
		{expr.NewAtan(vr), expr.NewReal(45.0)},
		{expr.NewCeil(vr), expr.NewReal(1.0)},
		{expr.NewFloor(vr), expr.NewReal(1.0)},
		{expr.NewSqrt(vr), expr.NewReal(1.0)},
		{expr.NewRound(vr), expr.NewReal(1.0)},
		{expr.NewLog(vr), expr.NewReal(0.0)},
		{expr.NewLn(vr), expr.NewReal(0.0)},
		{expr.NewEquality(n, vn), expr.NewBool(true)},
		{expr.NewEquality(vn, n), expr.NewBool(true)},
		{expr.NewEquality(vn, vn), expr.NewBool(true)},
		{expr.NewEquality(r, vr), expr.NewBool(true)},
		{expr.NewEquality(vr, r), expr.NewBool(true)},
		{expr.NewEquality(vr, vr), expr.NewBool(true)},
		{expr.NewNonEquality(n, vn), expr.NewBool(false)},
		{expr.NewNonEquality(vn, n), expr.NewBool(false)},
		{expr.NewNonEquality(vn, vn), expr.NewBool(false)},
		{expr.NewNonEquality(r, vr), expr.NewBool(false)},
		{expr.NewNonEquality(vr, r), expr.NewBool(false)},
		{expr.NewNonEquality(vr, vr), expr.NewBool(false)},
		{expr.NewLessThan(n, vn), expr.NewBool(false)},
		{expr.NewLessThan(vn, n), expr.NewBool(false)},
		{expr.NewLessThan(vn, vn), expr.NewBool(false)},
		{expr.NewLessThan(r, vr), expr.NewBool(false)},
		{expr.NewLessThan(vr, r), expr.NewBool(false)},
		{expr.NewLessThan(vr, vr), expr.NewBool(false)},
		{expr.NewGreaterThan(n, vn), expr.NewBool(false)},
		{expr.NewGreaterThan(vn, n), expr.NewBool(false)},
		{expr.NewGreaterThan(vn, vn), expr.NewBool(false)},
		{expr.NewGreaterThan(r, vr), expr.NewBool(false)},
		{expr.NewGreaterThan(vr, r), expr.NewBool(false)},
		{expr.NewGreaterThan(vr, vr), expr.NewBool(false)},
		{expr.NewGreaterThanOrEqual(n, vn), expr.NewBool(true)},
		{expr.NewGreaterThanOrEqual(vn, n), expr.NewBool(true)},
		{expr.NewGreaterThanOrEqual(vn, vn), expr.NewBool(true)},
		{expr.NewGreaterThanOrEqual(r, vr), expr.NewBool(true)},
		{expr.NewGreaterThanOrEqual(vr, r), expr.NewBool(true)},
		{expr.NewGreaterThanOrEqual(vr, vr), expr.NewBool(true)},
		{expr.NewLessThanOrEqual(n, vn), expr.NewBool(true)},
		{expr.NewLessThanOrEqual(vn, n), expr.NewBool(true)},
		{expr.NewLessThanOrEqual(vn, vn), expr.NewBool(true)},
		{expr.NewLessThanOrEqual(r, vr), expr.NewBool(true)},
		{expr.NewLessThanOrEqual(vr, r), expr.NewBool(true)},
		{expr.NewLessThanOrEqual(vr, vr), expr.NewBool(true)},
		{expr.NewAnd(b, vb), expr.NewBool(true)},
		{expr.NewAnd(vb, b), expr.NewBool(true)},
		{expr.NewAnd(vb, vb), expr.NewBool(true)},
		{expr.NewOr(b, vb), expr.NewBool(true)},
		{expr.NewOr(vb, b), expr.NewBool(true)},
		{expr.NewOr(vb, vb), expr.NewBool(true)},
		{expr.NewXor(b, vb), expr.NewBool(false)},
		{expr.NewXor(vb, b), expr.NewBool(false)},
		{expr.NewXor(vb, vb), expr.NewBool(false)},
		{expr.NewNor(b, vb), expr.NewBool(false)},
		{expr.NewNor(vb, b), expr.NewBool(false)},
		{expr.NewNor(vb, vb), expr.NewBool(false)},
		{expr.NewNand(b, vb), expr.NewBool(false)},
		{expr.NewNand(vb, b), expr.NewBool(false)},
		{expr.NewNand(vb, vb), expr.NewBool(false)},
		{expr.NewNot(vb), expr.NewBool(false)},
		{expr.NewConcat(s, vs), expr.NewString("citadelcitadel")},
		{expr.NewConcat(vs, s), expr.NewString("citadelcitadel")},
		{expr.NewConcat(vs, vs), expr.NewString("citadelcitadel")},
	}

	for i, testCase := range testCases {
		result, err := testCase.exp.Evaluate(table)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", i+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, table) {
			t.Errorf("case %d, got %v, want %v", i+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateArithmeticNoSymbolTableSuccess(t *testing.T) {
	i := expr.NewInt(3)
	j := expr.NewInt(6)
	x := expr.NewReal(16.0)
	y := expr.NewReal(4.0)

	testCases := []struct {
		exp    expr.Expr
		result expr.Expr
	}{
		{expr.NewAdd(i, j), expr.NewInt(9)},
		{expr.NewAdd(x, y), expr.NewReal(20.0)},
		{expr.NewSub(i, j), expr.NewInt(-3)},
		{expr.NewSub(x, y), expr.NewReal(12.0)},
		{expr.NewMul(i, j), expr.NewInt(18)},
		{expr.NewMul(x, y), expr.NewReal(64.0)},
		{expr.NewDiv(i, j), expr.NewReal(0.5)},
		{expr.NewDiv(x, y), expr.NewReal(4.0)},
		{expr.NewPow(i, j), expr.NewInt(729)},
		{expr.NewPow(x, y), expr.NewReal(65536.0)},
		{expr.NewNeg(i), expr.NewInt(-3)},
		{expr.NewNeg(x), expr.NewReal(-16.0)},
	}

	for n, testCase := range testCases {
		result, err := testCase.exp.Evaluate(nil)
		if err != nil {
			t.Errorf("case %d, couldn't evaluate expression: %v", n+1, err)
			continue
		}
		if !expr.Equals(result, testCase.result, nil) {
			t.Errorf("case %d, got %v, want %v", n+1,
				result, testCase.result)
		}
	}
}

func TestEvaluateUndefinedVariable(t *testing.T) {
	n := expr.NewInt(1)
	r := expr.NewReal(1.0)
	b := expr.NewBool(true)
	s := expr.NewString("citadel")
	vn := expr.NewVariable("n")
	vr := expr.NewVariable("r")
	vz := expr.NewVariable("z")
	vb := expr.NewVariable("b")
	vs := expr.NewVariable("s")

	testCases := []struct {
		exp expr.Expr
	}{
		{expr.NewAdd(n, vn)},
		{expr.NewAdd(vn, n)},
		{expr.NewAdd(vn, vn)},
		{expr.NewAdd(r, vr)},
		{expr.NewAdd(vr, r)},
		{expr.NewAdd(vr, vr)},
		{expr.NewSub(n, vn)},
		{expr.NewSub(vn, n)},
		{expr.NewSub(vn, vn)},
		{expr.NewSub(r, vr)},
		{expr.NewSub(vr, r)},
		{expr.NewSub(vr, vr)},
		{expr.NewMul(n, vn)},
		{expr.NewMul(vn, n)},
		{expr.NewMul(vn, vn)},
		{expr.NewMul(r, vr)},
		{expr.NewMul(vr, r)},
		{expr.NewMul(vr, vr)},
		{expr.NewDiv(n, vn)},
		{expr.NewDiv(vn, n)},
		{expr.NewDiv(vn, vn)},
		{expr.NewDiv(r, vr)},
		{expr.NewDiv(vr, r)},
		{expr.NewDiv(vr, vr)},
		{expr.NewPow(n, vn)},
		{expr.NewPow(vn, n)},
		{expr.NewPow(vn, vn)},
		{expr.NewPow(r, vr)},
		{expr.NewPow(vr, r)},
		{expr.NewPow(vr, vr)},
		{expr.NewNeg(vn)},
		{expr.NewNeg(vr)},
		{expr.NewCos(vz)},
		{expr.NewSin(vz)},
		{expr.NewTan(vz)},
		{expr.NewAcos(vr)},
		{expr.NewAsin(vr)},
		{expr.NewAtan(vr)},
		{expr.NewCeil(vr)},
		{expr.NewFloor(vr)},
		{expr.NewSqrt(vr)},
		{expr.NewRound(vr)},
		{expr.NewLog(vr)},
		{expr.NewLn(vr)},
		{expr.NewEquality(n, vn)},
		{expr.NewEquality(vn, n)},
		{expr.NewEquality(vn, vn)},
		{expr.NewEquality(r, vr)},
		{expr.NewEquality(vr, r)},
		{expr.NewEquality(vr, vr)},
		{expr.NewNonEquality(n, vn)},
		{expr.NewNonEquality(vn, n)},
		{expr.NewNonEquality(vn, vn)},
		{expr.NewNonEquality(r, vr)},
		{expr.NewNonEquality(vr, r)},
		{expr.NewNonEquality(vr, vr)},
		{expr.NewLessThan(n, vn)},
		{expr.NewLessThan(vn, n)},
		{expr.NewLessThan(vn, vn)},
		{expr.NewLessThan(r, vr)},
		{expr.NewLessThan(vr, r)},
		{expr.NewLessThan(vr, vr)},
		{expr.NewGreaterThan(n, vn)},
		{expr.NewGreaterThan(vn, n)},
		{expr.NewGreaterThan(vn, vn)},
		{expr.NewGreaterThan(r, vr)},
		{expr.NewGreaterThan(vr, r)},
		{expr.NewGreaterThan(vr, vr)},
		{expr.NewGreaterThanOrEqual(n, vn)},
		{expr.NewGreaterThanOrEqual(vn, n)},
		{expr.NewGreaterThanOrEqual(vn, vn)},
		{expr.NewGreaterThanOrEqual(r, vr)},
		{expr.NewGreaterThanOrEqual(vr, r)},
		{expr.NewGreaterThanOrEqual(vr, vr)},
		{expr.NewLessThanOrEqual(n, vn)},
		{expr.NewLessThanOrEqual(vn, n)},
		{expr.NewLessThanOrEqual(vn, vn)},
		{expr.NewLessThanOrEqual(r, vr)},
		{expr.NewLessThanOrEqual(vr, r)},
		{expr.NewLessThanOrEqual(vr, vr)},
		{expr.NewAnd(b, vb)},
		{expr.NewAnd(vb, b)},
		{expr.NewAnd(vb, vb)},
		{expr.NewOr(b, vb)},
		{expr.NewOr(vb, b)},
		{expr.NewOr(vb, vb)},
		{expr.NewXor(b, vb)},
		{expr.NewXor(vb, b)},
		{expr.NewXor(vb, vb)},
		{expr.NewNor(b, vb)},
		{expr.NewNor(vb, b)},
		{expr.NewNor(vb, vb)},
		{expr.NewNand(b, vb)},
		{expr.NewNand(vb, b)},
		{expr.NewNand(vb, vb)},
		{expr.NewNot(vb)},
		{expr.NewConcat(s, vs)},
		{expr.NewConcat(vs, s)},
		{expr.NewConcat(vs, vs)},
	}

	for n, testCase := range testCases {
		_, err := testCase.exp.Evaluate(nil)
		if err != expr.UndefinedVariableError {
			t.Errorf("case %d, got %v, want %v", n+1, err,
				expr.UndefinedVariableError)
			continue
		}
	}
}
