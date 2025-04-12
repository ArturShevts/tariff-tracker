include .env
export

.PHONY: run
run:
	# Check if port is already in use and stop any existing processes
	-@lsof -ti:${API_PORT} | xargs kill -9 2>/dev/null || true

	# Start just the database container
	docker-compose up -d db

	# Wait for PostgreSQL to fully initialize
	@echo "Waiting for PostgreSQL to start..."
	@sleep 5

	# Copy environment file and run the backend in background
	@cp .env apps/backend/.env
	cd apps/backend && go run cmd/server/main.go > /dev/null 2>&1 & echo $$! > .server.pid

	# Wait for server to initialize
	@echo "Waiting for server to start..."
	@sleep 3

	# Test endpoints and display results
	@echo "\n--- Testing /health endpoint ---"
	@curl -s http://localhost:8080/health | jq || echo "Failed to connect to health endpoint"

	@echo "\n--- Testing /api/leaderboard endpoint ---"
	@curl -s http://localhost:8080/api/leaderboard | jq || echo "Failed to connect to leaderboard endpoint"

	# Keep the server running in foreground
	@echo "\n--- Server is running at http://localhost:8080 ---"
	@echo "Press Ctrl+C to stop"
	@tail -f /dev/null

.PHONY: stop
stop:
	-@kill `cat .server.pid` 2>/dev/null || true
	-@rm .server.pid 2>/dev/null || true
	-@lsof -ti:${API_PORT} | xargs kill -9 2>/dev/null || true
	docker-compose down

.PHONY: test
test:
	cd apps/backend && go test ./...

.PHONY: stop-db
stop-db:
	docker-compose down -v
	docker-compose up -d db
	@echo "Database reset and restarted"