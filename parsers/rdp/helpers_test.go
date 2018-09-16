package rdp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"github.com/paulgriffiths/goeval/parsers/tree"
	"os"
	"testing"
)

func getParseTreeFromFile(t *testing.T, filename,
	input string) (*tree.Node, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	grammar, err := cfg.NewCfg(infile)
	infile.Close()
	if err != nil {
		return nil, err
	}

	parser, err := New(grammar)
	if err != nil {
		return nil, err
	}

	tree := parser.Parse(input)
	if tree == nil {
		return nil, nil
	}

	return tree, nil
}
