package robotmover

type direction string

const (
	North direction = "n"
	East  direction = "e"
	South direction = "s"
	West  direction = "w"
)

type Position struct {
	X, Y      uint
	Direction direction
}

type RoomLimits struct {
	X, Y uint
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
