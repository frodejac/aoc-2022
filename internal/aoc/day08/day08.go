package day08

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
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

func (r *rope) visualize(w, h int, palette []color.Color) *image.Paletted {
	centerX, centerY := w/2, h/2
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
	for i := len(r.knots) - 1; i >= 0; i-- {
		k := r.knots[i]
		knotX, knotY := centerX+k.position.X*9, centerY+k.position.Y*9
		for x := knotX - 4; x <= knotX+4; x++ {
			for y := knotY - 4; y <= knotY+4; y++ {
				img.Set(x, y, palette[i+1])
			}
		}
	}
	return img
}

func (d *Day08) Visualize() {
	var w, h int = 500, 500

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{250, 250, 110, 0xff},
		color.RGBA{207, 241, 116, 0xff},
		color.RGBA{166, 229, 125, 0xff},
		color.RGBA{129, 216, 134, 0xff},
		color.RGBA{93, 202, 143, 0xff},
		color.RGBA{60, 187, 150, 0xff},
		color.RGBA{28, 170, 153, 0xff},
		color.RGBA{1, 154, 151, 0xff},
		color.RGBA{13, 137, 146, 0xff},
		color.RGBA{33, 120, 136, 0xff},
	}
	var images []*image.Paletted

	r := getRope(10)
	for _, move := range d.moves[:300] {

		for i := 0; i < move.dst; i++ {
			r.head().move(move.dir)
			img := r.visualize(w, h, palette)
			images = append(images, img)
			for j := 1; j < len(r.knots); j++ {
				r.knots[j].update(r.knots[j-1].position.X, r.knots[j-1].position.Y)
				images = append(images, r.visualize(w, h, palette))
			}
		}
	}

	for i, img := range images {
		f, err := os.OpenFile(fmt.Sprintf("animation/rope%02d.png", i), os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		png.Encode(f, img)
	}
}
