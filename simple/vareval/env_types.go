package vareval

import (
	"io"
	"os"
)

type env struct {
	table  *symTab
	output io.Writer
	input  io.Reader
}

func NewStdEnv() *env {
	return &env{newTable(), os.Stdout, os.Stdin}
}

func NewEnvWithIO(output io.Writer, input io.Reader) *env {
	return &env{newTable(), output, input}
}
