APP_NAME := k8s-inspector

.PHONY: build clean run help test test-verbose test-coverage

# Build the application
build:
	go build -o $(APP_NAME) ./main.go

# Clean build artifacts
clean:
	rm -f $(APP_NAME)

# Run the application
run-pods: build
	./$(APP_NAME) pods

# Run all tests
test:
	go test ./...

# Run tests with verbose output
test-verbose:
	go test ./... -v

# Run tests with coverage
test-coverage:
	go test ./... -cover

# Display help
help:
	@echo "Available targets:"
	@echo "  make build          - Compile the application"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make run-pods       - Build and run the pods command"
	@echo "  make test           - Run all tests"
	@echo "  make test-verbose   - Run all tests with verbose output"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make help           - Display this help message"
