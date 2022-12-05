package day03

import (
	"strconv"
	"strings"
)

type elfRange struct {
	start, stop int
}

type elfPair struct {
	range1, range2 elfRange
}

type Day03 struct {
	pairs []*elfPair
}

func parseInput(input string) []*elfPair {
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

func Solver(input []byte) *Day03 {
	return &Day03{pairs: parseInput(string(input))}
}

func parseRange(input string) elfRange {
	r := strings.Split(input, "-")
	start, _ := strconv.Atoi(r[0])
	stop, _ := strconv.Atoi(r[1])
	return elfRange{start: start, stop: stop}
}

func contains(r1, r2 elfRange) bool {
	return r1.start <= r2.start && r2.stop <= r1.stop
}

func (d *Day03) SolvePart1() string {
	total := 0
	for i := 0; i < len(d.pairs); i++ {
		if contains(d.pairs[i].range1, d.pairs[i].range2) || contains(d.pairs[i].range2, d.pairs[i].range1) {
			total++
		}
	}

	return strconv.Itoa(total)
}

func overlaps(r1, r2 elfRange) bool {
	return r1.start <= r2.start && r2.start <= r1.stop || r2.start <= r1.start && r1.start <= r2.stop
}

func (d *Day03) SolvePart2() string {
	total := 0
	for i := 0; i < len(d.pairs); i++ {
		if overlaps(d.pairs[i].range1, d.pairs[i].range2) {
			total++
		}
	}

	return strconv.Itoa(total)
}
