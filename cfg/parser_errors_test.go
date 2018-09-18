package cfg

import (
	"github.com/paulgriffiths/goeval/lar"
	"testing"
)

func TestParserErrors(t *testing.T) {
	testCases := []struct {
		filename string
		err      parseError
	}{
		{tgBadMissingHead1,
			parseError{parseErrMissingHead, lar.FilePos{0, 4}}},
		{tgBadMissingBody1,
			parseError{parseErrEmptyBody, lar.FilePos{8, 4}}},
		{tgBadMissingBody2,
			parseError{parseErrEmptyBody, lar.FilePos{18, 4}}},
		{tgBadMissingBody3,
			parseError{parseErrEmptyBody, lar.FilePos{8, 4}}},
		{tgBadMissingBody4,
			parseError{parseErrEmptyBody, lar.FilePos{8, 5}}},
		{tgBadENotAlone1,
			parseError{parseErrEmptyNotAlone, lar.FilePos{24, 4}}},
		{tgBadENotAlone2,
			parseError{parseErrEmptyNotAlone, lar.FilePos{26, 4}}},
		{tgBadMissingArrow1,
			parseError{parseErrMissingArrow, lar.FilePos{1, 4}}},
	}

	for n, tc := range testCases {
		if _, err := getAndParseFile(t, tc.filename); err != tc.err {
			t.Errorf("case %d, got %v, want %v", n+1, err, tc.err)
		}
	}
}
