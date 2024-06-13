package robotmover

import "fmt"

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

type Position struct {
	Coord
	Direction Direction
}

type RoomLimits struct {
	Coord
}

type RobotMover struct {
	limits          RoomLimits
	currentPosition Position
}

func New(l RoomLimits, p Position) *RobotMover {
	return &RobotMover{
		limits:          l,
		currentPosition: p,
	}
}

func (m *RobotMover) PositionAsString() string {
	return fmt.Sprintf("%d %d %s", m.currentPosition.X, m.currentPosition.Y, m.currentPosition.Direction)
}
