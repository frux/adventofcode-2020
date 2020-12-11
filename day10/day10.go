package day10

import (
	"adventofcode-2020/input"
	"fmt"
	"sort"
)

const INPUT = "./day10/input.txt"

func Part1() string {
	entries := input.ReadIntegers(INPUT)
	sort.Ints(entries)
	threeJoltageCount := 0
	oneJoltageCount := 0
	device := entries[len(entries)-1] + 3
	adapters := append(entries, device)

	for i := 0; i < len(adapters); i++ {
		diff := 0
		if i == 0 {
			diff = adapters[i]
		} else {
			diff = adapters[i] - adapters[i-1]
		}

		if diff == 3 {
			threeJoltageCount++
		}

		if diff == 1 {
			oneJoltageCount++
		}
	}

	return fmt.Sprint(oneJoltageCount * threeJoltageCount)
}

func Part2() string {
	entries := input.ReadIntegers(INPUT)
	entries = append(entries, 0)
	sort.Ints(entries)
	device := entries[len(entries)-1] + 3
	entries = append(entries, device)
	variations := make([]int, len(entries))
	variations[0] = 1

	for i := 1; i < len(entries); i++ {
		if isSuitable(entries, i-3, i) {
			variations[i] += variations[i-3]
		}

		if isSuitable(entries, i-2, i) {
			variations[i] += variations[i-2]
		}

		if isSuitable(entries, i-1, i) {
			variations[i] += variations[i-1]
		}
	}

	return fmt.Sprint(variations[len(variations)-1])
}

func isSuitable(adapters []int, i1 int, i2 int) bool {
	if i1 < 0 || i2 >= len(adapters) {
		return false
	}

	return adapters[i2]-adapters[i1] <= 3
}
