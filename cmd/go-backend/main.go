package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sajinlama/go-backend/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

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
