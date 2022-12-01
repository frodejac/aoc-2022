package aoc

import (
	"strconv"
	"strings"
)

type Day00 Day

type elf struct {
	inventory []int
}

func parseInput(input string) []elf {
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

func sumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func (d Day00) SolvePart1() string {
	elves := parseInput(string(d.input))
	max := 0
	for _, elf := range elves {
		sum := sumArray(elf.inventory)
		if sum > max {
			max = sum
		}
	}
	return strconv.Itoa(max)
}

func sortElves(elves []elf) []elf {
	for i := 0; i < len(elves); i++ {
		for j := i + 1; j < len(elves); j++ {
			if sumArray(elves[i].inventory) < sumArray(elves[j].inventory) {
				elves[i], elves[j] = elves[j], elves[i]
			}
		}
	}
	return elves
}

func (d Day00) SolvePart2() string {
	elves := parseInput(string(d.input))
	elves = sortElves(elves)
	total := 0
	for i := 0; i < len(elves[:3]); i++ {
		total += sumArray(elves[i].inventory)
	}
	return strconv.Itoa(total)

}
