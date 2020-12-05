package main_test

import (
	"adventofcode-2020/day1"
	"adventofcode-2020/day2"
	"adventofcode-2020/day3"
	"adventofcode-2020/day4"
	"adventofcode-2020/day5"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func Test(t *testing.T) {
	assertDays(t, []tDay{
		{
			{day1.Part1, "1016619"},
			{day1.Part2, "218767230"},
		},
		{
			{day2.Part1, "460"},
			{day2.Part2, "251"},
		},
		{
			{day3.Part1, "247"},
			{day3.Part2, "2983070376"},
		},
		{
			{day4.Part1, "222"},
			{day4.Part2, "140"},
		},
		{
			{day5.Part1, "987"},
			{day5.Part2, "603"},
			{day5.Part2NotOptimal, "603"},
		},
	})
}

type tDay = []tPart
type tPart struct {
	Fn       func() string
	Expected string
}

func assertDays(t *testing.T, days []tDay) {
	for _, day := range days {
		for _, part := range day {
			assert(t, fmt.Sprint(getFuncName(part.Fn)), part.Fn(), part.Expected)
		}
	}
}

func assert(t *testing.T, message string, actual string, expected string) {
	if actual != expected {
		t.Fatalf("%s – Actual: %s, expected: %s", message, actual, expected)
	}
}

func getFuncName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
