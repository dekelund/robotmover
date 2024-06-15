package handlers

//TODO(dekelund): Bump to go 1.21 to enable log/slog

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"github.com/dekelund/robotmover/internal/robot"
	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func NewMux(controller *controllers.Controller) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/control", func(rw http.ResponseWriter, r *http.Request) {
		boundaries, position, actions, success := scanCommands(r.Body)

		if !success {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to scan commands")

			return
		}

		c, err := parseCommands(boundaries, position, actions)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to parse commands")

			return
		}

		controller.SetBoundaries(c.boundaries)
		controller.CalibratePosition(c.position)

		if err := controller.Exec(c.actions...); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			//slog.Debug("handler faild to execute actions")

			return
		}

		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "Report: %s", controller.String())
	})

	return mux
}

type commands struct {
	boundaries controllers.Boundaries
	position   robot.Position
	actions    []controllers.Action
}

func parseCommands(boundaries, position, actions string) (commands, error) {
	b, err := controllers.ParseBoundaries(boundaries)
	if err != nil {
		return commands{}, err
	}

	p, err := robot.ParsePosition(position)
	if err != nil {
		return commands{}, err
	}

	a, err := controllers.ParseActions(actions)
	if err != nil {
		return commands{}, err
	}

	return commands{
		boundaries: b,
		position:   p,
		actions:    a,
	}, nil
}

func scanCommands(r io.Reader) (boundaries, position, actions string, success bool) {
		scanner := bufio.NewScanner(r)

		success = scanner.Scan()
		if !success {
			return
		}

		boundaries = scanner.Text()

		success = scanner.Scan()
		if !success {
			return
		}

		position = scanner.Text()

		success = scanner.Scan()
		if !success {
			return
		}

		actions = scanner.Text()

	return
}
