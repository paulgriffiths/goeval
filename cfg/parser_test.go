package cfg

import "github.com/paulgriffiths/goeval/lar"

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
		return nil, fileErr
	}

	c, perr := parse(infile)
	if perr != nil {
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
			"test_grammars/01_arith_lr.grammar",
			"test_grammars/output/01_arith_lr_raw.grammar",
		},
		{
			"test_grammars/02_arith_nlr.grammar",
			"test_grammars/output/02_arith_nlr_raw.grammar",
		},
		{
			"test_grammars/03_arith_ambig.grammar",
			"test_grammars/output/03_arith_ambig_raw.grammar",
		},
		{
			"test_grammars/04_bal_parens_1.grammar",
			"test_grammars/output/04_bal_parens_1_raw.grammar",
		},
		{
			"test_grammars/05_bal_parens_2.grammar",
			"test_grammars/output/05_bal_parens_2_raw.grammar",
		},
		{
			"test_grammars/06_zero_one.grammar",
			"test_grammars/output/06_zero_one_raw.grammar",
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

func TestParserErrors(t *testing.T) {
	testCases := []struct {
		filename string
		err      parseError
	}{
		{
			"test_grammars/bad/missing_head_1.grammar",
			parseError{parseErrMissingHead, lar.FilePos{0, 4}},
		},
		{
			"test_grammars/bad/missing_body_1.grammar",
			parseError{parseErrEmptyBody, lar.FilePos{8, 4}},
		},
		{
			"test_grammars/bad/missing_body_2.grammar",
			parseError{parseErrEmptyBody, lar.FilePos{18, 4}},
		},
		{
			"test_grammars/bad/missing_body_3.grammar",
			parseError{parseErrEmptyBody, lar.FilePos{8, 4}},
		},
		{
			"test_grammars/bad/missing_body_4.grammar",
			parseError{parseErrEmptyBody, lar.FilePos{8, 5}},
		},
		{
			"test_grammars/bad/e_not_alone_1.grammar",
			parseError{parseErrEmptyNotAlone, lar.FilePos{24, 4}},
		},
		{
			"test_grammars/bad/e_not_alone_2.grammar",
			parseError{parseErrEmptyNotAlone, lar.FilePos{26, 4}},
		},
		{
			"test_grammars/bad/missing_arrow_1.grammar",
			parseError{parseErrMissingArrow, lar.FilePos{1, 4}},
		},
	}

	for n, tc := range testCases {
		_, err := getAndParseFile(t, tc.filename)
		if err != tc.err {
			t.Errorf("case %d, got %v, want %v", n+1, err, tc.err)
		}
	}
}
