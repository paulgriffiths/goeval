package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
)

type stmt interface {
	execute(e *env) error
}

type output struct {
	exp expr.Expr
}

func (o *output) execute(e *env) error {
	value, err := o.exp.Evaluate(e.table)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.output, "%v\n", value)
	return nil
}
