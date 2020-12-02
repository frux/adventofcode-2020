package main

import (
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
	}

	return "", "", fmt.Errorf("Can't find day with number %d", dayNumber)
}
