package day05

import (
	"strconv"
)

type Day05 struct {
	input string
}

func Solver(input []byte) *Day05 {
	return &Day05{input: string(input)}
}

func checkSegmentMarker(sequence string) bool {
	for i := 0; i < len(sequence); i++ {
		for j := i + 1; j < len(sequence); j++ {
			if sequence[i] == sequence[j] {
				return false
			}
		}
	}
	return true
}

func (d *Day05) SolvePart1() string {
	markerPos := 0
	for i := 4; i < len(d.input); i++ {
		if checkSegmentMarker(d.input[i-4 : i]) {
			markerPos = i
			break
		}
	}
	return strconv.Itoa(markerPos)
}

func (d *Day05) SolvePart2() string {
	markerPos := 0
	for i := 14; i < len(d.input); i++ {
		if checkSegmentMarker(d.input[i-14 : i]) {
			markerPos = i
			break
		}
	}
	return strconv.Itoa(markerPos)
}
