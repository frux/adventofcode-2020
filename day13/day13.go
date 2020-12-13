package day13

import (
	"adventofcode-2020/input"
	"fmt"
	"strconv"
	"strings"
)

const INPUT = "./day13/input.txt"

func Part1() string {
	entries := input.ReadStrings(INPUT)
	iArriveAt, buses := parse(entries)
	myBusId := buses[0][0]
	timeToWaitMyBus := getTimeToWait(iArriveAt, myBusId)

	for i := 1; i < len(buses); i++ {
		busId := buses[i][0]

		if busId == -1 {
			continue
		}

		timeToWaitThisBus := getTimeToWait(iArriveAt, busId)

		if timeToWaitThisBus < timeToWaitMyBus {
			timeToWaitMyBus = timeToWaitThisBus
			myBusId = busId
		}
	}

	return fmt.Sprint(myBusId * timeToWaitMyBus)
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	_, buses := parse(entries)
	matchAt := buses[0][0]
	step := buses[0][0]

	for i := 0; i < len(buses)-1; i++ {
		for !checkTime(buses[i:i+2], matchAt) {
			matchAt += step
		}
		step = step * buses[i+1][0]
	}

	return fmt.Sprint(matchAt)
}

func parse(entries []string) (int, [][2]int) {
	iArriveAt, _ := strconv.Atoi(entries[0])
	buses := [][2]int{}

	for i, bus := range strings.Split(entries[1], ",") {
		if bus == "x" {
			continue
		}

		busId, _ := strconv.Atoi(bus)

		buses = append(buses, [2]int{busId, i})
	}

	return iArriveAt, buses
}

func getTimeToWait(iArriveAt int, busId int) int {
	return (busId - iArriveAt%busId) % busId
}

func checkTime(buses [][2]int, time int) bool {
	for _, bus := range buses {
		if (time+bus[1])%bus[0] != 0 {
			return false
		}
	}

	return true
}
