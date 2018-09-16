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
			"test_grammars/arith_lr.grammar",
			[]int{0, 1, 2, 3},
			[]bool{true, true, false, false},
		},
		{
			"test_grammars/arith_nlr.grammar",
			[]int{0, 1, 2, 3, 4, 5},
			[]bool{false, false, false, false, false, false},
		},
		{
			"test_grammars/arith_ambig.grammar",
			[]int{0, 1},
			[]bool{true, false},
		},
		{
			"test_grammars/example_4_18.grammar",
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
		{"test_grammars/arith_lr.grammar", true},
		{"test_grammars/arith_nlr.grammar", false},
		{"test_grammars/arith_ambig.grammar", true},
		{"test_grammars/bal_parens.grammar", true},
		{"test_grammars/zero_one.grammar", false},
		{"test_grammars/example_4_18.grammar", true},
		{"test_grammars/three_level_lr_1.grammar", true},
		{"test_grammars/three_level_lr_2.grammar", true},
		{"test_grammars/three_level_lr_3.grammar", true},
		{"test_grammars/four_level_nlr_1.grammar", false},
		{"test_grammars/cycle_1.grammar", true},
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
		{"test_grammars/arith_lr.grammar", false},
		{"test_grammars/arith_nlr.grammar", false},
		{"test_grammars/arith_ambig.grammar", false},
		{"test_grammars/bal_parens.grammar", false},
		{"test_grammars/zero_one.grammar", false},
		{"test_grammars/example_4_18.grammar", false},
		{"test_grammars/three_level_lr_1.grammar", false},
		{"test_grammars/three_level_lr_2.grammar", false},
		{"test_grammars/three_level_lr_3.grammar", false},
		{"test_grammars/four_level_nlr_1.grammar", false},
		{"test_grammars/cycle_1.grammar", true},
		{"test_grammars/cycle_2.grammar", true},
		{"test_grammars/cycle_3.grammar", true},
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
		{"test_grammars/arith_lr.grammar", false},
		{"test_grammars/arith_nlr.grammar", true},
		{"test_grammars/arith_ambig.grammar", false},
		{"test_grammars/bal_parens.grammar", true},
		{"test_grammars/zero_one.grammar", false},
		{"test_grammars/example_4_18.grammar", true},
		{"test_grammars/three_level_lr_1.grammar", true},
		{"test_grammars/three_level_lr_2.grammar", true},
		{"test_grammars/three_level_lr_3.grammar", true},
		{"test_grammars/four_level_nlr_1.grammar", true},
		{"test_grammars/cycle_1.grammar", false},
		{"test_grammars/cycle_2.grammar", false},
		{"test_grammars/cycle_3.grammar", false},
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