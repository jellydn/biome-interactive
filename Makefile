.PHONY: dev
dev:
	@echo "Starting application..."
	@go run main.go

.PHONY: build
build:
	@echo "Building binary..."
	@go build -o bin/$(APP_NAME) main.go

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf bin

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  dev    - Start development server"
	@echo "  build  - Build binary"
	@echo "  test   - Run tests"
	@echo "  clean  - Clean up"
	@echo "  help   - Display this help message"
	@echo ""
	@echo "Variables:"
	@echo "  APP_NAME - Name of the application (default: $(APP_NAME))"
	@echo ""
