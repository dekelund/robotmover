package robotmover_test

import (
	"errors"
	"testing"

	"github.com/dekelund/robotmover/internal/robotmover"
)

func newRobotMover(xlimit, ylimit, xpos, ypos int, dir robotmover.Direction) (*robotmover.RobotMover, error) {
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
		t.Fatal("unexpected value", mover)
	}
}

func TestNew_invalidPositions(t *testing.T) {
	limits := robotmover.RoomLimits{robotmover.Coord{X: 10, Y: 15}}

	cases := map[string]struct {
		Coord robotmover.Coord
	}{
		"X outside of min-boundary": {
			robotmover.Coord{X: -1, Y: 5},
		},
		"Y outside of min-boundary": {
			robotmover.Coord{X: 5, Y: -1},
		},
		"X outside of max-boundary": {
			robotmover.Coord{X: 11, Y: 5},
		},
		"Y outside of max-boundary": {
			robotmover.Coord{X: 5, Y: 16},
		},
		"Both X and Y is outside of max-boundary": {
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

func TestRobotMover_Move(t *testing.T) {
	cases := map[string]struct {
		Start robotmover.Position
		End   string
	}{
		"forward - move north": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 5, Y: 5},
				Direction: robotmover.North,
			},
			End: "5 4 N",
		},
		"forward - move east": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 5, Y: 5},
				Direction: robotmover.East,
			},
			End: "6 5 E",
		},
		"forward - move south": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 5, Y: 5},
				Direction: robotmover.South,
			},
			End: "5 6 S",
		},
		"forward - move west": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 5, Y: 5},
				Direction: robotmover.West,
			},
			End: "4 5 W",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mover, err := robotmover.New(robotmover.RoomLimits{robotmover.NewCoord(10, 10)}, tc.Start)
			if err != nil {
				t.Fatal("unexpected error", err)
			}

			err = mover.Move(robotmover.WalkForward)
			if err != nil {
				t.Fatal("unexpected error", err)
			}

			str := mover.String()
			if str != tc.End {
				t.Fatal("unexpected position", str)
			}
		})
	}
}

func TestRobotMover_Move_outside_boundary(t *testing.T) {
	cases := map[string]struct {
		Start robotmover.Position
	}{
		"forward - move north": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 5, Y: 0},
				Direction: robotmover.North,
			},
		},
		"forward - move east": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 10, Y: 5},
				Direction: robotmover.East,
			},
		},
		"forward - move south": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 5, Y: 10},
				Direction: robotmover.South,
			},
		},
		"forward - move west": {
			Start: robotmover.Position{
				Coord:     robotmover.Coord{X: 0, Y: 5},
				Direction: robotmover.West,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mover, err := robotmover.New(robotmover.RoomLimits{robotmover.NewCoord(10, 10)}, tc.Start)
			if err != nil {
				t.Fatal("unexpected error", err)
			}

			err = mover.Move(robotmover.WalkForward)

			var expectedErrorType robotmover.InvalidPositionError
			if errors.As(err, &expectedErrorType) != true {
				t.Fatal("unexpected error", err)
			}
		})
	}
}
