package rdp

import (
	"bytes"
	"testing"
)

func TestParseWriteBracketed(t *testing.T) {
	testCases := []struct {
		filename, input string
		opts            []string
		output          string
	}{
		{
			"../../cfg/test_grammars/02_arith_nlr.grammar",
			"(3+4)*5",
			[]string{"", "[", "]"},
			"[E [T [F ( [E [T [F [Digits 3]] [T' e]] [E' + " +
				"[T [F [Digits 4]] [T' e]] [E' e]]] )] [T' * [F " +
				"[Digits 5]] [T' e]]] [E' e]]",
		},
		{
			"../../cfg/test_grammars/05_bal_parens_2.grammar",
			"((()))",
			[]string{"", "[", "]"},
			"[S ( [S ( [S ( [S e] ) [S e]] ) [S e]] ) [S e]]",
		},
		{
			"../../cfg/test_grammars/06_zero_one.grammar",
			"00001111",
			[]string{},
			"(S 0 (S 0 (S 0 (S 01) 1) 1) 1)",
		},
	}

	for n, tc := range testCases {
		tree, err := getParseTreeFromFile(t, tc.filename, tc.input)
		if err != nil {
			t.Errorf("couldn't get parse tree for file %q: %v",
				tc.filename, err)
			continue
		} else if tree == nil {
			t.Errorf("case %d, failed to parse", n+1)
			continue
		}

		outBuffer := bytes.NewBuffer(nil)
		tree.WriteBracketed(outBuffer, tc.opts...)
		output := string(outBuffer.Bytes())

		if output != tc.output {
			t.Errorf("case %d, got %q, want %q", n+1, output, tc.output)
		}
	}
}
