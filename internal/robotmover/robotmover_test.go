package robotmover_test

import (
	"errors"
	"testing"

	"github.com/dekelund/robotmover/internal/robotmover"
)

func newRobotMover(xlimit, ylimit, xpos, ypos uint, dir robotmover.Direction) (*robotmover.RobotMover, error) {
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

	mover, err := robotmover.New(limits, position)

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if mover == nil {
		t.Fatal("unexpected mover value (nil)")
	}
}

func TestNew_invalidPositions(t *testing.T) {
	limits := robotmover.RoomLimits{robotmover.Coord{X: 10, Y: 15}}

	cases := map[string]struct {
		Coord robotmover.Coord
	}{
		"X outside of boundary": {
			robotmover.Coord{X: 11, Y: 5},
		},
		"Y outside of boundary": {
			robotmover.Coord{X: 5, Y: 16},
		},
		"Both X and Y is outside of boundary": {
			robotmover.Coord{X: 100, Y: 3000},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			position := robotmover.Position{
				Coord:     tc.Coord,
				Direction: robotmover.South,
			}

			_, err := robotmover.New(limits, position)

			if err == nil {
				t.Fatal("unexpected nil, expected an error")
			}

			var expectedErrorType robotmover.InvalidPositionError

			if errors.As(err, &expectedErrorType) != true {
				t.Fatal("unexpected error", err)
			}
		})
	}

}

func TestCoord_String(t *testing.T) {
	str := robotmover.Coord{X: 1, Y: 2}.String()

	if str != "1 2" {
		t.Fatal("unexpected position-string:", str)
	}
}

func TestPosition_String(t *testing.T) {
	str := robotmover.Position{
		Coord:     robotmover.Coord{X: 1, Y: 2},
		Direction: robotmover.South,
	}.String()

	if str != "1 2 S" {
		t.Fatal("unexpected position-string:", str)
	}
}
