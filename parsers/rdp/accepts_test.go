package rdp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"os"
	"testing"
)

func TestAccepts(t *testing.T) {
	testCases := []struct {
		filename, input string
		match           bool
	}{
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"01",
			true,
		},
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"0011",
			true,
		},
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"000111",
			true,
		},
		{
			"../../cfg/test_grammars/zero_one.grammar",
			"00001111",
			true,
		},
		{
			"../../cfg/test_grammars/arith_nlr_alnum.grammar",
			"3p4t5",
			true,
		},
		{
			"../../cfg/test_grammars/arith_nlr_rev.grammar",
			"(3+4)*(5+8)",
			true,
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

		r := New(c)
		if r == nil {
			t.Errorf("couldn't create parser")
			continue
		}

		if result := r.Accepts(tc.input); result != tc.match {
			t.Errorf("case %d, got %t, want %t", n+1, result, tc.match)
		}
	}
}
