package day6

import (
	"adventofcode-2020/input"
	"fmt"
)

const INPUT = "./day6/input.txt"

func Part1() string {
	entries := input.ReadStrings(INPUT)
	sum := 0
	groupFrom := 0
	groupTo := 0

	for i, entry := range entries {
		if entry != "" {
			groupTo = i
		}

		if entry == "" || i == len(entries)-1 {
			sum += getRepeatedAnswersCount(
				entries[groupFrom:groupTo+1],
				1,
			)
			groupFrom = i + 1
		}
	}

	return fmt.Sprint(sum)
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	sum := 0
	groupFrom := 0
	groupTo := 0

	for i, entry := range entries {
		if entry != "" {
			groupTo = i
		}

		if entry == "" || i == len(entries)-1 {
			groupSize := groupTo - groupFrom + 1

			sum += getRepeatedAnswersCount(
				entries[groupFrom:groupTo+1],
				groupSize,
			)

			groupFrom = i + 1
		}
	}

	return fmt.Sprint(sum)
}

func getRepeatedAnswersCount(declarations []string, minAnswersRequired int) int {
	repeatedAnswersCount := 0
	answersCount := make(map[rune]int)

	for _, declaration := range declarations {
		for _, answer := range declaration {
			answersCount[answer]++

			if answersCount[answer] == minAnswersRequired {
				repeatedAnswersCount++
			}
		}
	}

	return repeatedAnswersCount
}
