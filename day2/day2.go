package day2

import (
	"adventofcode-2020/input"
	"fmt"
	"strconv"
	"strings"
)

const INPUT = "./day2/input.txt"

func Part1() string {
	entries := input.ReadStrings(INPUT)
	count := 0

	for _, entry := range entries {
		min, max, symb, password := parseEntry(entry)
		if isValidBySledPolicy(min, max, symb, password) {
			count++
		}
	}

	return fmt.Sprint(count)
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	count := 0

	for _, entry := range entries {
		pos1, pos2, symb, password := parseEntry(entry)
		if isValidByTobogganPolicy(pos1, pos2, symb, password) {
			count++
		}
	}

	return fmt.Sprint(count)
}

func parseEntry(entry string) (int, int, rune, string) {
	parts := strings.Split(entry, ": ")
	rule := parts[0]
	password := parts[1]

	parts = strings.Split(rule, " ")
	fromto := parts[0]
	symb := []rune(parts[1])[0]

	parts = strings.Split(fromto, "-")
	from, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return from, to, symb, password
}

func isValidBySledPolicy(min int, max int, symb rune, password string) bool {
	count := 0

	for _, char := range password {
		if char != symb {
			continue
		}

		count++

		if count > max {
			return false
		}
	}

	return count >= min
}

func isValidByTobogganPolicy(pos1 int, pos2 int, symb rune, password string) bool {
	symbols := []rune(password)
	symb1 := symbols[pos1-1]
	symb2 := symbols[pos2-1]

	if symb1 == symb {
		return symb2 != symb
	}

	return symb2 == symb
}
