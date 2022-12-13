package day12

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/frodejac/aoc-2022/pkg/datastructures/stack"
	"github.com/frodejac/aoc-2022/pkg/math"
)

type packet struct {
	data []interface{}
}

type pair struct {
	left  packet
	right packet
}

func parseLine(line string) packet {
	s := stack.New([]interface{}{})
	var numBuf string = ""
	var curr []interface{} = make([]interface{}, 0)
	for _, c := range line {
		switch c {
		case '[':
			s.Push(curr)
			curr = make([]interface{}, 0)
		case ']':
			if numBuf != "" {
				i, _ := strconv.Atoi(numBuf)
				curr = append(curr, i)
				numBuf = ""
			}
			tmp := s.Pop().([]interface{})
			tmp = append(tmp, curr)
			curr = tmp
		case ',':
			if numBuf != "" {
				i, _ := strconv.Atoi(numBuf)
				curr = append(curr, i)
				numBuf = ""
			}
		default:
			numBuf += string(c)
		}
	}
	return packet{data: curr[0].([]interface{})}
}

func parsePair(rawPair string) pair {
	parts := strings.Split(rawPair, "\n")
	return pair{
		left:  parseLine(parts[0]),
		right: parseLine(parts[1]),
	}
}

func parsePairs(input string) []pair {
	input = strings.TrimSpace(input)
	pairsRaw := strings.Split(input, "\n\n")
	pairs := make([]pair, len(pairsRaw))
	for i, p := range pairsRaw {
		pairs[i] = parsePair(p)
	}
	return pairs
}

func parsePackets(input string) []packet {
	input = strings.Replace(input, "\n\n", "\n", -1)
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	packets := make([]packet, len(lines))
	for i, line := range lines {
		packets[i] = parseLine(line)
	}
	return packets

}

func compareInts(left, right int) (bool, bool) {
	if left > right {
		return false, true
	}
	if left < right {
		return true, true
	}
	return false, false
}

func compareLists(left, right []interface{}) (bool, bool) {
	size := math.Min(len(left), len(right))
	for i := 0; i < size; i++ {
		ordered, done := checkOrder(left[i], right[i])
		if done {
			return ordered, true
		}
	}
	if len(left) > len(right) {
		return false, true
	} else if len(left) < len(right) {
		return true, true
	}
	return false, false
}

func checkOrder(left, right interface{}) (bool, bool) {
	lval, lok := left.(int)
	rval, rok := right.(int)

	if lok && rok {
		return compareInts(lval, rval)
	} else if lok && !rok {
		return checkOrder([]interface{}{left}, right)
	} else if !lok && rok {
		return checkOrder(left, []interface{}{right})
	} else {
		return compareLists(left.([]interface{}), right.([]interface{}))
	}
}

func (p *pair) isOrdered() bool {
	ordered, _ := checkOrder(p.left.data, p.right.data)
	return ordered
}

type Day12 struct {
	input string
}

func Solver(input []byte) *Day12 {
	return &Day12{input: string(input)}
}

func (d *Day12) SolvePart1() string {
	pairs := parsePairs(d.input)
	sum := 0
	for i, pair := range pairs {
		if pair.isOrdered() {
			sum += (i + 1)
		}
	}
	return strconv.Itoa(sum)
}

func (d *Day12) SolvePart2() string {
	packets := parsePackets(d.input + "[[2]]\n[[6]]")
	sort.Slice(packets[:], func(i, j int) bool {
		ordered, _ := checkOrder(packets[i].data, packets[j].data)
		return ordered
	})

	var idx1, idx2 int
	for i := 0; i < len(packets); i++ {
		switch fmt.Sprint(packets[i].data) {
		case "[[2]]":
			idx1 = i + 1
		case "[[6]]":
			idx2 = i + 1
		default:
			continue
		}
	}
	return strconv.Itoa(idx1 * idx2)
}
