package robot

const (
	ErrInvalidDirection robotError = "invalid direction"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func ParseDirection(s string) (Direction, error) {
	switch s {
	case "N":
		return North, nil
	case "E":
		return East, nil
	case "S":
		return South, nil
	case "W":
		return West, nil
	}

	return North, ErrInvalidDirection
}

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

	return "-" // NOTE(dekelund): It's tempting to panic at this point.
}
