package day12

import (
	"adventofcode-2020/input"
	"fmt"
	"strconv"
)

const INPUT = "./day12/input.txt"

const GO_NORTH = "N"
const GO_EAST = "E"
const GO_SOUTH = "S"
const GO_WEST = "W"
const GO_FORWARD = "F"
const TURN_RIGHT = "R"
const TURN_LEFT = "L"

func Part1() string {
	entries := input.ReadStrings(INPUT)
	ship := Ship{}
	ship.Angle = 90

	for _, entry := range entries {
		cmd := entry[:1]
		arg, _ := strconv.Atoi(entry[1:])

		switch cmd {
		case GO_NORTH:
			ship.GoNorth(arg)
		case GO_EAST:
			ship.GoEast(arg)
		case GO_SOUTH:
			ship.GoSouth(arg)
		case GO_WEST:
			ship.GoWest(arg)
		case GO_FORWARD:
			ship.GoForward(arg)
		case TURN_RIGHT:
			ship.Turn(arg)
		case TURN_LEFT:
			ship.Turn(-1 * arg)
		}
	}

	return fmt.Sprint(ship.GetDistance())
}

func Part2() string {
	entries := input.ReadStrings(INPUT)
	ship := ShipWithWayPoint{}
	ship.Waypoint.Coords[0] = 10
	ship.Waypoint.Coords[1] = 1

	for _, entry := range entries {
		cmd := entry[:1]
		arg, _ := strconv.Atoi(entry[1:])

		switch cmd {
		case GO_NORTH:
			ship.GoNorth(arg)
		case GO_EAST:
			ship.GoEast(arg)
		case GO_SOUTH:
			ship.GoSouth(arg)
		case GO_WEST:
			ship.GoWest(arg)
		case GO_FORWARD:
			ship.GoForward(arg)
		case TURN_RIGHT:
			ship.Turn(arg)
		case TURN_LEFT:
			ship.Turn(-1 * arg)
		}
	}

	return fmt.Sprint(ship.GetDistance())
}
