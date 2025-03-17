install-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: install-lint
	@echo "Running linter..."
	@if golangci-lint run; then \
		echo "Linting successful!"; \
	else \
		exit 1; \
	fi
