package main

import (
	"adventofcode-2020/day1"
	"adventofcode-2020/day10"
	"adventofcode-2020/day11"
	"adventofcode-2020/day12"
	"adventofcode-2020/day13"
	"adventofcode-2020/day2"
	"adventofcode-2020/day3"
	"adventofcode-2020/day4"
	"adventofcode-2020/day5"
	"adventofcode-2020/day6"
	"adventofcode-2020/day7"
	"adventofcode-2020/day8"
	"adventofcode-2020/day9"
	"flag"
	"fmt"
	"os"
)

func main() {
	dayNumber := flag.Int("day", 0, "Day of adventofcode: 1-31")
	partNumber := flag.Int("part", 0, "Part of the day: 1 or 2")

	flag.Parse()

	answer, err := getAnswerForPart(*dayNumber, *partNumber)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("The answer is %s\n", answer)
}

func getAnswerForPart(dayNumber int, partNumber int) (string, error) {
	getAnswer, err := getPart(dayNumber, partNumber)

	if err != nil {
		return "", err
	}

	return getAnswer(), nil
}

func getPart(dayNumber int, partNumber int) (func() string, error) {
	part1, part2, err := getDay(dayNumber)

	if err != nil {
		return nil, err
	}

	if partNumber == 1 {
		return part1, nil
	}

	if partNumber == 2 {
		return part2, nil
	}

	return nil, fmt.Errorf("%d is not a valid part of the day. Try 1 or 2.", dayNumber)
}

func getDay(dayNumber int) (func() string, func() string, error) {
	switch dayNumber {
	case 1:
		return day1.Part1, day1.Part2, nil
	case 2:
		return day2.Part1, day2.Part2, nil
	case 3:
		return day3.Part1, day3.Part2, nil
	case 4:
		return day4.Part1, day4.Part2, nil
	case 5:
		return day5.Part1, day5.Part2, nil
	case 6:
		return day6.Part1, day6.Part2, nil
	case 7:
		return day7.Part1, day7.Part2, nil
	case 8:
		return day8.Part1, day8.Part2, nil
	case 9:
		return day9.Part1, day9.Part2, nil
	case 10:
		return day10.Part1, day10.Part2, nil
	case 11:
		return day11.Part1, day11.Part2, nil
	case 12:
		return day12.Part1, day12.Part2, nil
	case 13:
		return day13.Part1, day13.Part2, nil
	}

	return nil, nil, fmt.Errorf("%d is not a valid day number. It's either not a December date (1-31) or I have no solution for that day yet.", dayNumber)
}
