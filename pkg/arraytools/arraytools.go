package arraytools

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](array []T) T {
	sum := T(0)
	for _, value := range array {
		sum += value
	}
	return sum
}

func Max[T Number](array []T) T {
	max := array[0]
	for _, value := range array {
		if value > max {
			max = value
		}
	}
	return max
}
