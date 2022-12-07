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

func TestMin(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	min := Min(array)
	if min != 1 {
		t.Errorf("Expected min to be 1, got %d", min)
	}
}

func TestIntersect(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	intersection := Intersect(a, b)
	if len(intersection) != 3 {
		t.Errorf("Expected intersection to be [3, 4, 5], got %v", intersection)
	}
}
