package day10

import (
	"sort"
	"strconv"
	"strings"

	"github.com/frodejac/aoc-2022/pkg/datastructures/stack"
)

type monkey struct {
	inventory  *stack.Stack[int]
	operation  func(int) int
	mod        int
	dstTrue    int
	dstFalse   int
	opsCounter *int
}

func parseInventory(inventoryRaw string) *stack.Stack[int] {
	inventory := stack.New([]int{})
	itemsRaw := strings.Split(inventoryRaw, ": ")[1]
	items := strings.Split(itemsRaw, ", ")
	for _, item := range items {
		worryLevel, _ := strconv.Atoi(item)
		inventory.Push(worryLevel)
	}
	return inventory
}

func parseOperation(operationRaw string) func(int) int {
	operationRaw = strings.Split(operationRaw, "new = ")[1]
	parts := strings.Split(operationRaw, " ")

	operation := func(old int) int {
		var op2 int
		switch parts[2] {
		case "old":
			op2 = old
		default:
			op2, _ = strconv.Atoi(parts[2])
		}
		switch parts[1] {
		case "+":
			return old + op2
		case "*":
			return old * op2
		default:
			panic("Unknown operation")
		}
	}
	return operation
}

func parseMod(modRaw string) int {
	modRaw = strings.TrimSpace(modRaw)
	parts := strings.Split(modRaw, " ")
	mod, _ := strconv.Atoi(parts[3])
	return mod
}

func parseDestination(destinationRaw string) int {
	destinationRaw = strings.TrimSpace(destinationRaw)
	parts := strings.Split(destinationRaw, " ")
	dst, _ := strconv.Atoi(parts[5])
	return dst
}

func parseMonkey(monkeyRaw string) monkey {
	monkey := monkey{}
	lines := strings.Split(monkeyRaw, "\n")
	monkey.inventory = parseInventory(lines[1])
	monkey.operation = parseOperation(lines[2])
	monkey.mod = parseMod(lines[3])
	monkey.dstTrue = parseDestination(lines[4])
	monkey.dstFalse = parseDestination(lines[5])
	monkey.opsCounter = new(int)
	return monkey
}

func parseInput(input string) []monkey {
	monkeysRaw := strings.Split(input, "\n\n")
	monkeys := []monkey{}
	for _, monkeyRaw := range monkeysRaw {
		monkeys = append(monkeys, parseMonkey(monkeyRaw))
	}
	return monkeys
}

type Day10 struct {
	input []byte
}

func Solver(input []byte) *Day10 {
	return &Day10{input: input}
}

func (d *Day10) SolvePart1() string {
	monkeys := parseInput(string(d.input))

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			itemCount := monkey.inventory.Size()
			for i := 0; i < itemCount; i++ {
				*monkey.opsCounter++
				item := monkey.inventory.Pop()
				newItem := monkey.operation(item)
				newItem = newItem / 3
				if (newItem % monkey.mod) == 0 {
					monkeys[monkey.dstTrue].inventory.Push(newItem)
				} else {
					monkeys[monkey.dstFalse].inventory.Push(newItem)
				}
			}
		}
	}
	opsCounters := []int{}
	for _, monkey := range monkeys {
		opsCounters = append(opsCounters, *monkey.opsCounter)
	}
	sort.Slice(opsCounters[:], func(i, j int) bool {
		return opsCounters[i] > opsCounters[j]
	})
	monkeyBusiness := opsCounters[0] * opsCounters[1]
	return strconv.Itoa(monkeyBusiness)
}

func (d *Day10) SolvePart2() string {
	monkeys := parseInput(string(d.input))

	commonDivisor := 1
	for _, monkey := range monkeys {
		commonDivisor *= monkey.mod
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			itemCount := monkey.inventory.Size()
			for i := 0; i < itemCount; i++ {
				*monkey.opsCounter++
				item := monkey.inventory.Pop()
				newItem := monkey.operation(item)
				newItem = newItem % commonDivisor
				if (newItem % monkey.mod) == 0 {
					monkeys[monkey.dstTrue].inventory.Push(newItem)
				} else {
					monkeys[monkey.dstFalse].inventory.Push(newItem)
				}
			}
		}
	}
	opsCounters := []int{}
	for _, monkey := range monkeys {
		opsCounters = append(opsCounters, *monkey.opsCounter)
	}
	sort.Slice(opsCounters[:], func(i, j int) bool {
		return opsCounters[i] > opsCounters[j]
	})
	monkeyBusiness := opsCounters[0] * opsCounters[1]
	return strconv.Itoa(monkeyBusiness)
}
