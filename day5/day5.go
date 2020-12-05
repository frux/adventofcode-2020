package day5

import (
	"adventofcode-2020/input"
	"fmt"
	"sort"
)

const INPUT = "./day5/input.txt"
const ROWS = 128
const COLS = 8

func Part1() string {
	entries := input.ReadStrings(INPUT)
	var maxSeatId int

	for _, entry := range entries {
		row, col := getSeatCoords(entry)
		seatId := getSeatId(row, col)

		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}

	return fmt.Sprint(maxSeatId)
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	seatIds := make(map[int]bool)
	var mySeatId int

	for _, entry := range entries {
		row, col := getSeatCoords(entry)
		seatId := getSeatId(row, col)
		seatIds[seatId] = true

		if hasTwoNeighbours(seatIds, seatId) {
			mySeatId -= seatId
			continue
		}

		if isPossibleLeftNeighbour(seatIds, seatId) {
			mySeatId += seatId + 1
		}

		if isPossibleRightNeighbour(seatIds, seatId) {
			mySeatId += seatId - 1
		}
	}

	return fmt.Sprint(mySeatId)
}

func Part2NotOptimal() string {
	entries := input.ReadStrings(INPUT)
	seatIds := make([]int, len(entries))

	for i, entry := range entries {
		row, col := getSeatCoords(entry)
		seatIds[i] = getSeatId(row, col)
	}

	sort.Ints(seatIds)

	for i, seatId := range seatIds {
		if seatIds[i+1]-seatIds[i] == 2 {
			return fmt.Sprint(seatId + 1)
		}
	}

	return ""
}

func getSeatCoords(seatPath string) (int, int) {
	rowPath := seatPath[:7]
	colPath := seatPath[7:]

	rowFrom := 0
	rowTo := ROWS - 1

	for _, dir := range rowPath {
		if dir == 'F' {
			rowFrom, rowTo = lowerHalf(rowFrom, rowTo)
			continue
		}

		if dir == 'B' {
			rowFrom, rowTo = upperHalf(rowFrom, rowTo)
		}
	}

	colFrom := 0
	colTo := COLS - 1

	for _, dir := range colPath {
		if dir == 'L' {
			colFrom, colTo = lowerHalf(colFrom, colTo)
			continue
		}

		if dir == 'R' {
			colFrom, colTo = upperHalf(colFrom, colTo)
		}
	}

	return rowFrom, colFrom
}

func isPossibleLeftNeighbour(seatIds map[int]bool, seatId int) bool {
	return !seatIds[seatId+1] && seatIds[seatId+2]
}

func isPossibleRightNeighbour(seatIds map[int]bool, seatId int) bool {
	return !seatIds[seatId-1] && seatIds[seatId-2]
}

func hasTwoNeighbours(seatIds map[int]bool, seatId int) bool {
	return seatIds[seatId-1] && seatIds[seatId+1]
}

func lowerHalf(from int, to int) (int, int) {
	return from, from + ((to - from) / 2)
}

func upperHalf(from int, to int) (int, int) {
	return to - ((to - from) / 2), to
}

func getSeatId(row int, col int) int {
	return row*8 + col
}
