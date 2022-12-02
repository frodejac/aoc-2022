package arraytools

import "testing"

func TestSum(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	sum := Sum(array)
	if sum != 15 {
		t.Errorf("Expected sum to be 15, got %d", sum)
	}
}

func TestMax(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	max := Max(array)
	if max != 5 {
		t.Errorf("Expected max to be 5, got %d", max)
	}
}
