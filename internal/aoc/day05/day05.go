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

func isSegmentMarker(sequence string) bool {
	for i := 0; i < len(sequence); i++ {
		for j := i + 1; j < len(sequence); j++ {
			if sequence[i] == sequence[j] {
				return false
			}
		}
	}
	return true
}

func findSegmentMarker(sequence string, markerLength int) int {
	for i := markerLength; i < len(sequence); i++ {
		if isSegmentMarker(sequence[i-markerLength : i]) {
			return i
		}
	}
	return -1
}

func (d *Day05) SolvePart1() string {
	markerPos := findSegmentMarker(d.input, 4)
	return strconv.Itoa(markerPos)
}

func (d *Day05) SolvePart2() string {
	markerPos := findSegmentMarker(d.input, 14)
	return strconv.Itoa(markerPos)
}
