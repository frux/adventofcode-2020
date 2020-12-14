package day14

import (
	"adventofcode-2020/input"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const INPUT = "./day14/input.txt"
const MASK = "mask"
const MEM = "mem"

func Part1() string {
	entries := input.ReadStrings(INPUT)
	mask := ""
	memory := newMemory()

	for _, entry := range entries {
		cmd, arg1, arg2 := parseEntry(entry)

		switch cmd {
		case MASK:
			mask = arg1
		case MEM:
			addr, _ := strconv.Atoi(arg1)
			val, _ := strconv.Atoi(arg2)
			val = applyMask1(val, mask)
			memory.Write(addr, val)
		}
	}

	sum := 0

	for _, val := range memory.Values {
		sum += val
	}

	return fmt.Sprint(sum)
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	mask := ""
	memory := newMemory()

	for _, entry := range entries {
		cmd, arg1, arg2 := parseEntry(entry)

		switch cmd {
		case MASK:
			mask = arg1
		case MEM:
			addr, _ := strconv.Atoi(arg1)
			val, _ := strconv.Atoi(arg2)
			addresses := applyMask2(addr, mask)
			for _, addr := range addresses {
				memory.Write(addr, val)
			}
		}
	}

	sum := 0

	for _, val := range memory.Values {
		sum += val
	}

	return fmt.Sprint(sum)
}

type Memory struct {
	Values    []int
	Addresses map[int]int
}

func newMemory() *Memory {
	return &Memory{
		Values:    []int{},
		Addresses: make(map[int]int),
	}
}

func (m *Memory) Read(ptr int) int {
	return m.Values[m.Addresses[ptr]]
}

func (m *Memory) Write(ptr int, val int) {
	addr, isMemmoryAllocated := m.Addresses[ptr]

	if isMemmoryAllocated {
		m.Values[addr] = val
		return
	}

	m.Values = append(m.Values, val)
	m.Addresses[ptr] = len(m.Values) - 1
}

func parseEntry(entry string) (string, string, string) {
	if entry[:4] == MASK {
		return MASK, entry[7:], ""
	}

	re, _ := regexp.Compile(`^mem\[(\d+)\] = (\d+)$`)
	parts := re.FindStringSubmatch(entry)

	return MEM, parts[1], parts[2]
}

func applyMask1(val int, mask string) int {
	for i := 0; i < len(mask); i++ {
		bit := mask[len(mask)-1-i]

		switch bit {
		case '1':
			val = bitwiseSet(val, i, 1)
		case '0':
			val = bitwiseSet(val, i, 0)
		}
	}

	return val
}

func applyMask2(val int, mask string) []int {
	floatingBits := []int{}

	for i := 0; i < len(mask); i++ {
		bit := mask[len(mask)-1-i]

		switch bit {
		case '1':
			val = bitwiseSet(val, i, 1)
		case 'X':
			floatingBits = append(floatingBits, i)
		}
	}

	valuesLength := int(math.Pow(2, float64(len(floatingBits))))
	values := make([]int, valuesLength)

	for i := 0; i < valuesLength; i++ {
		for j, index := range floatingBits {
			bit := bitwiseGet(i, j)
			val = bitwiseSet(val, index, bit)
		}

		values[i] = val
	}

	return values
}

func bitwiseSet(val int, index int, bit int) int {
	mod := int(math.Pow(2, float64(index)))

	if bit == 0 {
		return val &^ mod
	}

	return val | mod
}

func bitwiseGet(val int, index int) int {
	if val&int(math.Pow(2, float64(index))) == 0 {
		return 0
	}

	return 1
}
