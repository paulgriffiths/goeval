package rdp

import (
	"bytes"
	"github.com/paulgriffiths/goeval/cfg"
	"os"
	"testing"
)

func TestParseWriteTerminals(t *testing.T) {
	testCases := []struct {
		filename, input string
	}{
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"01",
		},
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"0011",
		},
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"000111",
		},
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"00001111",
		},
		{
			"../../cfg/test_grammars/arith_nlr_alnum.grammar",
			"3p4t5",
		},
		{
			"../../cfg/test_grammars/arith_nlr_rev.grammar",
			"(3+4)*(5+8)",
		},
	}

	for n, tc := range testCases {
		infile, err := os.Open(tc.filename)
		if err != nil {
			t.Errorf("couldn't open file %q: %v", tc.filename, err)
			continue
		}

		c, err := cfg.NewCfg(infile)
		infile.Close()
		if err != nil {
			t.Errorf("couldn't create cfg: %v", err)
			continue
		}

		r, err := New(c)
		if err != nil {
			t.Errorf("couldn't create parser: %v", err)
			continue
		}

		tree := r.Parse(tc.input)
		if tree == nil {
			t.Errorf("case %d, couldn't create parse tree", n+1)
		}

		outBuffer := bytes.NewBuffer(nil)
		tree.WriteTerminals(outBuffer)
		terms := string(outBuffer.Bytes())

		if terms != tc.input {
			t.Errorf("case %d, got %q, want %q", n+1, terms, tc.input)
		}
	}
}
