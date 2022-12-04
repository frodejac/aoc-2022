package day03

import (
	"strconv"
	"strings"
)

type Day03 struct {
	input []byte
}

func Solver(input []byte) *Day03 {
	return &Day03{input: input}
}

type elfRange struct {
	start, stop int
}

type elfPair struct {
	range1, range2 elfRange
}

func parseRange(input string) elfRange {
	r := strings.Split(input, "-")
	start, _ := strconv.Atoi(r[0])
	stop, _ := strconv.Atoi(r[1])
	return elfRange{start: start, stop: stop}
}

func parseDay03Input(input string) []*elfPair {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	elfpairs := make([]*elfPair, len(lines))
	for i := 0; i < len(lines); i++ {
		pair := strings.Split(lines[i], ",")
		range1 := parseRange(pair[0])
		range2 := parseRange(pair[1])
		elfpairs[i] = &elfPair{range1: range1, range2: range2}
	}
	return elfpairs
}

func contains(r1, r2 elfRange) bool {
	return r1.start <= r2.start && r2.stop <= r1.stop
}

func (d *Day03) SolvePart1() string {
	pairs := parseDay03Input(string(d.input))

	total := 0
	for i := 0; i < len(pairs); i++ {
		if contains(pairs[i].range1, pairs[i].range2) || contains(pairs[i].range2, pairs[i].range1) {
			total++
		}
	}

	return strconv.Itoa(total)
}

func overlaps(r1, r2 elfRange) bool {
	return r1.start <= r2.start && r2.start <= r1.stop || r2.start <= r1.start && r1.start <= r2.stop
}

func (d *Day03) SolvePart2() string {
	pairs := parseDay03Input(string(d.input))

	total := 0
	for i := 0; i < len(pairs); i++ {
		if overlaps(pairs[i].range1, pairs[i].range2) {
			total++
		}
	}

	return strconv.Itoa(total)
}
