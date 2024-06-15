package controllers

import (
	"errors"

	"github.com/dekelund/robotmover/internal/robot"
)

type Action interface {
	exec(c *Controller) error
}

func ParseActions(actions string) ([]Action, error) {
	var errs []error

	result := make([]Action, 0, len(actions))

	for _, c := range actions {
		action, err := parseAction(c)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		result = append(result, action)
	}

	return result, errors.Join(errs...)
}

func parseAction(r rune) (Action, error) {
	switch r {
	case 'F':
		return WalkForward, nil

	case 'R':
		return TurnRight, nil

	case 'L':
		return TurnLeft, nil

	default:
		return TurnLeft, errors.New("unexpected action character")
	}
}

const (
	WalkForward walkForward = "F"
	TurnRight   turnRight   = "R"
	TurnLeft    turnLeft    = "L"
)

type walkForward string

func (f walkForward) exec(c *Controller) error {
	newPosition := c.currentPosition

	switch c.currentPosition.Direction {
	case robot.North:
		newPosition.Y--

	case robot.West:
		newPosition.X--

	case robot.South:
		newPosition.Y++

	case robot.East:
		newPosition.X++
	}

	if err := c.boundaries.validatePosition(newPosition); err != nil {
		return err
	}

	c.currentPosition = newPosition

	return nil
}

type turnRight string

func (f turnRight) exec(c *Controller) error {
	newPosition := c.currentPosition

	newPosition.Direction++
	if newPosition.Direction > robot.West {
		newPosition.Direction = robot.North
	}

	if err := c.boundaries.validatePosition(newPosition); err != nil {
		return err
	}

	c.currentPosition = newPosition

	return nil
}

type turnLeft string

func (f turnLeft) exec(c *Controller) error {
	newPosition := c.currentPosition

	newPosition.Direction--
	if newPosition.Direction < robot.North {
		newPosition.Direction = robot.West
	}

	if err := c.boundaries.validatePosition(newPosition); err != nil {
		return err
	}

	c.currentPosition = newPosition

	return nil
}
