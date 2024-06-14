package controllers_test

import (
	"testing"

	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func Test_ParseActions(t *testing.T) {
	input := "RFRFFRFRF"

	expected := []controllers.Action{
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

	result, err := controllers.ParseActions(input)

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if len(expected) != len(result) {
		t.Fatal("unexpected number of actions")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Fatal("unexpected result, expected", expected, "received", result)
		}
	}

}

// TODO(dekelund): Test invalid input
