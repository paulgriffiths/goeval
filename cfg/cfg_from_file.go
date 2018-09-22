package cfg

import "os"

// GrammarFromFile constructs a context-free grammar object from
// a representation in a text file.
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
