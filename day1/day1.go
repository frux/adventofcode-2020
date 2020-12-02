package day1

import (
	"adventofcode-2020/input"
	"fmt"
)

const INPUT = "./day1/input.txt"

func Part1() string {
	entries := input.ReadIntegers(INPUT)

	for i := range entries {
		for j := i; j < len(entries); j++ {
			if entries[i]+entries[j] == 2020 {
				return fmt.Sprint(entries[i] * entries[j])
			}
		}
	}

	return ""
}

func Part2() string {
	entries := input.ReadIntegers(INPUT)

	for i := range entries {
		for j := i; j < len(entries); j++ {
			for k := i; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == 2020 {
					return fmt.Sprint(entries[i] * entries[j] * entries[k])
				}
			}
		}
	}

	return ""
}
