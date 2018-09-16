package tree

import (
	"fmt"
	"github.com/paulgriffiths/goeval/cfg"
	"io"
)

// Node represents a parse tree node.
type Node struct {
	Comp     cfg.BodyComp
	Value    string
	Children []*Node
}

// NewNode creates a new parse tree node.
func NewNode(comp cfg.BodyComp, value string, children []*Node) *Node {
	node := Node{comp, value, children}
	return &node
}

// WriteTerminals writes the terminals in the parse tree to
// the provided io.Writer.
func (t *Node) WriteTerminals(writer io.Writer) {
	if t.Comp.T == cfg.BodyTerminal {
		writer.Write([]byte(fmt.Sprintf("%s", t.Value)))
	}
	for _, child := range t.Children {
		child.WriteTerminals(writer)
	}
}

// WriteBracketed outputs the parse tree in a bracketed-parse format.
func (t *Node) WriteBracketed(writer io.Writer, opts ...string) {
	qc := ""
	ob := "("
	cb := ")"
	if len(opts) > 0 {
		qc = opts[0]
	}
	if len(opts) > 2 {
		ob = opts[1]
		cb = opts[2]
	}

	switch t.Comp.T {
	case cfg.BodyTerminal:
		writer.Write([]byte(fmt.Sprintf("%s%s%s", qc, t.Value, qc)))
	case cfg.BodyEmpty:
		writer.Write([]byte(fmt.Sprintf("e")))
	case cfg.BodyNonTerminal:
		writer.Write([]byte(fmt.Sprintf("%s%s ", ob, t.Value)))
		for n, child := range t.Children {
			child.WriteBracketed(writer, qc, ob, cb)
			if n < len(t.Children)-1 {
				writer.Write([]byte(fmt.Sprintf(" ")))
			}
		}
		writer.Write([]byte(fmt.Sprintf("%s", cb)))
	}
}
