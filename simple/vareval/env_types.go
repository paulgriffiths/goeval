package vareval

import (
	"github.com/paulgriffiths/goeval/expr"
	"io"
	"os"
)

// Env represents an execution environment.
type Env struct {
	table  *expr.SymTab
	output io.Writer
	input  io.Reader
}

// NewStdEnv returns a new execution environment with an empty
// symbol table, and with IO set to standard output and input.
func NewStdEnv() *Env {
	return &Env{expr.NewTable(), os.Stdout, os.Stdin}
}

// NewEnvWithIO returns a new execution environment with an empty
// symbol table, and with IO set to the specified writer and reader.
func NewEnvWithIO(output io.Writer, input io.Reader) *Env {
	return &Env{expr.NewTable(), output, input}
}
