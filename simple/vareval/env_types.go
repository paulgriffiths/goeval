package vareval

import (
	"github.com/paulgriffiths/goeval/expr"
	"io"
	"os"
)

type env struct {
	table  *expr.SymTab
	output io.Writer
	input  io.Reader
}

func NewStdEnv() *env {
	return &env{expr.NewTable(), os.Stdout, os.Stdin}
}

func NewEnvWithIO(output io.Writer, input io.Reader) *env {
	return &env{expr.NewTable(), output, input}
}
