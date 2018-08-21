package vareval

import (
	"bufio"
	"bytes"
	"github.com/paulgriffiths/goeval/expr"
	"testing"
)

func TestOutput(t *testing.T) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	e := NewEnvWithIO(writer, nil)
	p := outputStmt{expr.NewInt(35)}
	if err := p.Execute(e); err != nil {
		t.Errorf("couldn't execute statement: %v", err)
		return
	}

	writer.Flush()
	if buffer.String() != "35\n" {
		t.Errorf("got %s, want %s", buffer.String(), "35\n")
	}
}

func TestAssignmemt(t *testing.T) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	e := NewEnvWithIO(writer, nil)
	p := assignStmt{"foobar", expr.NewInt(35)}
	if err := p.Execute(e); err != nil {
		t.Errorf("couldn't execute statement: %v", err)
		return
	}

	val, ok := e.table.Retrieve("foobar")
	if !ok {
		t.Errorf("couldn't retrieve value")
		return
	}
	if val != expr.NewInt(35) {
		t.Errorf("got %v, want %v", val, expr.NewInt(35))
	}
}
