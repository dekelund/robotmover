package robotmover

import "fmt"

type InvalidPositionError string

func (e InvalidPositionError) Error() string {
	return string(e)
}

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

type Coord struct {
	X, Y uint
}

func NewCoord(x, y uint) Coord {
	return Coord{X: x, Y: y}
}

func (c Coord) String() string {
	return fmt.Sprintf("%d %d", c.X, c.Y)
}

type Position struct {
	Coord
	Direction Direction
}

func (p Position) String() string {
	return fmt.Sprintf("%s %s", p.Coord, p.Direction)
}

type RoomLimits struct {
	Coord
}

func (l RoomLimits) validate(pos Position) error {
	if pos.X > l.X {
		return InvalidPositionError(fmt.Sprintf("x-position %d is outside of the mesh", pos.X))
	}

	if pos.Y > l.Y {
		return InvalidPositionError(fmt.Sprintf("y-position %d is outside of the mesh", pos.Y))
	}

	return nil
}

type RobotMover struct {
	limits          RoomLimits
	currentPosition Position
}

func New(l RoomLimits, p Position) (*RobotMover, error) {
	if err := l.validate(p); err != nil {
		return nil, err
	}

	return &RobotMover{
		limits:          l,
		currentPosition: p,
	}, nil
}
