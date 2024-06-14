package controllers

import (
	"fmt"

	"github.com/dekelund/robotmover/internal/robot"
)

type InvalidPositionError controllerError

func (e InvalidPositionError) Error() string {
	return string(e)
}

type InvalidDirectionError controllerError

func (e InvalidDirectionError) Error() string {
	return string(e)
}

type Boundaries struct {
	X int
	Y int
}

func ParseBoundaries(b string) (Boundaries, error) {
	var x, y int

	n, err := fmt.Sscanf(b, "%d %d\n", &x, &y)
	if err != nil {
		return Boundaries{}, err
	}

	if n != 2 {
		return Boundaries{}, fmt.Errorf("string not fully parsed, read %d of 2", n)
	}

	return Boundaries{X: x, Y: y}, nil
}

func (b Boundaries) ValidatePosition(pos robot.Position) error {
	if pos.X < 0 {
		return InvalidPositionError(fmt.Sprintf("x-position %d is outside of the mesh", pos.X))
	}

	if pos.Y < 0 {
		return InvalidPositionError(fmt.Sprintf("y-position %d is outside of the mesh", pos.Y))
	}

	if pos.X > b.X {
		return InvalidPositionError(fmt.Sprintf("x-position %d is outside of the mesh", pos.X))
	}

	if pos.Y > b.Y {
		return InvalidPositionError(fmt.Sprintf("y-position %d is outside of the mesh", pos.Y))
	}

	if pos.Direction < robot.North || pos.Direction > robot.West {
		return InvalidDirectionError(fmt.Sprintf("invalid direction %d", pos.Direction))
	}

	return nil
}
