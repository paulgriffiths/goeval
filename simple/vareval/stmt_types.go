package vareval

import (
	"fmt"
	"github.com/paulgriffiths/goeval/expr"
)

// Stmt represents a programming language statement.
type Stmt interface {
	Execute(e *Env) error
}

type outputStmt struct {
	exp expr.Expr
}

func (o *outputStmt) Execute(e *Env) error {
	value, err := o.exp.Evaluate(e.table)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.output, "%v\n", value)
	return nil
}

// NewOutputStatement returns a new output statement.
func NewOutputStatement(exp expr.Expr) Stmt {
	return &outputStmt{exp}
}

type outputExprStmt struct {
	exp expr.Expr
}

func (o *outputExprStmt) Execute(e *Env) error {
	value, err := o.exp.Evaluate(e.table)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.output, "  = %v\n", value)
	return nil
}

// NewOutputExprStatement returns a new single-expression output statement.
func NewOutputExprStatement(exp expr.Expr) Stmt {
	return &outputExprStmt{exp}
}

type assignStmt struct {
	variable string
	exp      expr.Expr
}

func (a *assignStmt) Execute(e *Env) error {
	value, err := a.exp.Evaluate(e.table)
	if err != nil {
		return err
	}

	e.table.Store(a.variable, value)
	return nil
}

// NewAssignStatement returns a new assignment statement.
func NewAssignStatement(name string, exp expr.Expr) Stmt {
	return &assignStmt{name, exp}
}
