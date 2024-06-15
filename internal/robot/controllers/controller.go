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

// SetBoundaries tells the controller how big the room is.
func (c *Controller) SetBoundaries(b Boundaries) {
	c.boundaries = b
}

// CalibratePosition tells the controller what field the robot is positioned
// with, and what direction it is facing.
func (c *Controller) CalibratePosition(p robot.Position) {
	c.currentPosition = p
}

// Exec controls the robot by navigating based on actions, such as:
// - walk forward
// - turn left
// - turn right
//
// It may return an error indicating invalid position or direction.
func (c *Controller) Exec(actions ...Action) error {
	var errs []error

	for _, a := range actions {
		if err := a.exec(c); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

// State will return the current position of the robot, in the following format:
// "X Y D", where X and Y correspond to which field, and D is a one of the
// following letters "NWSE", indicating what direction the robot is facing.
//
// For instance 3 1 E, if the robot is located in field (3, 1), facing east.
func (c *Controller) State() string {
	return c.currentPosition.String()
}
