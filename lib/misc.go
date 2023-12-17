package lib

type OrthogonalDirection int8

const (
	UP    = OrthogonalDirection(0)
	RIGHT = OrthogonalDirection(1)
	DOWN  = OrthogonalDirection(2)
	LEFT  = OrthogonalDirection(3)
)

type CompassDirection4 int8

const (
	NORTH = CompassDirection4(0)
	EAST  = CompassDirection4(1)
	SOUTH = CompassDirection4(2)
	WEST  = CompassDirection4(3)
)

type CompassDirection8 int8

const (
	N  = CompassDirection4(0)
	NE = CompassDirection4(1)
	E  = CompassDirection4(2)
	SE = CompassDirection4(3)
	S  = CompassDirection4(4)
	SW = CompassDirection4(5)
	W  = CompassDirection4(6)
	NW = CompassDirection4(7)
)
