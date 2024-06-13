package robotmover_test

import (
	"testing"

	"github.com/dekelund/robotmover/internal/robotmover"
)

func newRobotMover() *robotmover.RobotMover {
	limits := robotmover.RoomLimits{robotmover.Coord{X: 10, Y: 10}}

	position := robotmover.Position{
		Coord:     robotmover.Coord{X: 0, Y: 0},
		Direction: robotmover.South,
	}

	return robotmover.New(limits, position)
}

func TestNew(t *testing.T) {
	limits := robotmover.RoomLimits{robotmover.Coord{X: 10, Y: 10}}

	position := robotmover.Position{
		Coord:     robotmover.Coord{X: 0, Y: 0},
		Direction: robotmover.South,
	}

	mover := robotmover.New(limits, position)

	if mover == nil {
		t.Fatal("unexpected nil")
	}
}
