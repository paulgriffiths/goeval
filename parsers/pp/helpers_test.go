package pp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"github.com/paulgriffiths/goeval/parsers/tree"
	"os"
	"testing"
)

func getParserFromFile(t *testing.T, filename string) (*Pp, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	grammar, err := cfg.NewCfg(infile)
	infile.Close()
	if err != nil {
		return nil, err
	}

	parser := NewPp(grammar)
	if parser == nil {
		return nil, err
	}

	return parser, nil
}

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

	parser := NewPp(grammar)
	if parser == nil {
		return nil, err
	}

	tree := parser.Parse(input)
	if tree == nil {
		return nil, nil
	}

	return tree, nil
}
