package aoc

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
		return &Day00{input: input}
	case 1:
		return &Day01{input: input}
	case 2:
		return &Day02{input: input}
	case 3:
		return &Day03{input: input}
	case 4:
		return &Day04{input: input}
	case 5:
		return &Day05{input: input}
	case 6:
		return &Day06{input: input}
	case 7:
		return &Day07{input: input}
	case 8:
		return &Day08{input: input}
	case 9:
		return &Day09{input: input}
	case 10:
		return &Day10{input: input}
	case 11:
		return &Day11{input: input}
	case 12:
		return &Day12{input: input}
	case 13:
		return &Day13{input: input}
	case 14:
		return &Day14{input: input}
	case 15:
		return &Day15{input: input}
	case 16:
		return &Day16{input: input}
	case 17:
		return &Day17{input: input}
	case 18:
		return &Day18{input: input}
	case 19:
		return &Day19{input: input}
	case 20:
		return &Day20{input: input}
	case 21:
		return &Day21{input: input}
	case 22:
		return &Day22{input: input}
	case 23:
		return &Day23{input: input}
	case 24:
		return &Day24{input: input}
	default:
		panic("Day not implemented")
	}

}
