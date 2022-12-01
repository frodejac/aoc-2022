package main

import (
	"fmt"

	"github.com/frodejac/aoc-2022/internal/aoc"
	"github.com/frodejac/aoc-2022/internal/io"
)

func main() {
	for day := 0; day < 24; day++ {
		data, err := io.GetInput(day)
		if err != nil {
			fmt.Printf("Day %02d: Failed to get input data, skipping...\n", day)
			continue
		}
		solver := aoc.GetAocSolver(day, data)
		part1 := solver.SolvePart1()
		fmt.Printf("Day %02d, part 1: %s\n", day, part1)
		part2 := solver.SolvePart2()
		fmt.Printf("Day %02d, part 2: %s\n", day, part2)
	}
}
