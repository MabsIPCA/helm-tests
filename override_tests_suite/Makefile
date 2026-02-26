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

# Run a specific test (usage: make test-01)
test-%:
	@go run run-tests.go -test $*

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	@rm -f run-tests
	@find . -name "TEST-RESULTS.md" -delete
	@echo "Clean complete."

# Validate all charts
validate:
	@echo "Validating charts..."
	@for dir in test-*; do \
		if [ -f "$$dir/Chart.yaml" ]; then \
			echo "Validating $$dir..."; \
			helm lint "$$dir" || true; \
		fi \
	done

# Help target
help:
	@echo "Helm Values Override Tests"
	@echo ""
	@echo "Usage:"
	@echo "  make          - Run all tests"
	@echo "  make test     - Run all tests"
	@echo "  make test-01  - Run specific test (e.g., test-01)"
	@echo "  make build    - Build the test runner executable"
	@echo "  make validate - Validate all Helm charts"
	@echo "  make clean    - Remove generated files"
	@echo "  make help     - Show this help message"
	@echo ""
	@echo "Test results are written to each test folder's TEST-RESULTS.md"

