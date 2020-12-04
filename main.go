package main

import (
	"adventofcode-2020/day1"
	"adventofcode-2020/day2"
	"adventofcode-2020/day3"
	"adventofcode-2020/day4"
	"flag"
	"fmt"
	"os"
)

func main() {
	dayNumber := flag.Int("day", 0, "Day of adventofcode")
	flag.Parse()

	answer1, answer2, err := getAnswers(*dayNumber)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Answers are %s and %s\n", answer1, answer2)
}

func getAnswers(dayNumber int) (string, string, error) {
	switch dayNumber {
	case 1:
		return day1.Part1(), day1.Part2(), nil
	case 2:
		return day2.Part1(), day2.Part2(), nil
	case 3:
		return day3.Part1(), day3.Part2(), nil
	case 4:
		return day4.Part1(), day4.Part2(), nil
	}

	return "", "", fmt.Errorf("Can't find day with number %d", dayNumber)
}
