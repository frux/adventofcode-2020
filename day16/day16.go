package day16

import (
	"adventofcode-2020/input"
	"fmt"
	"strconv"
	"strings"
)

const INPUT = "./day16/input.txt"

func Part1() string {
	entries := input.ReadStrings(INPUT)
	fields := []*Field{}

	for _, entry := range entries {
		if entry == "" {
			break
		}

		fields = append(fields, parseField(entry))
	}

	myTicketIndex := len(fields) + 2
	nearbyTicketsIndex := myTicketIndex + 3
	errorRate := 0

	for i := nearbyTicketsIndex; i < len(entries); i++ {
		ticket := parseTicket(entries[i])

		for _, value := range ticket {
			if !isValidForAnyField(fields, value) {
				errorRate += value
			}
		}
	}

	return fmt.Sprint(errorRate)
}

func Part2() string {
	return "skipped"
}

func isValidForAnyField(fields []*Field, value int) bool {
	for _, field := range fields {
		if isValidForField(field, value) {
			return true
		}
	}

	return false
}

func isValidForField(field *Field, value int) bool {
	for _, limits := range field.Ranges {
		if value >= limits[0] && value <= limits[1] {
			return true
		}
	}

	return false
}

func parseTicket(line string) []int {
	parts := strings.Split(line, ",")
	ticket := make([]int, len(parts))

	for i, part := range parts {
		value, _ := strconv.Atoi(part)
		ticket[i] = value
	}

	return ticket
}

func parseField(line string) *Field {
	parts := strings.Split(line, ": ")
	return &Field{
		Name:   parts[0],
		Ranges: parseRanges(parts[1]),
	}
}

func parseRanges(line string) [][2]int {
	parts := strings.Split(line, " or ")
	ranges := make([][2]int, len(parts))
	for i, part := range parts {
		numbers := strings.Split(part, "-")
		min, _ := strconv.Atoi(numbers[0])
		max, _ := strconv.Atoi(numbers[1])
		ranges[i] = [2]int{min, max}
	}

	return ranges
}

type Field struct {
	Name   string
	Ranges [][2]int
}
