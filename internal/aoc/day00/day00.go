package day00

import (
	"sort"
	"strconv"
	"strings"

	"github.com/frodejac/aoc-2022/pkg/arraytools"
)

type Day00 struct {
	input []byte
}

func Solver(input []byte) *Day00 {
	return &Day00{input: input}
}

type elf struct {
	inventory []int
}

func parseDay00Input(input string) []elf {
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
	elves := parseDay00Input(string(d.input))

	inventories := make([]int, len(elves))
	for i, elf := range elves {
		inventories[i] = arraytools.Sum(elf.inventory)
	}
	max := arraytools.Max(inventories)

	return strconv.Itoa(max)
}

func (d *Day00) SolvePart2() string {
	elves := parseDay00Input(string(d.input))

	inventories := make([]int, len(elves))
	for i, elf := range elves {
		inventories[i] = arraytools.Sum(elf.inventory)
	}
	sort.Slice(inventories[:], func(i, j int) bool {
		return inventories[i] > inventories[j]
	})

	total := arraytools.Sum(inventories[:3])

	return strconv.Itoa(total)

}