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
