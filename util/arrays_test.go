package util

import (
	"testing"
)

func TestArrayRemove(t *testing.T) {
	a := []string{"a", "b", "d", "c"}
	a, b := ArrayRemove(a, 2)
	if b != "d" {
		t.Errorf("method removed %v instead of d from array", b)
	}
	if len(a) != 3 {
		t.Errorf("method didn't actually change array")
	}
}
