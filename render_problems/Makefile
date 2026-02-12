.PHONY: all test build clean help

# Default target - run all tests
all: test

# Run all tests using the main runner
test:
	@go run run-tests.go

# Build the test runner executable
build:
	@echo "Building test runner..."
	@go build -o run-tests run-tests.go
	@echo "Build complete: ./run-tests"

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	@rm -f run-tests
	@echo "Clean complete."

# Help target
help:
	@echo "Helm Template Error Tests"
	@echo ""
	@echo "Usage:"
	@echo "  make          - Run all tests (13, 14, 15, 16)"
	@echo "  make test     - Run all tests"
	@echo "  make build    - Build the test runner executable"
	@echo "  make clean    - Remove generated executables"
	@echo "  make help     - Show this help message"
	@echo ""
	@echo "Test results are written to each test folder's TEST-RESULTS.md"

