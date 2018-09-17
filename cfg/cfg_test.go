package cfg

import (
	"testing"
)

func TestImmediateLeftRecursive(t *testing.T) {
	testCases := []struct {
		filename string
		nt       []int
		result   []bool
	}{
		{
			tgArithLr,
			[]int{0, 1, 2, 3},
			[]bool{true, true, false, false},
		},
		{
			tgArithNlr,
			[]int{0, 1, 2, 3, 4, 5},
			[]bool{false, false, false, false, false, false},
		},
		{
			tgArithAmbig,
			[]int{0, 1},
			[]bool{true, false},
		},
		{
			tgIndirectLr1,
			[]int{0, 1},
			[]bool{false, true},
		},
	}

	for n, tc := range testCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		for i, nt := range tc.nt {
			if r := c.IsImmediateLeftRecursive(nt); r != tc.result[i] {
				t.Errorf("case %d nt %d, got %t, want %t",
					n+1, nt, r, tc.result[i])
			}
		}
	}
}

func TestIsLeftRecursive(t *testing.T) {
	testCases := []struct {
		filename string
		result   bool
	}{
		{tgArithLr, true},
		{tgArithNlr, false},
		{tgArithAmbig, true},
		{tgBalParens1, true},
		{tgBalParens2, false},
		{tgZeroOne, false},
		{tgIndirectLr1, true},
		{tgIndirectLr2, true},
		{tgIndirectLr3, true},
		{tgCycle1, true},
		{tgCycle2, true},
		{tgCycle3, true},
		{tgCycle4, true},
	}

	for n, tc := range testCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		if r := c.IsLeftRecursive(); r != tc.result {
			t.Errorf("case %d, got %t, want %t", n+1, r, tc.result)
		}
	}
}

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		filename string
		result   bool
	}{
		{tgArithLr, false},
		{tgArithNlr, false},
		{tgArithAmbig, false},
		{tgBalParens1, false},
		{tgBalParens2, false},
		{tgZeroOne, false},
		{tgIndirectLr1, false},
		{tgIndirectLr2, false},
		{tgIndirectLr3, false},
		{tgCycle1, true},
		{tgCycle2, true},
		{tgCycle3, true},
		{tgCycle4, true},
	}

	for n, tc := range testCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		if r := c.HasCycle(); r != tc.result {
			t.Errorf("case %d, got %t, want %t", n+1, r, tc.result)
		}
	}
}

func TestHasEProduction(t *testing.T) {
	testCases := []struct {
		filename string
		result   bool
	}{
		{tgArithLr, false},
		{tgArithNlr, true},
		{tgArithAmbig, false},
		{tgBalParens1, true},
		{tgBalParens2, true},
		{tgZeroOne, false},
		{tgIndirectLr1, true},
		{tgIndirectLr2, true},
		{tgIndirectLr3, true},
		{tgCycle1, false},
		{tgCycle2, false},
		{tgCycle3, false},
		{tgCycle4, false},
	}

	for n, tc := range testCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		if r := c.HasEProduction(); r != tc.result {
			t.Errorf("case %d, got %t, want %t", n+1, r, tc.result)
		}
	}
}

func TestIsNullable(t *testing.T) {
	testCases := []struct {
		filename string
		nt       []int
		result   []bool
	}{
		{
			tgArithNlr,
			[]int{0, 1, 2, 3, 4, 5},
			[]bool{false, false, true, false, true, false},
		},
		{
			tgNullable1,
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]bool{true, false, false, true, true, false, false},
		},
		{
			tgNullable2,
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]bool{true, false, true, true, true, false, false},
		},
		{
			tgNullable3,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
			[]bool{true, false, true, true, true, true, false, false},
		},
	}

	for n, tc := range testCases {
		c, err := getAndParseFile(t, tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		for i, nt := range tc.nt {
			if r := c.IsNullable(nt); r != tc.result[i] {
				t.Errorf("case %d nt %d, got %t, want %t",
					n+1, nt, r, tc.result[i])
			}
		}
	}
}
