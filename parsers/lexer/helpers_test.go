package lexer

import (
	"github.com/paulgriffiths/goeval/cfg"
	"os"
	"testing"
)

func getAndParseFile(t *testing.T, filename string) (*cfg.Cfg, error) {
	infile, fileErr := os.Open(filename)
	if fileErr != nil {
		return nil, fileErr
	}

	c, perr := cfg.NewCfg(infile)
	if perr != nil {
		return nil, perr
	}

	infile.Close()

	return c, nil
}
