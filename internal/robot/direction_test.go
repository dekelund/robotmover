package robot

import (
	"testing"
)

func Test_parseDirection(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected Direction
	}{
		"north": {
			input:    "N",
			expected: North,
		}, "east": {
			input:    "E",
			expected: East,
		}, "south": {
			input:    "S",
			expected: South,
		}, "west": {
			input:    "W",
			expected: West,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			dir, err := parseDirection(tc.input)
			if err != nil {
				t.Fatal("unexpected exception: ", err)
			}

			if dir != tc.expected {
				t.Fatal("unexpected direction: ", dir)
			}
		})
	}
}

func Test_parseDirection_errors(t *testing.T) {
	cases := map[string]string{
		"lowercase - west":  "w",
		"lowercase - east":  "e",
		"lowercase - north": "n",
		"lowercase - south": "s",
		"unknown character": "k",
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := parseDirection(tc); err != ErrInvalidDirection {
				t.Fatal("unexpected exception: ", err)
			}
		})
	}
}
