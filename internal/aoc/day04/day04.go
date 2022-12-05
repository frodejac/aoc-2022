package day04

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/frodejac/aoc-2022/pkg/datastructures/stack"
)

type Day04 struct {
	cb    *cargoBay
	moves []move
}

type cargoBay struct {
	stacks []*stack.Stack[string]
}

type move struct {
	from, to int
	count    int
}

func (c *cargoBay) move9000(from, to, count int) {
	for i := 0; i < count; i++ {
		c.stacks[to].Push(c.stacks[from].Pop())
	}
}

func (c *cargoBay) move9001(from, to, count int) {
	values := c.stacks[from].PopN(count)
	c.stacks[to].PushN(values)
}

func (c *cargoBay) toString() string {
	var sb strings.Builder
	for _, stack := range c.stacks {
		sb.WriteString(stack.Peek())
	}
	return sb.String()
}

func newCargoBay(stackCount int) *cargoBay {
	stacks := make([]*stack.Stack[string], stackCount)
	for i := 0; i < stackCount; i++ {
		stacks[i] = stack.New([]string{})
	}
	return &cargoBay{stacks: stacks}
}

func parseCargoBay(input string) *cargoBay {
	lines := strings.SplitAfter(input, "\n")
	stackCount := len(lines[0]) / 4
	cb := newCargoBay(stackCount)

	for _, line := range lines[:len(lines)-1] {
		for i := 0; i < stackCount; i++ {
			char := line[i*4 : i*4+4]
			char = strings.TrimSpace(char)
			char = strings.TrimLeft(char, "[")
			char = strings.TrimRight(char, "]")
			if char != "" {
				cb.stacks[i].PushLeft(char)
			}
		}
	}
	return cb
}

func parseMoves(input string) []move {
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := r.FindAllStringSubmatch(input, -1)
	moves := make([]move, len(matches))
	for i, match := range matches {
		count, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])
		moves[i] = move{from: from - 1, to: to - 1, count: count}
	}
	return moves
}

func parseInput(input string) (*cargoBay, []move) {
	sections := strings.Split(input, "\n\n")
	cargoBay := parseCargoBay(sections[0])
	moves := parseMoves(sections[1])
	return cargoBay, moves
}

func Solver(input []byte) *Day04 {
	cb, moves := parseInput(string(input))
	return &Day04{cb: cb, moves: moves}
}

func (d *Day04) SolvePart1() string {
	for _, move := range d.moves {
		d.cb.move9000(move.from, move.to, move.count)
	}
	return d.cb.toString()
}

func (d *Day04) SolvePart2() string {
	for _, move := range d.moves {
		d.cb.move9001(move.from, move.to, move.count)
	}
	return d.cb.toString()
}
