package aoc

import (
	"strconv"
	"strings"

	"github.com/frodejac/aoc-2022/pkg/arraytools"
	"github.com/frodejac/aoc-2022/pkg/maptools"
)

type Day02 Day

type rucksack struct {
	compartment1 map[string]int
	compartment2 map[string]int
}

var prorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func parseDay02Input(input string) []rucksack {
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

func findDuplicates(r rucksack) string {
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
	rucksacks := parseDay02Input(string(d.input))
	duplicates := make([]string, len(rucksacks))
	total := 0
	for i := 0; i < len(rucksacks); i++ {
		duplicates[i] = findDuplicates(rucksacks[i])
		total += prorities[duplicates[i]]
	}
	return strconv.Itoa(total)
}

func findGroupBadge(group []rucksack) string {
	g := make([][]string, len(group))
	for i := 0; i < len(group); i++ {
		r := maptools.Merge(group[i].compartment1, group[i].compartment2)
		g[i] = maptools.Keys(r)
	}

	i := arraytools.Intersect(g[0], g[1])
	i = arraytools.Intersect(i, g[2])

	if len(i) == 1 {
		return i[0]
	}

	return ""
}

func (d *Day02) SolvePart2() string {
	rucksacks := parseDay02Input(string(d.input))
	total := 0
	for i := 0; i < len(rucksacks); i += 3 {
		group := rucksacks[i : i+3]
		badge := findGroupBadge(group)
		total += prorities[badge]
	}
	return strconv.Itoa(total)
}
