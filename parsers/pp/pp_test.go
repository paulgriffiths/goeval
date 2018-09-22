package pp

import (
	"github.com/paulgriffiths/goeval/cfg"
	"testing"
)

func TestTable(t *testing.T) {
	grammar, err := cfg.GrammarFromFile(tgArithNlr)
	if err != nil {
		t.Errorf("couldn't get grammar from file %s: %v", tgArithNlr, err)
		return
	}

	pp := NewPp(grammar)

	matrix := [][]int{
		//    +  *  (  )  n  $
		[]int{0, 0, 1, 0, 1, 0}, // E
		[]int{0, 0, 1, 0, 1, 0}, // T
		[]int{1, 0, 0, 1, 0, 1}, // E'
		[]int{0, 0, 1, 0, 1, 0}, // F
		[]int{1, 1, 0, 1, 0, 1}, // T'
		[]int{0, 0, 0, 0, 1, 0}, // Digits
	}

	for i, row := range matrix {
		for j, l := range row {
			v := len(pp.table[i][j])
			if v != l {
				t.Errorf("For (%d,%d), got %d, want %d", i, j, v, l)
			}
		}
	}
}
