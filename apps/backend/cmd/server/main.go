// apps/backend/cmd/server/main.go
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArturShevts/tariff-tracker/apps/backend/internal/api/handlers"
	"github.com/ArturShevts/tariff-tracker/apps/backend/internal/api/middleware"
	"github.com/ArturShevts/tariff-tracker/apps/backend/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or cannot be loaded")
	}

	// Use API_PORT from env or default to 8080
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	flag.StringVar(&port, "port", port, "API server port")
	flag.Parse()

	// Get host from environment or default to 0.0.0.0
	host := os.Getenv("API_HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	// Initialize the database
	db.InitDB()

	// Initialize router with handlers
	router := handlers.NewRouter()

	// Add middleware
	router.Use(middleware.Logger)
	router.Use(middleware.CORS)

	// Create server
	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: router,
	}

	// Server in a goroutine so shutdown can be handled gracefully
	go func() {
		log.Printf("Server starting on %s:%s", host, port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
