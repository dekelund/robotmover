package robot_test

import (
	"testing"

	"github.com/dekelund/robotmover/internal/robot"
)

func Test_ParsePosition(t *testing.T) {
	cases := map[string]struct {
		position    string
		expectedX   int
		expectedY   int
		expectedDir robot.Direction
	}{
		"north": {
			position:    "10 5 N",
			expectedX:   10,
			expectedY:   5,
			expectedDir: robot.North,
		},
		"west": {
			position:    "19 5 W",
			expectedX:   19,
			expectedY:   5,
			expectedDir: robot.West,
		},
		"east": {
			position:    "10 500 E",
			expectedX:   10,
			expectedY:   500,
			expectedDir: robot.East,
		},
		"south": {
			position:    "10001 5 S",
			expectedX:   10001,
			expectedY:   5,
			expectedDir: robot.South,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			limits, err := robot.ParsePosition(tc.position)

			if err != nil {
				t.Fatal("unexpected error", err)
			}

			if limits.X != tc.expectedX {
				t.Fatal("unexpected x boundary", limits.X)
			}

			if limits.Y != tc.expectedY {
				t.Fatal("unexpected number of actions")
			}

			if limits.Direction != tc.expectedDir {
				t.Fatal("unexpected number of actions")
			}
		})
	}
}

func Test_ParsePosition_invalidFormat(t *testing.T) {
	cases := map[string]string{
		"x as float":         "9.0 5 W",
		"y as float":         "10 5.0 N",
		"trailing character": "10 500 EW",
		"invalid direction":  "10001 5 A",
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := robot.ParsePosition(tc); err == nil {
				t.Fatal("unexpected success")
			}
		})
	}
}
