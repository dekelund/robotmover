package robotmover_test

import (
	"testing"

	"github.com/dekelund/robotmover/internal/robotmover"
)

func TestNew(t *testing.T) {
	limits := robotmover.RoomLimits{
		X: 10,
		Y: 10,
	}

	position := robotmover.Position{
		X: 0,
		Y: 0,
		Direction: robotmover.South,
	}

	mover := robotmover.New(limits, position)

	if mover == nil {
		t.Fatal("unexpected nil")
	}
}
