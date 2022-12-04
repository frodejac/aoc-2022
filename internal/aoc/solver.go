package aoc

import (
	"github.com/frodejac/aoc-2022/internal/aoc/day00"
	"github.com/frodejac/aoc-2022/internal/aoc/day01"
	"github.com/frodejac/aoc-2022/internal/aoc/day02"
	"github.com/frodejac/aoc-2022/internal/aoc/day03"
	"github.com/frodejac/aoc-2022/internal/aoc/day04"
	"github.com/frodejac/aoc-2022/internal/aoc/day05"
	"github.com/frodejac/aoc-2022/internal/aoc/day06"
	"github.com/frodejac/aoc-2022/internal/aoc/day07"
	"github.com/frodejac/aoc-2022/internal/aoc/day08"
	"github.com/frodejac/aoc-2022/internal/aoc/day09"
	"github.com/frodejac/aoc-2022/internal/aoc/day10"
	"github.com/frodejac/aoc-2022/internal/aoc/day11"
	"github.com/frodejac/aoc-2022/internal/aoc/day12"
	"github.com/frodejac/aoc-2022/internal/aoc/day13"
	"github.com/frodejac/aoc-2022/internal/aoc/day14"
	"github.com/frodejac/aoc-2022/internal/aoc/day15"
	"github.com/frodejac/aoc-2022/internal/aoc/day16"
	"github.com/frodejac/aoc-2022/internal/aoc/day17"
	"github.com/frodejac/aoc-2022/internal/aoc/day18"
	"github.com/frodejac/aoc-2022/internal/aoc/day19"
	"github.com/frodejac/aoc-2022/internal/aoc/day20"
	"github.com/frodejac/aoc-2022/internal/aoc/day21"
	"github.com/frodejac/aoc-2022/internal/aoc/day22"
	"github.com/frodejac/aoc-2022/internal/aoc/day23"
	"github.com/frodejac/aoc-2022/internal/aoc/day24"
)

type AocSolver interface {
	SolvePart1() string
	SolvePart2() string
}

type Day struct {
	input []byte
}

func GetAocSolver(day int, input []byte) AocSolver {
	switch day {
	case 0:
		return day00.Solver(input)
	case 1:
		return day01.Solver(input)
	case 2:
		return day02.Solver(input)
	case 3:
		return day03.Solver(input)
	case 4:
		return day04.Solver(input)
	case 5:
		return day05.Solver(input)
	case 6:
		return day06.Solver(input)
	case 7:
		return day07.Solver(input)
	case 8:
		return day08.Solver(input)
	case 9:
		return day09.Solver(input)
	case 10:
		return day10.Solver(input)
	case 11:
		return day11.Solver(input)
	case 12:
		return day12.Solver(input)
	case 13:
		return day13.Solver(input)
	case 14:
		return day14.Solver(input)
	case 15:
		return day15.Solver(input)
	case 16:
		return day16.Solver(input)
	case 17:
		return day17.Solver(input)
	case 18:
		return day18.Solver(input)
	case 19:
		return day19.Solver(input)
	case 20:
		return day20.Solver(input)
	case 21:
		return day21.Solver(input)
	case 22:
		return day22.Solver(input)
	case 23:
		return day23.Solver(input)
	case 24:
		return day24.Solver(input)
	default:
		panic("Day not implemented")
	}
}
