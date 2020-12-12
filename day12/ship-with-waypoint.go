package day12

import (
	"math"
)

type ShipWithWayPoint struct {
	Coords   [2]int
	Waypoint struct {
		Coords [2]int
	}
}

func (s *ShipWithWayPoint) GoNorth(distance int) {
	s.Waypoint.Coords[1] += distance
}
func (s *ShipWithWayPoint) GoSouth(distance int) {
	s.Waypoint.Coords[1] -= distance
}
func (s *ShipWithWayPoint) GoEast(distance int) {
	s.Waypoint.Coords[0] += distance
}
func (s *ShipWithWayPoint) GoWest(distance int) {
	s.Waypoint.Coords[0] -= distance
}

func (s *ShipWithWayPoint) Turn(angle int) {
	angleRad := rad(angle)
	cos := int(math.Cos(angleRad))
	sin := int(math.Sin(angleRad))

	nextX := s.Waypoint.Coords[0]*cos + s.Waypoint.Coords[1]*sin
	nextY := s.Waypoint.Coords[1]*cos - s.Waypoint.Coords[0]*sin

	s.Waypoint.Coords[0] = nextX
	s.Waypoint.Coords[1] = nextY
}

func (s *ShipWithWayPoint) GoForward(times int) {
	s.Coords[0] += s.Waypoint.Coords[0] * times
	s.Coords[1] += s.Waypoint.Coords[1] * times
}

func (s *ShipWithWayPoint) GetDistance() int {
	return int(math.Abs(float64(s.Coords[0])) + math.Abs(float64(s.Coords[1])))
}

func rad(deg int) float64 {
	return float64(deg) * (math.Pi / 180.0)
}
