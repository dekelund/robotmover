package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log/slog"

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

		slog.Info("listening for signals")

		<-ch

		slog.Info("shutting down")

		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		if err := serv.Shutdown(ctx); err != nil {
			slog.Error("shutdown failed", "error", err)
		}
	}()

	slog.Info("starts listening for connections")
	err := serv.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("server closed and shutdown")
	} else if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
