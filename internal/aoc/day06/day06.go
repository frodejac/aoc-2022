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

func (d *directory) calcSize() {
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, c := range d.children {
		c.calcSize()
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
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "$":
			cmd := parts[1]
			switch cmd {
			case "cd":
				dir := parts[2]
				switch dir {
				case "..":
					curdir = curdir.parent
				case "/":
					curdir = root
				default:
					curdir = curdir.children[dir]
				}
			case "ls":
				continue
			}
		case "dir":
			dirname := parts[1]
			dir := &directory{parent: curdir, children: map[string]*directory{}, files: map[string]*file{}}
			curdir.children[dirname] = dir
		default:
			fsize, _ := strconv.Atoi(parts[0])
			fname := parts[1]
			file := &file{parent: curdir, size: fsize}
			curdir.files[fname] = file
		}
	}
	root.calcSize()
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
