package main

import (
	"errors"
	"net/http"

	"github.com/dekelund/robotmover/cmd/handlers"
	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func main() {
	controller := controllers.New()

	mux := handlers.NewMux(controller)

	err := http.ListenAndServe(":8000", mux)

	if errors.Is(err, http.ErrServerClosed) {
	} else if err != nil {
		panic(err)
	}
}
