package vareval

import (
	"bufio"
	"bytes"
	"testing"
)

func TestOutput(t *testing.T) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	e := NewEnvWithIO(writer, nil)
	p := output{intValue{35}}
	if err := p.execute(e); err != nil {
		t.Errorf("couldn't execute statement: %v", err)
		return
	}

	writer.Flush()
	if buffer.String() != "35\n" {
		t.Errorf("got %s, want %s", buffer.String(), "35\n")
	}
}
