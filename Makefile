# App Variables
APP_NAME=api
APP_DIR=cmd/$(APP_NAME)
BINARY_DIR=bin

# Docker Variables
DOCKER_DIR=deployments/docker

.PHONY: all
all: clean generate-docs build run

.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(BINARY_DIR)/$(APP_NAME) $(APP_DIR)/main.go
	@echo "Build completed."

.PHONY: run
run: build
	@echo "Running the application..."
	@./$(BINARY_DIR)/$(APP_NAME)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BINARY_DIR)/$(APP_NAME)
	@echo "Clean up completed."

.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v
	@echo "Tests completed."

.PHONY: format
format:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Code formatted."

.PHONY: generate-docs
generate-docs:
	@echo "Generating documentation..."
	@mkdir -p docs
	@swag init -d $(APP_DIR) -o docs --pdl 3
	@echo "Documentation generated in docs/ directory."

.PHONY: air
air:
	@echo "Starting Air for hot reloading..."
	@air
