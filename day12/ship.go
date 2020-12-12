package day12

import (
	"math"
)

type Ship struct {
	Angle  int
	Coords [2]int
}

func (s *Ship) GoNorth(distance int) {
	s.Coords[1] += distance
}
func (s *Ship) GoSouth(distance int) {
	s.Coords[1] -= distance
}
func (s *Ship) GoEast(distance int) {
	s.Coords[0] += distance
}
func (s *Ship) GoWest(distance int) {
	s.Coords[0] -= distance
}

func (s *Ship) Turn(angle int) {
	s.Angle = (s.Angle + angle) % 360

	if s.Angle < 0 {
		s.Angle += 360
	}
}

func (s *Ship) GoForward(distance int) {
	switch s.Angle {
	case 0:
		s.GoNorth(distance)
	case 90:
		s.GoEast(distance)
	case 180:
		s.GoSouth(distance)
	case 270:
		s.GoWest(distance)
	}
}

func (s *Ship) GetDistance() int {
	return int(math.Abs(float64(s.Coords[0])) + math.Abs(float64(s.Coords[1])))
}
