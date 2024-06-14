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

	parsedDir, err := ParseDirection(dir)
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
