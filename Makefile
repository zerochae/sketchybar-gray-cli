.PHONY: build run clean install help

BINARY_NAME=gsbar
BUILD_DIR=bin

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

run: build
	@$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "Clean complete"

install: build
	@echo "Installing $(BINARY_NAME)..."
	@go install
	@echo "Install complete"

release:
	@echo "Building release binary..."
	@go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Release build complete: $(BUILD_DIR)/$(BINARY_NAME)"

help:
	@echo "Available targets:"
	@echo "  build    - Build the binary to $(BUILD_DIR)/"
	@echo "  run      - Build and run the binary"
	@echo "  clean    - Remove build artifacts"
	@echo "  install  - Install to GOPATH/bin"
	@echo "  release  - Build optimized release binary"
	@echo "  help     - Show this help message"
