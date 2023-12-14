package lib

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("[X=%d, Y=%d]", p.X, p.Y)
}

func (p *Point) ManhattanDistance(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}
