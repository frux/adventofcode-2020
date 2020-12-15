package day15

import (
	"adventofcode-2020/input"
	"fmt"
	"strconv"
	"strings"
)

const INPUT = "./day15/input.txt"

func Part1() string {
	return fmt.Sprint(playGame(2020))
}

func Part2() string {
	return fmt.Sprint(playGame(30000000))
}

func playGame(turnsNumber int) int {
	entries := input.ReadStrings(INPUT)
	startingNumbers := parseStartingNumbers(entries[0])
	game := NewGame(startingNumbers)

	for game.TurnNumber < turnsNumber {
		prevNum := game.LastSpokenNumber

		if game.GetSpokenCount(prevNum) == 1 {
			game.Speak(0)
			continue
		}

		recentTimesSpoken := game.GetRecentTimesSpoken(prevNum)
		game.Speak(recentTimesSpoken[1] - recentTimesSpoken[0])
	}

	return game.LastSpokenNumber
}

type Game struct {
	TurnNumber       int
	LastSpokenNumber int
	numbersCounter   map[int]int
	history          map[int][2]int
}

func NewGame(startingNumbers []int) *Game {
	game := Game{
		numbersCounter: make(map[int]int),
		history:        make(map[int][2]int),
	}

	for _, num := range startingNumbers {
		game.Speak(num)
	}

	return &game
}

func (g *Game) GetSpokenCount(num int) int {
	return g.numbersCounter[num]
}

func (g *Game) GetRecentTimesSpoken(num int) [2]int {
	return g.history[num]
}

func (g *Game) Speak(num int) {
	g.TurnNumber++
	g.LastSpokenNumber = num
	g.history[num] = [2]int{
		g.GetRecentTimesSpoken(num)[1],
		g.TurnNumber,
	}
	g.numbersCounter[num]++
}

func parseStartingNumbers(line string) []int {
	startingNumbers := []int{}

	for _, num := range strings.Split(line, ",") {
		parsedNum, _ := strconv.Atoi(num)
		startingNumbers = append(startingNumbers, parsedNum)
	}

	return startingNumbers
}
