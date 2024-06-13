package robotmover_test

import (
	"testing"

	"github.com/dekelund/robotmover/internal/robotmover"
)

func newRobotMover(xlimit, ylimit, xpos, ypos uint, dir robotmover.Direction) *robotmover.RobotMover {
	limits := robotmover.RoomLimits{robotmover.Coord{X: xlimit, Y: ylimit}}

	position := robotmover.Position{
		Coord:     robotmover.Coord{X: xpos, Y: ypos},
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
		t.Fatal("unexpected mover value (nil)")
	}
}

func TestRobotMover_PositionAsString(t *testing.T) {
	mover := newRobotMover(10, 10, 1, 2, robotmover.South)

	if mover == nil {
		t.Fatal("unexpected nil")
	}

	position := mover.PositionAsString()
	if position != "1 2 S" {
		t.Fatal("unexpected position-string:", position)
	}
}
