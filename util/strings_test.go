package util

import (
	"testing"
)

func TestWordMatchesFill(t *testing.T) {
	w := "AspaRTamE"
	f := "a  a  a  "
	if !WordMatchesFill(w, f) {
		t.Errorf("word should match fill")
	}
}

func TestNormalizeString(t *testing.T) {
	w := "GREATLY"
	n := NormalizeString(w)
	m := "aeglrty"
	if n != m {
		t.Errorf("%v normalized to %v instead of %v", w, n, m)
	}
}

func TestWordMatchesSet(t *testing.T) {
	if !WordMatchesSet("style", "elyts") {
		t.Errorf("word style should match set elyts")
	}
	if !WordMatchesSet("style", "styleeeee") {
		t.Errorf("word should match set that has more letters")
	}
}
