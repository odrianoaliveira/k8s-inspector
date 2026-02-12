APP_NAME := k8s-inspector

.PHONY: build clean run help

# Build the application
build:
	go build -o $(APP_NAME) ./main.go

# Clean build artifacts
clean:
	rm -f $(APP_NAME)

# Run the application
run-pods: build
	./$(APP_NAME) pods

# Display help
help:
	@echo "Available targets:"
	@echo "  make build    - Compile the application"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make run-pods - Build and run the pods command"
	@echo "  make help     - Display this help message"
