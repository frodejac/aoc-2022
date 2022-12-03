package maptools

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Merge[K comparable, V Number](a, b map[K]V) map[K]V {
	merged := make(map[K]V)
	for k, v := range a {
		merged[k] += v
	}
	for k, v := range b {
		merged[k] += v
	}
	return merged
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
