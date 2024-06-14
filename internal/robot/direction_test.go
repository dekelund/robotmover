package robot_test

import (
	"testing"

	"github.com/dekelund/robotmover/internal/robot"
)

func TestParseDirection(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected robot.Direction
	}{
		"north": {
			input:    "N",
			expected: robot.North,
		}, "east": {
			input:    "E",
			expected: robot.East,
		}, "south": {
			input:    "S",
			expected: robot.South,
		}, "west": {
			input:    "W",
			expected: robot.West,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			dir, err := robot.ParseDirection(tc.input)
			if err != nil {
				t.Fatal("unexpected exception: ", err)
			}

			if dir != tc.expected {
				t.Fatal("unexpected direction: ", dir)
			}
		})
	}
}

func TestParseDirection_errors(t *testing.T) {
	cases := map[string]string{
		"lowercase - west":  "w",
		"lowercase - east":  "e",
		"lowercase - north": "n",
		"lowercase - south": "s",
		"unknown character": "k",
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := robot.ParseDirection(tc); err != robot.ErrInvalidDirection {
				t.Fatal("unexpected exception: ", err)
			}
		})
	}
}
