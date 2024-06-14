package controllers_test

import (
	"testing"

	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func Test_ParseBoundaries(t *testing.T) {
	limits, err := controllers.ParseBoundaries("10 5")

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if limits.X != 10 {
		t.Fatal("unexpected x boundary", limits.X)
	}

	if limits.Y != 5 {
		t.Fatal("unexpected number of actions")
	}
}

func Test_ParseBoundaries_withLetter(t *testing.T) {
	_, err := controllers.ParseBoundaries("A 500")

	if err == nil {
		t.Fatal("x must not be a letter")
	}

	_, err = controllers.ParseBoundaries("6 B")

	if err == nil {
		t.Fatal("y must not be a letter")
	}

	_, err = controllers.ParseBoundaries("6 5B")

	if err == nil {
		t.Fatal("must fail on trailing letters")
	}
}

func Test_ParseBoundaries_withFloat(t *testing.T) {
	_, err := controllers.ParseBoundaries("10.0 5")

	if err == nil {
		t.Fatal("x must not be a float")
	}

	_, err = controllers.ParseBoundaries("10 5.0")

	if err == nil {
		t.Fatal("y must not be a float")
	}

	_, err = controllers.ParseBoundaries("10 5.0B")

	if err == nil {
		t.Fatal("y must not end with a letter")
	}
}
