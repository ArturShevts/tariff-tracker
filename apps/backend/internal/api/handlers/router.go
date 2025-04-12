// apps/backend/internal/api/handlers/router.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates and configures a new router with all application routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Register routes
	api.HandleFunc("/leaderboard", GetLeaderboard).Methods("GET")
	api.HandleFunc("/country/{code}", GetCountry).Methods("GET")
	api.HandleFunc("/refresh-tariffs", RefreshTariffs).Methods("POST")

	// Health check endpoint
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	return r
}

// HealthCheck handles the health endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

// GetLeaderboard returns the tariff leaderboard data
func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`[{"country":"India","country_code":"IN","max_tariff":72.0,"direction":"retaliation"}]`))
}

// GetCountry returns data for a specific country
func GetCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"country":"Country %s","tariffs_from_country":[],"tariffs_from_us":[]}`, code)))
}

// RefreshTariffs triggers the tariff update process
func RefreshTariffs(w http.ResponseWriter, r *http.Request) {
	// Will be implemented later with the AI updater
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"update started"}`))
}
