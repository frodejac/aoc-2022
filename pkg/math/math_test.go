package math

import "testing"

func TestMax(t *testing.T) {
	x := 1
	y := 2
	max := Max(x, y)
	if max != 2 {
		t.Errorf("Expected max to be 2, got %d", max)
	}

	x = 2
	max = Max(x, y)
	if max != 2 {
		t.Errorf("Expected max to be 2, got %d", max)
	}
}

func TestMin(t *testing.T) {
	x := 1
	y := 2
	min := Min(x, y)
	if min != 1 {
		t.Errorf("Expected min to be 1, got %d", min)
	}

	x = 2
	min = Min(x, y)
	if min != 2 {
		t.Errorf("Expected min to be 2, got %d", min)
	}
}
