package controllers_test

import (
	"errors"
	"testing"

	"github.com/dekelund/robotmover/internal/robot"
	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func TestControllers_Exec_WalkForward(t *testing.T) {
	cases := map[string]struct {
		Start robot.Position
		End   string
	}{
		"forward - move north": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 5},
				Direction: robot.North,
			},
			End: "5 4 N",
		},
		"forward - move east": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 5},
				Direction: robot.East,
			},
			End: "6 5 E",
		},
		"forward - move south": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 5},
				Direction: robot.South,
			},
			End: "5 6 S",
		},
		"forward - move west": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 5},
				Direction: robot.West,
			},
			End: "4 5 W",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := controllers.New()
			c.SetBoundaries(controllers.Boundaries{10, 10})
			c.CalibratePosition(tc.Start)

			if err := c.Exec(controllers.WalkForward); err != nil {
				t.Fatal("unexpected error", err)
			}

			if str := c.State(); str != tc.End {
				t.Fatal("unexpected position", str)
			}
		})
	}
}

func TestControllers_Exec_TurnLeft(t *testing.T) {
	start := robot.Position{
		Coord:     robot.Coord{X: 5, Y: 5},
		Direction: robot.North,
	}

	c := controllers.New()
	c.SetBoundaries(controllers.Boundaries{10, 10})
	c.CalibratePosition(start)

	expected := []string{"W", "S", "E", "N"}

	for _, dir := range expected {
		if err := c.Exec(controllers.TurnLeft); err != nil {
			t.Fatal("unexpected error", err)
		}

		expectedPosition := "5 5 " + dir

		if str := c.State(); str != expectedPosition {
			t.Fatal("unexpected position/direction", str, "expected", expectedPosition)
		}
	}
}

func TestControllers_Exec_TurnRight(t *testing.T) {
	start := robot.Position{
		Coord:     robot.Coord{X: 5, Y: 5},
		Direction: robot.North,
	}

	c := controllers.New()
	c.SetBoundaries(controllers.Boundaries{10, 10})
	c.CalibratePosition(start)

	expected := []string{"E", "S", "W", "N"}

	for _, dir := range expected {
		if err := c.Exec(controllers.TurnRight); err != nil {
			t.Fatal("unexpected error", err)
		}

		expectedPosition := "5 5 " + dir

		if str := c.State(); str != expectedPosition {
			t.Fatal("unexpected position/direction", str, "expected", expectedPosition)
		}
	}
}

func TestControllers_Exec_InvalidPositionError(t *testing.T) {
	cases := map[string]struct {
		Start robot.Position
	}{
		"forward - move north": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 0},
				Direction: robot.North,
			},
		},
		"forward - move east": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 10, Y: 5},
				Direction: robot.East,
			},
		},
		"forward - move south": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 10},
				Direction: robot.South,
			},
		},
		"forward - move west": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 0, Y: 5},
				Direction: robot.West,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := controllers.New()
			c.SetBoundaries(controllers.Boundaries{10, 10})
			c.CalibratePosition(tc.Start)

			err := c.Exec(controllers.WalkForward)
			if err == nil {
				t.Fatal("unexpected success")
			}

			var expectedErrorType controllers.InvalidPositionError
			if errors.As(err, &expectedErrorType) != true {
				t.Fatal("unexpected error", err)
			}
		})
	}
}

func TestControllers_Exec_InvalidDirectionError(t *testing.T) {
	cases := map[string]struct {
		Start robot.Position
	}{
		"Undefined direction (-1)": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 5},
				Direction: robot.Direction(-1),
			},
		},
		"Undefined direction (West + 1)": {
			Start: robot.Position{
				Coord:     robot.Coord{X: 5, Y: 5},
				Direction: robot.West + 1,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := controllers.New()
			c.SetBoundaries(controllers.Boundaries{10, 10})
			c.CalibratePosition(tc.Start)

			err := c.Exec(controllers.WalkForward)
			if err == nil {
				t.Fatal("unexpected success")
			}

			var expectedErrorType controllers.InvalidDirectionError
			if errors.As(err, &expectedErrorType) != true {
				t.Fatal("unexpected error", err)
			}
		})
	}
}

func TestControllers_Exec_example1(t *testing.T) {
	boundaries := controllers.Boundaries{X: 5, Y: 5}

	start := robot.Position{
		Coord:     robot.NewCoord(1, 2),
		Direction: robot.North,
	}

	// Example: RFRFFRFRF
	actions := []controllers.Action{
		controllers.TurnRight,
		controllers.WalkForward,
		controllers.TurnRight,
		controllers.WalkForward,
		controllers.WalkForward,
		controllers.TurnRight,
		controllers.WalkForward,
		controllers.TurnRight,
		controllers.WalkForward,
	}

	c := controllers.New()
	c.SetBoundaries(boundaries)
	c.CalibratePosition(start)

	if err := c.Exec(actions...); err != nil {
		t.Fatal("unexpected error", err)
	}

	if str := c.State(); str != "1 3 N" {
		t.Fatal("unexpected position", str)
	}
}

func TestControllers_Exec_example2(t *testing.T) {
	boundaries := controllers.Boundaries{X: 5, Y: 5}

	start := robot.Position{
		Coord:     robot.NewCoord(0, 0),
		Direction: robot.East,
	}

	// Example: RFLFFLRF
	actions := []controllers.Action{
		controllers.TurnRight,
		controllers.WalkForward,
		controllers.TurnLeft,
		controllers.WalkForward,
		controllers.WalkForward,
		controllers.TurnLeft,
		controllers.TurnRight,
		controllers.WalkForward,
	}

	c := controllers.New()
	c.SetBoundaries(boundaries)
	c.CalibratePosition(start)

	if err := c.Exec(actions...); err != nil {
		t.Fatal("unexpected error", err)
	}

	if str := c.State(); str != "3 1 E" {
		t.Fatal("unexpected position", str)
	}
}
