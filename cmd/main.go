package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// Bump to Go 1.21 to enable "log/slog"

	"github.com/dekelund/robotmover/cmd/handlers"
	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func main() {
	ctx := context.Background()

	controller := controllers.New()

	serv := http.Server{
		Addr:    ":8000",
		Handler: handlers.NewMux(controller),
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

		println("listening for signals")

		<-ch

		println("shutting down")

		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		if err := serv.Shutdown(ctx); err != nil {
			println("shutdown failed: ", err)
			//slog.Errorf("shutdown failed: %s", err)
		}
	}()

	println("listening for connections")
	err := serv.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		//slog.Errorf("failed to start server: %s", err)
	} else if err != nil {
		//slog.Errorf("failed to start server: %s", err)
	}
}
