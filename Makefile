install-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: install-lint
	@echo "Running linter..."
	@golangci-lint run
