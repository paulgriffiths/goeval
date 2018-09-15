package cfg

import (
	"bufio"
	"bytes"
	_ "fmt"
	"os"
	"testing"
)

func stringArraysEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for n, s := range a {
		if s != b[n] {
			return false
		}
	}

	return true
}

func getAndParseFile(t *testing.T, filename string) (*Cfg, error) {
	infile, fileErr := os.Open(filename)
	if fileErr != nil {
		t.Errorf("couldn't open file %q: %v", filename, fileErr)
		return nil, fileErr
	}

	c, perr := parse(infile)
	if perr != nil {
		t.Errorf("couldn't get tokens for file %q: %v", filename, perr)
		return nil, perr
	}

	infile.Close()

	return c, nil
}

func TestParserOutput(t *testing.T) {
	testCases := []struct {
		infile, cmpfile string
	}{
		{
			"test_grammars/arith_lr.grammar",
			"test_grammars/output/arith_lr_raw.grammar",
		},
		{
			"test_grammars/arith_nlr.grammar",
			"test_grammars/output/arith_nlr_raw.grammar",
		},
		{
			"test_grammars/arith_ambig.grammar",
			"test_grammars/output/arith_ambig_raw.grammar",
		},
		{
			"test_grammars/zero_one.grammar",
			"test_grammars/output/zero_one_raw.grammar",
		},
		{
			"test_grammars/bal_parens.grammar",
			"test_grammars/output/bal_parens_raw.grammar",
		},
	}

	for _, tc := range testCases {
		c, err := getAndParseFile(t, tc.infile)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.infile, err)
			continue
		}

		outBuffer := bytes.NewBuffer(nil)
		c.outputCfg(outBuffer)
		outScanner := bufio.NewScanner(outBuffer)

		infile, fileErr := os.Open(tc.cmpfile)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", tc.infile, fileErr)
			continue
		}

		cmpScanner := bufio.NewScanner(infile)

		for cmpScanner.Scan() {
			if !outScanner.Scan() {
				t.Errorf("fewer lines than %q", tc.infile)
				break
			}
			if cmpScanner.Text() != outScanner.Text() {
				t.Errorf("%q: got %q, want %q", tc.infile,
					outScanner.Text(), cmpScanner.Text())
			}
		}

		if outScanner.Scan() {
			t.Errorf("more lines than %q", tc.infile)
		}

		infile.Close()
	}
}
