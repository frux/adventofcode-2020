package day8

import (
	"adventofcode-2020/input"
	"fmt"
	"strconv"
	"strings"
)

const INPUT = "./day8/input.txt"
const ACC = "acc"
const JMP = "jmp"
const NOP = "nop"

func Part1() string {
	entries := input.ReadStrings(INPUT)
	comp := computer{}
	debugger := make(map[int]bool)

	for comp.Ptr < len(entries) && !debugger[comp.Ptr] {
		debugger[comp.Ptr] = true
		cmd, arg := parseCommand(entries[comp.Ptr])
		comp.Exec(cmd, arg)
	}

	return fmt.Sprint(comp.Acc)
}

func Part2() string {
	entries := input.ReadStrings(INPUT)

	for i, entry := range entries {
		cmd, _ := parseCommand(entry)

		if cmd == ACC {
			continue
		}

		if ok, result := tryFix(entries, i); ok {
			return fmt.Sprint(result)
		}
	}

	return ""
}

func tryFix(programm []string, corruptedAt int) (bool, int) {
	comp := computer{}
	debugger := make(map[int]bool)

	for comp.Ptr < len(programm) {
		if debugger[comp.Ptr] {
			return false, 0
		}
		debugger[comp.Ptr] = true

		cmd, arg := parseCommand(programm[comp.Ptr])
		if comp.Ptr == corruptedAt {
			if cmd == JMP {
				cmd = NOP
			} else {
				cmd = JMP
			}
		}

		comp.Exec(cmd, arg)
	}

	return true, comp.Acc
}

type computer struct {
	Ptr int
	Acc int
}

func (comp *computer) Exec(cmd string, arg int) {
	switch cmd {
	case ACC:
		comp.Acc += arg
		comp.Ptr++
	case JMP:
		comp.Ptr += arg
	case NOP:
		comp.Ptr++
	}
}

func parseCommand(cmdLine string) (string, int) {
	parts := strings.Split(cmdLine, " ")
	arg, _ := strconv.Atoi(parts[1])

	return parts[0], arg
}
