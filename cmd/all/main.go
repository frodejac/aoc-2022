package main

import (
	"fmt"
	"time"

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
		start := time.Now()
		part1 := solver.SolvePart1()
		elapsed := time.Since(start)
		fmt.Printf("Day %02d, part 1: %-15s %10v\n", day, part1, elapsed)
		start = time.Now()
		part2 := solver.SolvePart2()
		elapsed = time.Since(start)
		fmt.Printf("Day %02d, part 2: %-15s %10v\n", day, part2, elapsed)
	}
}
