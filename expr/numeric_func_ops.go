package expr

import (
	"math"
)

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func toDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

type cosOp struct {
	value Expr
}

func (op cosOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Cos(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewCos(value Expr) cosOp {
	return cosOp{value}
}

type sinOp struct {
	value Expr
}

func (op sinOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Sin(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewSin(value Expr) sinOp {
	return sinOp{value}
}

type tanOp struct {
	value Expr
}

func (op tanOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Tan(toRadians(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewTan(value Expr) tanOp {
	return tanOp{value}
}

type acosOp struct {
	value Expr
}

func (op acosOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Acos(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewAcos(value Expr) acosOp {
	return acosOp{value}
}

type asinOp struct {
	value Expr
}

func (op asinOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Asin(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewAsin(value Expr) asinOp {
	return asinOp{value}
}

type atanOp struct {
	value Expr
}

func (op atanOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := toDegrees(math.Atan(exps[0].(arithmeticValue).floatValue()))
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewAtan(value Expr) atanOp {
	return atanOp{value}
}

type roundOp struct {
	value Expr
}

func (op roundOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Round(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

func NewRound(value Expr) roundOp {
	return roundOp{value}
}

type ceilOp struct {
	value Expr
}

func (op ceilOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Ceil(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

func NewCeil(value Expr) ceilOp {
	return ceilOp{value}
}

type floorOp struct {
	value Expr
}

func (op floorOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Floor(exps[0].(arithmeticValue).floatValue())
	return realValue{result}, nil
}

func NewFloor(value Expr) floorOp {
	return floorOp{value}
}

type sqrtOp struct {
	value Expr
}

func (op sqrtOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Sqrt(exps[0].(arithmeticValue).floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewSqrt(value Expr) sqrtOp {
	return sqrtOp{value}
}

type logOp struct {
	value Expr
}

func (op logOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Log10(exps[0].(arithmeticValue).floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewLog(value Expr) logOp {
	return logOp{value}
}

type lnOp struct {
	value Expr
}

func (op lnOp) Evaluate(table *symTab) (Expr, error) {
	exps, err := evaluateExprs(table, IsNumeric, op.value)
	if err != nil {
		return nil, err
	}
	result := math.Log(exps[0].(arithmeticValue).floatValue())
	if math.IsNaN(result) {
		return nil, DomainError
	}
	return realValue{result}, nil
}

func NewLn(value Expr) lnOp {
	return lnOp{value}
}
