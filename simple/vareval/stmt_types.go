package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
)

type Stmt interface {
	Execute(e *env) error
}

type outputStmt struct {
	exp expr.Expr
}

func (o *outputStmt) Execute(e *env) error {
	value, err := o.exp.Evaluate(e.table)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.output, "%v\n", value)
	return nil
}

func NewOutputStatement(exp expr.Expr) *outputStmt {
	return &outputStmt{exp}
}

type outputExprStmt struct {
	exp expr.Expr
}

func (o *outputExprStmt) Execute(e *env) error {
	value, err := o.exp.Evaluate(e.table)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.output, "  = %v\n", value)
	return nil
}

func NewOutputExprStatement(exp expr.Expr) *outputExprStmt {
	return &outputExprStmt{exp}
}

type assignStmt struct {
	variable string
	exp      expr.Expr
}

func (a *assignStmt) Execute(e *env) error {
	value, err := a.exp.Evaluate(e.table)
	if err != nil {
		return err
	}

	e.table.Store(a.variable, value)
	return nil
}

func NewAssignStatement(name string, exp expr.Expr) *assignStmt {
	return &assignStmt{name, exp}
}
