package day09

import (
	"strconv"
	"strings"
)

type operation int

const (
	addX = iota
	noop
)

type cpu struct {
	register int
	clock    int
	buf      []int
}

type instruction struct {
	op  operation
	val int
}

func parseInput(input string) []instruction {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "noop":
			instructions[i] = instruction{op: noop}
		case "addx":
			val, _ := strconv.Atoi(parts[1])
			instructions[i] = instruction{op: addX, val: val}
		default:
			continue
		}
	}
	return instructions
}

type Day09 struct {
	instructions []instruction
}

func Solver(input []byte) *Day09 {
	return &Day09{instructions: parseInput(string(input))}
}

func (c *cpu) tick() {
	c.buf = append(c.buf, c.register)
	c.clock++
}

func (c *cpu) run(instructions []instruction) {
	for i := 0; i < len(instructions); i++ {
		switch instructions[i].op {
		case noop:
			c.tick()
		case addX:
			c.tick()
			c.tick()
			c.register += instructions[i].val
		}
	}
}

func (c *cpu) renderCRT() string {
	str := ""
	for i := 0; i < 240; i++ {
		if i%40 == 0 {
			str += "\n"
		}
		diff := i%40 - c.buf[i]
		if diff >= -1 && diff <= 1 {
			str += "#"
		} else {
			str += "."
		}
	}
	return str
}

func (d *Day09) SolvePart1() string {
	cpu := cpu{register: 1}
	cpu.run(d.instructions)

	sum := 0
	for i := 20; i < 240; i += 40 {
		sum += cpu.buf[i-1] * i
	}

	return strconv.Itoa(sum)
}

func (d *Day09) SolvePart2() string {
	cpu := cpu{register: 1}
	cpu.run(d.instructions)
	return cpu.renderCRT()
}
