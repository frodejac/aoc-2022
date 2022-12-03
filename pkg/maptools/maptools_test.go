package maptools

import "testing"

func TestMerge(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"b": 3, "c": 4}
	merged := Merge(a, b)
	if merged["a"] != 1 {
		t.Errorf("Expected merged[\"a\"] to be 1, got %d", merged["a"])
	}
	if merged["b"] != 5 {
		t.Errorf("Expected merged[\"b\"] to be 5, got %d", merged["b"])
	}
	if merged["c"] != 4 {
		t.Errorf("Expected merged[\"c\"] to be 4, got %d", merged["c"])
	}
}

func TestKeys(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	keys := Keys(a)
	if len(keys) != 2 {
		t.Errorf("Expected keys to be [\"a\", \"b\"], got %v", keys)
	}
}
