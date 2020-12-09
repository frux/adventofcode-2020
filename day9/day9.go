package day9

import (
	"adventofcode-2020/input"
	"fmt"
)

const INPUT = "./day9/input.txt"
const PREAMBLE_SIZE = 25

func Part1() string {
	entries := input.ReadIntegers(INPUT)
	numbersIndex := make(map[int]bool)

	for i, entry := range entries {
		windowStart := i - PREAMBLE_SIZE

		if windowStart > 0 {
			numbersIndex[entries[windowStart-1]] = false
		}

		numbersIndex[entry] = true

		if windowStart < 0 {
			continue
		}

		isValid := false

		for j := i - PREAMBLE_SIZE; j < i; j++ {
			num1 := entries[j]
			num2 := entry - num1

			if num1 != num2 && numbersIndex[num2] {
				isValid = true
				break
			}
		}

		if !isValid {
			return fmt.Sprint(entry)
		}
	}

	return ""
}

func Part2() string {
	entries := input.ReadIntegers(INPUT)
	numbersIndex := make(map[int]bool)

	for i, entry := range entries {
		windowStart := i - PREAMBLE_SIZE

		if windowStart > 0 {
			numbersIndex[entries[windowStart-1]] = false
		}

		numbersIndex[entry] = true

		if windowStart < 0 {
			continue
		}

		isValid := false

		for j := i - PREAMBLE_SIZE; j < i; j++ {
			num1 := entries[j]
			num2 := entry - num1

			if num1 != num2 && numbersIndex[num2] {
				isValid = true
				break
			}
		}

		if isValid {
			continue
		}

		for j := 0; j < i; j++ {
			sum := entries[j]
			min := entries[j]
			max := entries[j]

			for k := j + 1; k < i; k++ {
				sum += entries[k]

				if sum > entry {
					break
				}

				if entries[k] < min {
					min = entries[k]
				}

				if entries[k] > max {
					max = entries[k]
				}

				if sum < entry {
					continue
				}

				return fmt.Sprint(min + max)
			}
		}
	}

	return ""
}
