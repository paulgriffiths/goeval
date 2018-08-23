package expr_test

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
)

func Example() {

	// Addition of two integers, with integer result

	add := expr.NewAdd(expr.NewInt(5), expr.NewInt(8))
	sum, _ := add.Evaluate(nil)
	fmt.Printf("%v = %v\n", add, sum)

	// Multiplication by a real number, with real result

	mul := expr.NewMul(add, expr.NewReal(1.5))
	product, _ := mul.Evaluate(nil)
	fmt.Printf("%v = %v\n", mul, product)

	// Comparison with an integer stored in a variable, with boolean result

	table := expr.NewTable()
	table.Store("m", expr.NewInt(20))
	cmp := expr.NewLessThan(mul, expr.NewVariable("m"))
	result, _ := cmp.Evaluate(table)
	fmt.Printf("%v = %v, when m = 20\n", cmp, result)

	// Division by zero, after redefining variable m, with error

	table.Store("m", expr.NewInt(0))
	_, err := expr.NewDiv(mul, expr.NewVariable("m")).Evaluate(table)
	if err != nil {
		fmt.Printf("couldn't evaluate div expression: %v\n", err)
	}

	// Output:
	// (5)+(8) = 13
	// ((5)+(8))*(1.5) = 19.5
	// (((5)+(8))*(1.5))<(m) = true, when m = 20
	// couldn't evaluate div expression: divide by zero
}
