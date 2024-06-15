package robot

import (
	"fmt"
)

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

// ParsePosition parses a string with following format:
// "X Y D", where X and Y correspond to which field, and D is a one of the
// following letters "NWSE", indicating what direction the robot is facing.
//
// For instance 3 1 E, if the robot is located in field (3, 1), facing east.
//
// It returns a Position based on parsed values, or an error for malformed
// strings.
func ParsePosition(p string) (Position, error) {
	var x, y int
	var dir string

	n, err := fmt.Sscanf(p, "%d %d %s\n", &x, &y, &dir)
	if err != nil {
		return Position{}, err
	}

	if n != 3 {
		return Position{}, fmt.Errorf("string not fully parsed, read %d of 2", n)
	}

	parsedDir, err := parseDirection(dir)
	if err != nil {
		return Position{}, err
	}

	return Position{
		Coord: Coord{
			X: x, Y: y,
		},
		Direction: parsedDir,
	}, nil
}
