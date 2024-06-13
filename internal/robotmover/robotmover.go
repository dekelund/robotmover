package robotmover

type direction string

const (
	North direction = "n"
	East  direction = "e"
	South direction = "s"
	West  direction = "w"
)

type Coord struct {
	X, Y uint
}

func NewCoord(x, y uint) Coord {
	return Coord{X: x, Y: y}
}

type Position struct {
	Coord
	Direction direction
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
