package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("SSL_MODE"),
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	// Run migrations
	err = db.AutoMigrate(&Country{}, &Tariff{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed")
}

type Country struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Code    string `gorm:"size:2;unique;not null"`
	FlagURL string
}

type Tariff struct {
	ID            uint    `gorm:"primaryKey"`
	CountryID     uint    `gorm:"not null"`
	TargetCountry uint    `gorm:"not null"`
	Product       string  `gorm:"not null"`
	Type          string  `gorm:"not null"`
	Tariff        float64 `gorm:"not null"`
	LastUpdated   string
}

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

	// Initialize the database
	initDB()

	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Register routes
	api.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"country":"India","country_code":"IN","max_tariff":72.0,"direction":"retaliation"}]`))
	}).Methods("GET")

	api.HandleFunc("/country/{code}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		code := vars["code"]
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"country":"Country %s","tariffs_from_country":[],"tariffs_from_us":[]}`, code)))
	}).Methods("GET")

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	// Get host from environment or default to 0.0.0.0
	host := os.Getenv("API_HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	// Start server
	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: r,
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
}
