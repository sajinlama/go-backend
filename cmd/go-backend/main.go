package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sajinlama/go-backend/internal/config"
	"github.com/sajinlama/go-backend/internal/http/handlers/students"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", students.New())

	server := http.Server{
		Addr:    cfg.HttpsServer.Add,
		Handler: router,
	}
	slog.Info(" server started ", slog.String("address", cfg.HttpsServer.Add))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done
	slog.Info("shutting down the sever")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
