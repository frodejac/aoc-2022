package arraytools

import (
	"github.com/frodejac/aoc-2022/pkg/math"
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

func Min[T Number](array []T) T {
	min := array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
	}
	return min
}

func Intersect[T comparable](as, bs []T) []T {
	i := make([]T, 0, math.Max(len(as), len(bs)))
	for _, a := range as {
		for _, b := range bs {
			if a == b {
				i = append(i, a)
			}
		}
	}
	return i
}
