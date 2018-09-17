package cfg

import "github.com/paulgriffiths/goeval/lar"

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func TestParserOutput(t *testing.T) {
	testCases := []struct {
		infile, cmpfile string
	}{
		{tgArithLr, tgOutArithLrRaw},
		{tgArithNlr, tgOutArithNlrRaw},
		{tgArithAmbig, tgOutArithAmbigRaw},
		{tgBalParens1, tgOutBalParens1Raw},
		{tgBalParens2, tgOutBalParens2Raw},
		{tgZeroOne, tgOutZeroOneRaw},
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
			tgBadMissingHead1,
			parseError{parseErrMissingHead, lar.FilePos{0, 4}},
		},
		{
			tgBadMissingBody1,
			parseError{parseErrEmptyBody, lar.FilePos{8, 4}},
		},
		{
			tgBadMissingBody2,
			parseError{parseErrEmptyBody, lar.FilePos{18, 4}},
		},
		{
			tgBadMissingBody3,
			parseError{parseErrEmptyBody, lar.FilePos{8, 4}},
		},
		{
			tgBadMissingBody4,
			parseError{parseErrEmptyBody, lar.FilePos{8, 5}},
		},
		{
			tgBadENotAlone1,
			parseError{parseErrEmptyNotAlone, lar.FilePos{24, 4}},
		},
		{
			tgBadENotAlone2,
			parseError{parseErrEmptyNotAlone, lar.FilePos{26, 4}},
		},
		{
			tgBadMissingArrow1,
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
