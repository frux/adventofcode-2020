package main_test

import (
	"adventofcode-2020/day1"
	"adventofcode-2020/day10"
	"adventofcode-2020/day11"
	"adventofcode-2020/day12"
	"adventofcode-2020/day13"
	"adventofcode-2020/day14"
	"adventofcode-2020/day15"
	"adventofcode-2020/day2"
	"adventofcode-2020/day3"
	"adventofcode-2020/day4"
	"adventofcode-2020/day5"
	"adventofcode-2020/day6"
	"adventofcode-2020/day7"
	"adventofcode-2020/day8"
	"adventofcode-2020/day9"
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
		{
			{day6.Part1, "6161"},
			{day6.Part2, "2971"},
		},
		{
			{day7.Part1, "124"},
			{day7.Part2, "34862"},
		},
		{
			{day8.Part1, "1723"},
			{day8.Part2, "846"},
		},
		{
			{day9.Part1, "1309761972"},
			{day9.Part2, "177989832"},
		},
		{
			{day10.Part1, "2368"},
			{day10.Part2, "1727094849536"},
		},
		{
			{day11.Part1, "2247"},
			{day11.Part2, "2011"},
		},
		{
			{day12.Part1, "1601"},
			{day12.Part2, "13340"},
		},
		{
			{day13.Part1, "2305"},
			{day13.Part2, "552612234243498"},
		},
		{
			{day14.Part1, "11926135976176"},
			{day14.Part2, "4330547254348"},
		},
		{
			{day15.Part1, "447"},
			{day15.Part2, "11721679"},
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
