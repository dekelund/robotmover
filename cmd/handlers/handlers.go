package handlers

//TODO(dekelund): Bump to go 1.21 to enable log/slog

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/dekelund/robotmover/internal/robot"
	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func NewMux(controller *controllers.Controller) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/control", func(rw http.ResponseWriter, r *http.Request) {
		scanner := bufio.NewScanner(r.Body)
		if !scanner.Scan() {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to scan boundaries")

			return
		}

		boundaries, err := controllers.ParseBoundaries(scanner.Text())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to parse boundaries")

			return
		}

		if !scanner.Scan() {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to scan position")

			return
		}

		position, err := robot.ParsePosition(scanner.Text())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to parse position")

			return
		}

		if !scanner.Scan() {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to scan actions")

			return
		}

		actions, err := controllers.ParseActions(scanner.Text())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to parse actions")

			return
		}

		controller.SetBoundaries(boundaries)
		controller.CalibratePosition(position)

		if err := controller.Exec(actions...); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			//slog.Debug("handler faild to execute actions")

			return
		}

		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "Report: %s", controller.String())
	})

	return mux
}
