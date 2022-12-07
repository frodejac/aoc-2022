package day06

import (
	"strconv"
	"strings"
)

type Day06 struct {
	root *directory
}

type file struct {
	parent *directory
	size   int
}

type directory struct {
	parent   *directory
	children map[string]*directory
	files    map[string]*file
	size     int
}

func (d *directory) recursiveCalcSizes() {
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, c := range d.children {
		c.recursiveCalcSizes()
		size += c.size
	}
	d.size = size
}

func parseInput(input string) *directory {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	root := &directory{parent: nil, children: map[string]*directory{}, files: map[string]*file{}}
	curdir := root
	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		switch lineParts[0] {
		case "$": // command
			switch lineParts[1] {
			case "cd":
				dir := lineParts[2]
				// change directory
				switch dir {
				case "..":
					curdir = curdir.parent
				case "/":
					curdir = root
				default:
					curdir = curdir.children[dir]
				}
			case "ls":
				// list directory, noop
				continue
			}
		case "dir": // directory
			{
				dirname := lineParts[1]
				dir := &directory{parent: curdir, children: map[string]*directory{}, files: map[string]*file{}}
				curdir.children[dirname] = dir
			}

		default: // file
			{
				fsize, _ := strconv.Atoi(lineParts[0])
				fname := lineParts[1]
				file := &file{parent: curdir, size: fsize}
				curdir.files[fname] = file
			}
		}
	}
	root.recursiveCalcSizes()
	return root
}

func sumSizesLessThanEqual(d *directory, size int) int {
	sum := 0
	if d.size <= size {
		sum += d.size
	}
	for _, c := range d.children {
		sum += sumSizesLessThanEqual(c, size)
	}
	return sum
}

func findSizeOfSmallestDirectoryToDelete(d *directory, freeRequired int, minSize *int) {
	if d.size >= freeRequired && d.size < *minSize {
		*minSize = d.size
	}
	for _, c := range d.children {
		findSizeOfSmallestDirectoryToDelete(c, freeRequired, minSize)
	}
}

func Solver(input []byte) *Day06 {
	return &Day06{root: parseInput(string(input))}
}

func (d *Day06) SolvePart1() string {
	return strconv.Itoa(sumSizesLessThanEqual(d.root, 100000))
}

func (d *Day06) SolvePart2() string {
	const availableSpace, necessarySpace int = 70000000, 30000000
	unusedSpace := availableSpace - d.root.size
	freeRequired := necessarySpace - unusedSpace

	dirsize := availableSpace
	findSizeOfSmallestDirectoryToDelete(d.root, freeRequired, &dirsize)

	return strconv.Itoa(dirsize)
}
