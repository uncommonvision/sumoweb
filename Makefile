.PHONY: api-dev api-dev+ api-build api-test api-clean dragonfly-up dragonfly-down dragonfly-logs dragonfly-cli dragonfly-clean worker-dev worker-dev-queue worker-build worker-clean help

api-dev:
	@echo "Starting api in development mode..."
	@set -a && . ./env/api.env && set +a && cd api && ENVIRONMENT=development go run cmd/api/main.go

api-dev+:
	@echo "Starting api in development mode (publicly accessible)..."
	@set -a && . ./env/api.env && set +a && cd api && ENVIRONMENT=development HOST=0.0.0.0 go run cmd/api/main.go

api-build:
	@echo "Building api..."
	@cd api && go build -o ./bin/api cmd/api/main.go
	@echo "Build complete: ./api/bin/api"

api-test:
	@echo "Running api tests..."
	@set -a && . ./env/api.env && set +a && cd api && ENVIRONMENT=development HOST=0.0.0.0 go test ./api/tests/... -v

api-clean:
	@echo "Cleaning api build artifacts..."
	@rm -rf api/bin/*
	@echo "Clean complete"

dragonfly-up:
	@echo "Starting Dragonfly container..."
	@docker compose up -d dragonfly
	@echo "Dragonfly is running on localhost:6379"

dragonfly-down:
	@echo "Stopping Dragonfly container..."
	@docker compose down dragonfly

dragonfly-logs:
	@echo "Showing Dragonfly logs..."
	@docker compose logs -f dragonfly

dragonfly-cli:
	@echo "Connecting to Dragonfly CLI..."
	@docker compose exec dragonfly redis-cli

dragonfly-clean:
	@echo "Removing Dragonfly data..."
	@docker compose down dragonfly -v

web-dev:
	@echo "Starting web in development mode..."
	@set -a && . ./env/web.env && set +a && cd web && bun run dev+

web-dev+:
	@echo "Starting web in development mode (publically accessible)..."
	@set -a && . ./env/web.env && set +a && cd web && bun run dev+

web-build:
	@echo "Building web..."
	@cd web && bun run build
	@echo "Build complete"

web-preview:
	@echo "Starting web in preview mode..."
	@set -a && . ./env/web.env && set +a && cd web && bun run preview

web-preview+:
	@echo "Starting web in preview mode (publically accessible)..."
	@set -a && . ./env/web.env && set +a && cd web && bun run preview+

worker-dev:
	@echo "Starting worker in development mode..."
	@set -a && . ./env/worker.env && set +a && cd api && ENVIRONMENT=development go run cmd/worker/main.go

worker-dev-queue:
	@echo "Starting worker for queue: $(QUEUE)"
	@cd api && ENVIRONMENT=development QUEUE_NAME=$(QUEUE) go run cmd/worker/main.go
	@set -a && . ./env/worker.env && set +a && cd api && ENVIRONMENT=development QUEUE_NAME=$(QUEUE) go run cmd/worker/main.go

worker-build:
	@echo "Building worker..."
	@set -a && . ./env/worker.env && set +a && cd api && go build -o ./bin/api-worker cmd/worker/main.go
	@echo "Build complete: ./api/bin/api-worker"

worker-clean:
	@echo "Cleaning worker build artifacts..."
	@rm -f api/bin/api-worker
	@echo "Clean complete"

help:
	@echo "Available targets:"
	@echo ""
	@echo "Api:"
	@echo "  make api-dev     - Start api in development mode (localhost only)"
	@echo "  make api-dev+    - Start api in development mode (publicly accessible)"
	@echo "  make api-build   - Build api binary"
	@echo "  make api-test    - Run api tests"
	@echo "  make api-clean   - Remove api build artifacts"
	@echo ""
	@echo "Worker:"
	@echo "  make worker-dev       - Start worker in development mode"
	@echo "  make worker-dev-queue - Start worker for specific queue (QUEUE=name)"
	@echo "  make worker-build     - Build worker binary"
	@echo "  make worker-clean     - Remove worker build artifacts"
	@echo ""
	@echo "Dragonfly:"
	@echo "  make dragonfly-up   - Start Dragonfly container"
	@echo "  make dragonfly-down - Stop Dragonfly container"
	@echo "  make dragonfly-logs - Show Dragonfly logs"
	@echo "  make dragonfly-cli  - Connect to Dragonfly CLI"
	@echo "  make dragonfly-clean - Remove Dragonfly container and data"
	@echo ""
	@echo "General:"
	@echo "  make help           - Show this help message"
