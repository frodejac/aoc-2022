package day08

import (
	"math"
	"strconv"
	"strings"
)

type direction int

const (
	up = iota
	down
	left
	right
)

type move struct {
	dir direction
	dst int
}

type point struct {
	X int
	Y int
}
type knot struct {
	position point
	visited  map[point]int
}

type rope struct {
	knots []knot
}

type Day08 struct {
	moves []move
}

func parseInput(raw []byte) []move {
	input := string(raw)
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	moves := make([]move, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		var dir direction
		switch parts[0] {
		case "U":
			dir = up
		case "D":
			dir = down
		case "L":
			dir = left
		case "R":
			dir = right
		}
		distance, _ := strconv.Atoi(parts[1])
		moves[i] = move{
			dir: dir,
			dst: distance,
		}
	}
	return moves
}

func Solver(input []byte) *Day08 {
	return &Day08{
		moves: parseInput(input),
	}
}

func (k *knot) move(d direction) {
	switch d {
	case up:
		k.position.Y++
	case down:
		k.position.Y--
	case left:
		k.position.X--
	case right:
		k.position.X++
	}
	k.visited[k.position]++
}

func (k *knot) update(headX, headY int) {
	dx := headX - k.position.X
	dy := headY - k.position.Y

	// Don't move tail if it's already adjacent to head
	if dx >= -1 && dx <= 1 && dy >= -1 && dy <= 1 {
		return
	}

	// Move tail in direction of head
	if math.Abs(float64(dx)) > math.Abs(float64(dy)) {
		if dx > 0 {
			k.position.X++
		} else {
			k.position.X--
		}
		// Also update row/column if necessary
		if k.position.Y != headY {
			if dy > 0 {
				k.position.Y++
			} else {
				k.position.Y--
			}
		}
	} else {
		if dy > 0 {
			k.position.Y++
		} else {
			k.position.Y--
		}
		// Also update row/column if necessary
		if k.position.X != headX {
			if dx > 0 {
				k.position.X++
			} else {
				k.position.X--
			}
		}
	}
	k.visited[k.position]++
}

func (r *rope) move(d direction, dst int) {
	for i := 0; i < dst; i++ {
		r.head().move(d)
		for j := 1; j < len(r.knots); j++ {
			r.knots[j].update(r.knots[j-1].position.X, r.knots[j-1].position.Y)
		}
	}
}

func (r *rope) tail() *knot {
	return &r.knots[len(r.knots)-1]
}

func (r *rope) head() *knot {
	return &r.knots[0]
}

func getRope(len int) rope {
	r := rope{knots: make([]knot, len)}
	for i := 0; i < len; i++ {
		r.knots[i] = knot{position: point{X: 0, Y: 0}, visited: map[point]int{}}
	}
	return r
}

func (d *Day08) SolvePart1() string {
	r := getRope(2)
	for _, move := range d.moves {
		r.move(move.dir, move.dst)
	}
	return strconv.Itoa(len(r.tail().visited))
}

func (d *Day08) SolvePart2() string {
	r := getRope(10)
	for _, move := range d.moves {
		r.move(move.dir, move.dst)
	}
	return strconv.Itoa(len(r.tail().visited))
}
