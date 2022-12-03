package math

import "testing"

func TestMax(t *testing.T) {
	x := 1
	y := 2
	max := Max(x, y)
	if max != 2 {
		t.Errorf("Expected max to be 2, got %d", max)
	}
}
