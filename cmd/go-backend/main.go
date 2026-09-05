package main

import (
	"context"
	"github.com/sajinlama/go-backend/internal/config"
	"github.com/sajinlama/go-backend/internal/http/handlers/students"
	"github.com/sajinlama/go-backend/internal/storage/sqlite"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	//setups database

	storage, err1 := sqlite.New(cfg)
	if err1 != nil {
		log.Fatal(err1)
	}
	slog.Info("storage initialize ", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", students.New(storage))

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
