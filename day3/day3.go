package day3

import (
	"adventofcode-2020/input"
	"fmt"
	"math"
)

const INPUT = "./day3/input.txt"
const TREE = "#"

func Part1() string {
	entries := input.ReadStrings(INPUT)

	return fmt.Sprint(getTreesCount(entries, 3, 1))
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	result := getTreesCount(entries, 1, 1) *
		getTreesCount(entries, 3, 1) *
		getTreesCount(entries, 5, 1) *
		getTreesCount(entries, 7, 1) *
		getTreesCount(entries, 1, 2)

	return fmt.Sprint(result)
}

func getTreesCount(entries []string, right int, down int) int {
	mapWidth := len(entries[0])
	mapHeight := len(entries)
	x := 0
	y := 0
	maxY := mapHeight - 1
	treesCounter := 0

	for y < maxY {
		x = (x + right) % mapWidth
		y = int(math.Min(float64(y+down), float64(maxY)))

		if string(entries[y][x]) == TREE {
			treesCounter++
		}
	}

	return treesCounter
}
