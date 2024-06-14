package controllers

import (
	"errors"

	"github.com/dekelund/robotmover/internal/robot"
)

type controllerError string

func (e controllerError) Error() string {
	return string(e)
}

type Controller struct {
	boundaries      Boundaries
	currentPosition robot.Position
}

// New returns a newly allocated Controller, with current state
// in position (0, 0), facing in south direction on a 5x5 grid.
func New() *Controller {
	return &Controller{
		boundaries: Boundaries{5, 5},
		currentPosition: robot.Position{
			Coord:     robot.Coord{X: 0, Y: 0},
			Direction: robot.South,
		},
	}
}

func (c *Controller) SetBoundaries(b Boundaries) {
	c.boundaries = b
}

func (c *Controller) CalibratePosition(p robot.Position) {
	c.currentPosition = p
}

func (c *Controller) Exec(actions ...Action) error {
	var errs []error

	for _, a := range actions {
		if err := a.exec(c); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (c *Controller) String() string {
	return c.currentPosition.String()
}
