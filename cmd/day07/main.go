package main

import (
	"fmt"

	"github.com/frodejac/aoc-2022/internal/aoc"
	"github.com/frodejac/aoc-2022/internal/io"
)

const DAY = 7

func main() {
	data, err := io.GetInput(DAY)
	if err != nil {
		panic(err)
	}
	solver := aoc.GetAocSolver(DAY, data)
	part1 := solver.SolvePart1()
	fmt.Printf("Day %02d, part 1: %s\n", DAY, part1)
	part2 := solver.SolvePart2()
	fmt.Printf("Day %02d, part 2: %s\n", DAY, part2)
}
