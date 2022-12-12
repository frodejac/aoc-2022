package day11

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type Day11 struct {
	heightMap [][]int
	startPos  point
	endPos    point
}

func parseInput(input string) ([][]int, point, point) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	var startPos, endPos point
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		for j, char := range line {
			if char == 'S' {
				startPos = point{x: i, y: j}
				heightMap[i][j] = 0
			} else if char == 'E' {
				endPos = point{x: i, y: j}
				heightMap[i][j] = int('z') - 97
			} else {
				heightMap[i][j] = int(char) - 97
			}
		}
	}

	return heightMap, startPos, endPos
}

func Solver(input []byte) *Day11 {
	heightMap, startPos, endPos := parseInput(string(input))
	return &Day11{heightMap: heightMap, startPos: startPos, endPos: endPos}
}

func dijkstra(w, h int, startPos point, finished func(point) bool, hasPath func(point, point) bool) int16 {
	dist := make([][]int16, w)
	visited := make([][]bool, w)
	for i := range dist {
		visited[i] = make([]bool, h)
		dist[i] = make([]int16, h)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt16
		}
	}
	dist[startPos.x][startPos.y] = 0
	queue := []point{startPos}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Skip visited nodes
		if visited[curr.x][curr.y] {
			continue
		} else {
			visited[curr.x][curr.y] = true
		}

		// Check if we're done
		if finished(curr) {
			return dist[curr.x][curr.y]
		}

		for _, neighbor := range []point{
			{x: curr.x - 1, y: curr.y},
			{x: curr.x + 1, y: curr.y},
			{x: curr.x, y: curr.y - 1},
			{x: curr.x, y: curr.y + 1},
		} {
			// Out of bounds check
			if neighbor.x < 0 || neighbor.x >= w || neighbor.y < 0 || neighbor.y >= h {
				continue
			}

			// Skip inaccessible neighbors
			if !hasPath(curr, neighbor) {
				continue
			}

			// Update shortest path distance
			alt := dist[curr.x][curr.y] + 1
			if alt <= dist[neighbor.x][neighbor.y] {
				dist[neighbor.x][neighbor.y] = alt
			}
			queue = append(queue, neighbor)
		}

		// Sort queue by distance to get a priority queue
		sort.Slice(queue[:], func(i, j int) bool {
			return dist[queue[i].x][queue[i].y] < dist[queue[j].x][queue[j].y]
		})
	}
	return -1
}

func (d *Day11) SolvePart1() string {
	finished := func(p point) bool {
		return p.x == d.endPos.x && p.y == d.endPos.y
	}
	hasPath := func(p1, p2 point) bool {
		return (d.heightMap[p2.x][p2.y] - d.heightMap[p1.x][p1.y]) <= 1
	}
	w, h := len(d.heightMap), len(d.heightMap[0])
	cost := dijkstra(w, h, d.startPos, finished, hasPath)
	return strconv.Itoa(int(cost))
}

func (d *Day11) SolvePart2() string {
	finished := func(p point) bool {
		return d.heightMap[p.x][p.y] == 0
	}
	hasPath := func(p1, p2 point) bool {
		return (d.heightMap[p1.x][p1.y] - d.heightMap[p2.x][p2.y]) <= 1
	}
	w, h := len(d.heightMap), len(d.heightMap[0])
	cost := dijkstra(w, h, d.endPos, finished, hasPath)
	return strconv.Itoa(int(cost))
}
