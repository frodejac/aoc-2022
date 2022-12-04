package aoc_test

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/frodejac/aoc-2022/internal/aoc"
	"github.com/frodejac/aoc-2022/internal/io"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestSolveDay03Part1(t *testing.T) {
	const day, part = 3, 1
	input, err := io.GetInput(day)
	if err != nil {
		t.Log("Failed to get input for day 03")
		t.Fail()
	}
	solver := aoc.GetAocSolver(day, input)
	part1 := solver.SolvePart1()
	expected, err := io.GetSolution(day, part)
	if err != nil {
		t.Log("Failed to get solution for day 03, part 1")
		t.Fail()
	}
	if part1 != string(expected) {
		t.Logf("Expected %s, got %s", expected, part1)
		t.Fail()
	}
}

func TestSolveDay03Part2(t *testing.T) {
	const day, part = 3, 2
	input, err := io.GetInput(day)
	if err != nil {
		t.Log("Failed to get input for day 03")
		t.Fail()
	}
	solver := aoc.GetAocSolver(day, input)
	part2 := solver.SolvePart2()
	expected, err := io.GetSolution(day, part)
	if err != nil {
		t.Log("Failed to get solution for day 03, part 2")
		t.Fail()
	}
	if part2 != string(expected) {
		t.Logf("Expected %s, got %s", expected, part2)
		t.Fail()
	}
}
