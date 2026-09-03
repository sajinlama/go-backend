package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Printf("Server starting on %s...\n", cfg.HttpsServer.Add)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}
