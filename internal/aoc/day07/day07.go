package day07

import (
	"strconv"
	"strings"
)

type tree struct {
	height      int
	visible     bool
	scenicScore int
}

type Day07 struct {
	forest [][]tree
}

func parseInput(input string) [][]tree {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	trees := make([][]tree, len(lines))
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], "")
		trees[i] = make([]tree, len(line))
		for j := 0; j < len(line); j++ {
			height, _ := strconv.Atoi(line[j])
			trees[i][j] = tree{height: height}
		}
	}
	return trees
}

func Solver(input []byte) *Day07 {
	return &Day07{forest: parseInput(string(input))}
}

func visibleUp(forest [][]tree, i, j int) bool {
	for k := i - 1; k >= 0; k-- {
		if forest[k][j].height >= forest[i][j].height {
			break
		}
		if k == 0 {
			return true
		}
	}
	return false
}

func visibleDown(forest [][]tree, i, j int) bool {
	for k := i + 1; k < len(forest); k++ {
		if forest[k][j].height >= forest[i][j].height {
			break
		}
		if k == len(forest)-1 {
			return true
		}
	}
	return false
}

func visibleLeft(forest [][]tree, i, j int) bool {
	for k := j - 1; k >= 0; k-- {
		if forest[i][k].height >= forest[i][j].height {
			break
		}
		if k == 0 {
			return true
		}
	}
	return false
}

func visibleRight(forest [][]tree, i, j int) bool {
	for k := j + 1; k < len(forest[i]); k++ {
		if forest[i][k].height >= forest[i][j].height {
			break
		}
		if k == len(forest[i])-1 {
			return true
		}
	}
	return false
}

func (d *Day07) calculateVisibility() {
	for i := 0; i < len(d.forest); i++ {
		for j := 0; j < len(d.forest[i]); j++ {
			if i == 0 || j == 0 || i == len(d.forest)-1 || j == len(d.forest[i])-1 {
				d.forest[i][j].visible = true
			} else {
				d.forest[i][j].visible = visibleUp(d.forest, i, j) || visibleDown(d.forest, i, j) || visibleLeft(d.forest, i, j) || visibleRight(d.forest, i, j)
			}
		}
	}
}

func (d *Day07) SolvePart1() string {
	d.calculateVisibility()
	total := 0
	for i := 0; i < len(d.forest); i++ {
		for j := 0; j < len(d.forest[i]); j++ {
			if d.forest[i][j].visible {
				total++
			}
		}
	}
	return strconv.Itoa(total)
}

func viewDistanceUp(forest [][]tree, i, j int) int {
	score := 0
	for k := i - 1; k >= 0; k-- {
		score++
		if forest[k][j].height >= forest[i][j].height {
			break
		}
	}
	return score
}

func viewDistanceDown(forest [][]tree, i, j int) int {
	score := 0
	for k := i + 1; k < len(forest); k++ {
		score++
		if forest[k][j].height >= forest[i][j].height {
			break
		}
	}
	return score
}

func viewDistanceLeft(forest [][]tree, i, j int) int {
	score := 0
	for k := j - 1; k >= 0; k-- {
		score++
		if forest[i][k].height >= forest[i][j].height {
			break
		}
	}
	return score
}

func viewDistanceRight(forest [][]tree, i, j int) int {
	score := 0
	for k := j + 1; k < len(forest[i]); k++ {
		score++
		if forest[i][k].height >= forest[i][j].height {
			break
		}
	}
	return score
}

func (d *Day07) calculateScenicScores() {
	for i := 0; i < len(d.forest); i++ {
		for j := 0; j < len(d.forest[i]); j++ {
			scenicScore := viewDistanceUp(d.forest, i, j) * viewDistanceDown(d.forest, i, j) * viewDistanceLeft(d.forest, i, j) * viewDistanceRight(d.forest, i, j)
			d.forest[i][j].scenicScore = scenicScore
		}
	}
}

func (d *Day07) SolvePart2() string {
	d.calculateScenicScores()
	max := 0
	for i := 0; i < len(d.forest); i++ {
		for j := 0; j < len(d.forest[i]); j++ {
			if d.forest[i][j].scenicScore > max {
				max = d.forest[i][j].scenicScore
			}
		}
	}
	return strconv.Itoa(max)
}
