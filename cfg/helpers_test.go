package cfg

import (
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
