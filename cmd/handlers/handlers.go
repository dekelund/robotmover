package handlers

//TODO(dekelund): Bump to go 1.21 to enable log/slog

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/dekelund/robotmover/internal/robot"
	"github.com/dekelund/robotmover/internal/robot/controllers"
)

var mutex sync.Mutex

func NewMux(controller *controllers.Controller) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/control", func(rw http.ResponseWriter, r *http.Request) {
		if !mutex.TryLock() {
			// Misusing WebDAV to indicate that robot is already
			// being controlled. We will not reply with a XML WebDAV
			// body.
			rw.WriteHeader(http.StatusLocked)
			return
		}
		defer mutex.Unlock()

		boundaries, position, actions, success := scanBody(r.Body)

		if !success {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to scan body")

			return
		}

		c, err := parseBody(boundaries, position, actions)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to parse body")

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

type controlBody struct {
	boundaries controllers.Boundaries
	position   robot.Position
	actions    []controllers.Action
}

func parseBody(boundaries, position, actions string) (controlBody, error) {
	b, err := controllers.ParseBoundaries(boundaries)
	if err != nil {
		return controlBody{}, err
	}

	p, err := robot.ParsePosition(position)
	if err != nil {
		return controlBody{}, err
	}

	a, err := controllers.ParseActions(actions)
	if err != nil {
		return controlBody{}, err
	}

	return controlBody{
		boundaries: b,
		position:   p,
		actions:    a,
	}, nil
}

func scanBody(r io.Reader) (boundaries, position, actions string, success bool) {
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
