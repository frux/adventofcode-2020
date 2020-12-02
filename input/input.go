package input

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadStrings(path string) []string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(data)

	return strings.Split(text, "\n")
}

func ReadIntegers(path string) []int {
	lines := ReadStrings(path)
	numbers := make([]int, len(lines))

	for i, line := range lines {
		number, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		numbers[i] = number
	}

	return numbers
}
