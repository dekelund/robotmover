package robotmover

import (
	"errors"
	"fmt"
)

type InvalidDirectionError string

func (e InvalidDirectionError) Error() string {
	return string(e)
}

type InvalidPositionError string

func (e InvalidPositionError) Error() string {
	return string(e)
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	}

	return "" // TODO(dekelund): how do we want to handle this scenario
}

type Coord struct {
	X, Y int
}

func NewCoord(x, y int) Coord {
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
	if pos.X < 0 {
		return InvalidPositionError(fmt.Sprintf("x-position %d is outside of the mesh", pos.X))
	}

	if pos.Y < 0 {
		return InvalidPositionError(fmt.Sprintf("y-position %d is outside of the mesh", pos.Y))
	}

	if pos.X > l.X {
		return InvalidPositionError(fmt.Sprintf("x-position %d is outside of the mesh", pos.X))
	}

	if pos.Y > l.Y {
		return InvalidPositionError(fmt.Sprintf("y-position %d is outside of the mesh", pos.Y))
	}

	if pos.Direction < North || pos.Direction > West {
		return InvalidDirectionError(fmt.Sprintf("invalid direction %d", pos.Direction))
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

type Action string

const (
	WalkForward Action = "F"
	TurnRight   Action = "R"
	TurnLeft    Action = "L"
)

func (m *RobotMover) Exec(actions ...Action) error {
	var errs []error

	for _, a := range actions {
		if err := m.Move(a); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

// Move provides functionalities to turn left, right and walk forward.
func (m *RobotMover) Move(a Action) error {
	switch a {
	case WalkForward:
		return m.walkForward()

	case TurnRight:
		return m.turnRight()

	case TurnLeft:
		return m.turnLeft()
	}

	return errors.New("invalid action")
}

func (m *RobotMover) String() string {
	return m.currentPosition.String()
}

func (m *RobotMover) walkForward() error {
	newPosition := m.currentPosition

	switch m.currentPosition.Direction {
	case North:
		newPosition.Y--

	case West:
		newPosition.X--

	case South:
		newPosition.Y++

	case East:
		newPosition.X++
	}

	if err := m.limits.validate(newPosition); err != nil {
		return err
	}

	m.currentPosition = newPosition

	return nil
}

func (m *RobotMover) turnLeft() error {
	newPosition := m.currentPosition

	newPosition.Direction--
	if newPosition.Direction < North {
		newPosition.Direction = West
	}

	if err := m.limits.validate(newPosition); err != nil {
		return err
	}

	m.currentPosition = newPosition

	return nil
}

func (m *RobotMover) turnRight() error {
	newPosition := m.currentPosition

	newPosition.Direction++
	if newPosition.Direction > West {
		newPosition.Direction = North
	}

	if err := m.limits.validate(newPosition); err != nil {
		return err
	}

	m.currentPosition = newPosition

	return nil
}
