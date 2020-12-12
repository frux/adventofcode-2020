package day11

import (
	"adventofcode-2020/input"
	"fmt"
)

const INPUT = "./day11/input.txt"
const EMPTY = "L"
const OCCUPIED = "#"
const FLOOR = "."

func Part1() string {
	entries := input.ReadStrings(INPUT)
	seatMap := make([]string, len(entries))
	copy(seatMap, entries)

	for {
		hasChanges := false
		occupiedSeatsCount := 0
		nextSeatMap := make([]string, len(seatMap))
		copy(nextSeatMap, seatMap)

		for row := range seatMap {
			for col := range seatMap[row] {
				nextState := apllyRules1(seatMap, row, col)
				set(nextSeatMap, row, col, nextState)
				hasChanges = hasChanges || seatMap[row][col] != nextSeatMap[row][col]

				if get(nextSeatMap, row, col) == OCCUPIED {
					occupiedSeatsCount++
				}
			}
		}

		if !hasChanges {
			return fmt.Sprint(occupiedSeatsCount)
		}

		seatMap = nextSeatMap
	}
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	seatMap := make([]string, len(entries))
	copy(seatMap, entries)

	for {
		hasChanges := false
		occupiedSeatsCount := 0
		nextSeatMap := make([]string, len(seatMap))
		copy(nextSeatMap, seatMap)

		for row := range seatMap {
			for col := range seatMap[row] {
				nextState := apllyRules2(seatMap, row, col)
				set(nextSeatMap, row, col, nextState)
				hasChanges = hasChanges || seatMap[row][col] != nextSeatMap[row][col]

				if get(nextSeatMap, row, col) == OCCUPIED {
					occupiedSeatsCount++
				}
			}
		}

		if !hasChanges {
			return fmt.Sprint(occupiedSeatsCount)
		}

		seatMap = nextSeatMap
	}
}

func apllyRules1(seatMap []string, row int, col int) string {
	state := get(seatMap, row, col)

	if state == EMPTY && !hasAdjacentOccupiedSeats(seatMap, row, col, 1) {
		return OCCUPIED
	}

	if state == OCCUPIED && hasAdjacentOccupiedSeats(seatMap, row, col, 4) {
		return EMPTY
	}

	return state
}

func apllyRules2(seatMap []string, row int, col int) string {
	state := get(seatMap, row, col)

	if state == EMPTY {
		if !canSeeOccupiedSeatsAround(seatMap, row, col, 1) {
			return OCCUPIED
		}
	}

	if state == OCCUPIED && canSeeOccupiedSeatsAround(seatMap, row, col, 5) {
		return EMPTY
	}

	return state
}

func hasAdjacentOccupiedSeats(seatMap []string, row int, col int, required int) bool {
	count := 0
	seats := [][2]int{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col + 1},
		{row, col - 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}

	for _, coords := range seats {
		if get(seatMap, coords[0], coords[1]) == OCCUPIED {
			if count++; count == required {
				return true
			}
		}
	}

	return false
}

func canSeeOccupiedSeatsAround(seatMap []string, row int, col int, required int) bool {
	count := 0
	directions := []func([]string, int, int) [][2]int{
		top,
		topRight,
		right,
		bottomRight,
		bottom,
		bottomLeft,
		left,
		topLeft,
	}

	for _, direction := range directions {
		if canSeeOccupiedSeat(seatMap, direction(seatMap, row, col)) {
			if count++; count == required {
				return true
			}
		}
	}

	return false
}

func canSeeOccupiedSeat(seatMap []string, direction [][2]int) bool {
	for _, coords := range direction {
		state := get(seatMap, coords[0], coords[1])

		if state == FLOOR {
			continue
		}

		return state == OCCUPIED
	}

	return false
}

func top(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; row-i >= 0; i++ {
		direction = append(direction, [2]int{row - i, col})
	}
	return direction
}

func bottom(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; row+i < len(seatMap); i++ {
		direction = append(direction, [2]int{row + i, col})
	}
	return direction
}

func right(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; col+i < len(seatMap[row]); i++ {
		direction = append(direction, [2]int{row, col + i})
	}
	return direction
}

func left(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; col-i >= 0; i++ {
		direction = append(direction, [2]int{row, col - i})
	}
	return direction
}

func topLeft(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; row-i >= 0 && col-i >= 0; i++ {
		direction = append(direction, [2]int{row - i, col - i})
	}
	return direction
}

func topRight(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; row-i >= 0 && col+i < len(seatMap[row-i]); i++ {
		direction = append(direction, [2]int{row - i, col + i})
	}
	return direction
}

func bottomRight(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; row+i < len(seatMap) && col+i < len(seatMap[row+i]); i++ {
		direction = append(direction, [2]int{row + i, col + i})
	}
	return direction
}

func bottomLeft(seatMap []string, row int, col int) [][2]int {
	direction := [][2]int{}
	for i := 1; row+i < len(seatMap) && col-i >= 0; i++ {
		direction = append(direction, [2]int{row + i, col - i})
	}
	return direction
}

func get(seatMap []string, row int, col int) string {
	if row < 0 || row >= len(seatMap) || col < 0 || col >= len(seatMap[row]) {
		return ""
	}

	return string(seatMap[row][col])
}

func set(seatMap []string, row int, col int, nextState string) {
	seatMap[row] = seatMap[row][:col] + nextState + seatMap[row][col+1:]
}
