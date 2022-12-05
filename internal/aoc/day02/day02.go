package day02

import (
	"strconv"
	"strings"

	"github.com/frodejac/aoc-2022/pkg/arraytools"
	"github.com/frodejac/aoc-2022/pkg/maptools"
)

func getPriorities() map[string]int {
	priorities := make(map[string]int)
	for i, c := range strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "") {
		priorities[c] = i + 1
	}
	return priorities
}

type rucksack struct {
	compartment1 map[string]int
	compartment2 map[string]int
}
type Day02 struct {
	rucksacks  []rucksack
	priorities map[string]int
}

func Solver(input []byte) *Day02 {
	return &Day02{rucksacks: parseInput(string(input)), priorities: getPriorities()}
}

func parseInput(input string) []rucksack {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	rucksacks := make([]rucksack, len(lines))
	for i := 0; i < len(lines); i++ {
		rucksacks[i] = rucksack{compartment1: make(map[string]int), compartment2: make(map[string]int)}
		chars := strings.Split(lines[i], "")
		for j := 0; j < len(chars)/2; j++ {
			rucksacks[i].compartment1[chars[j]]++
		}
		for j := len(chars) / 2; j < len(chars); j++ {
			rucksacks[i].compartment2[chars[j]]++
		}
	}
	return rucksacks
}

func findDuplicate(r rucksack) string {
	for a := range r.compartment1 {
		for b := range r.compartment2 {
			if a == b {
				return a
			}
		}
	}
	return ""
}

func (d *Day02) SolvePart1() string {
	total := 0
	for _, r := range d.rucksacks {
		commonItem := findDuplicate(r)
		total += d.priorities[commonItem]
	}
	return strconv.Itoa(total)
}

func findGroupBadge(group []rucksack) string {
	r := maptools.Merge(group[0].compartment1, group[0].compartment2)
	intersection := maptools.Keys(r)

	for _, g := range group[1:] {
		r := maptools.Merge(g.compartment1, g.compartment2)
		intersection = arraytools.Intersect(intersection, maptools.Keys(r))
	}

	if len(intersection) == 1 {
		return intersection[0]
	}

	return ""
}

func (d *Day02) SolvePart2() string {
	groupSize := 3
	total := 0
	for i := 0; i < len(d.rucksacks); i += groupSize {
		group := d.rucksacks[i : i+groupSize]
		badge := findGroupBadge(group)
		total += d.priorities[badge]
	}
	return strconv.Itoa(total)
}
