package day00

import (
	"sort"
	"strconv"
	"strings"

	"github.com/frodejac/aoc-2022/pkg/arraytools"
)

type Day00 struct {
	elves []elf
}

func Solver(input []byte) *Day00 {
	return &Day00{elves: parseInput(string(input))}
}

type elf struct {
	inventory []int
}

func parseInput(input string) []elf {
	input = strings.TrimSpace(input)
	inputs := strings.Split(input, "\n\n")
	elves := make([]elf, len(inputs))
	for i, input := range inputs {
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			} else {
				elves[i].inventory = append(elves[i].inventory, calories)
			}
		}
	}
	return elves
}

func (d *Day00) SolvePart1() string {
	max := 0
	for _, elf := range d.elves {
		sum := arraytools.Sum(elf.inventory)
		if sum > max {
			max = sum
		}
	}
	return strconv.Itoa(max)
}

func (d *Day00) SolvePart2() string {
	invSizes := make([]int, len(d.elves))
	for i, elf := range d.elves {
		invSizes[i] = arraytools.Sum(elf.inventory)
	}
	sort.Slice(invSizes[:], func(i, j int) bool {
		return invSizes[i] > invSizes[j]
	})

	total := arraytools.Sum(invSizes[:3])

	return strconv.Itoa(total)
}
