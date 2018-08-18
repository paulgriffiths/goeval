package vareval

import (
	"fmt"
)

type stmt interface {
	execute(e *env) error
}

type output struct {
	exp expr
}

func (o *output) execute(e *env) error {
	value, err := o.exp.evaluate(e.table)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.output, "%v\n", value)
	return nil
}
