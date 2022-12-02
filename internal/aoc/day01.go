package aoc

import (
	"strconv"
	"strings"
)

type Day01 Day

type rpsRound struct {
	elf    int
	player int
}

var rpsMoves = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
	"A": 1,
	"B": 2,
	"C": 3,
}

var roundScore = map[int]int{
	0: 3,
	1: 0,
	2: 6,
}

var roundMod = map[int]int{
	1: 1,
	2: -1,
	3: 0,
}

func parseDay01Input(input string) []rpsRound {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	rounds := make([]rpsRound, len(lines))
	for i := 0; i < len(lines); i++ {
		plays := strings.Split(lines[i], " ")
		rounds[i] = rpsRound{elf: rpsMoves[plays[0]], player: rpsMoves[plays[1]]}
	}
	return rounds
}

func (d *Day01) SolvePart1() string {
	total := 0
	for _, r := range parseDay01Input(string(d.input)) {
		total += roundScore[(r.elf-r.player+3)%3] + r.player
	}
	return strconv.Itoa(total)
}

func (d *Day01) SolvePart2() string {
	total := 0
	for _, r := range parseDay01Input(string(d.input)) {
		total += (r.elf+roundMod[r.player])%3 + 1 + (r.player-1)*3
	}
	return strconv.Itoa(total)
}
