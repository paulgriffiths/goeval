package cfg

import "os"

func GrammarFromFile(filename string) (*Cfg, error) {
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
