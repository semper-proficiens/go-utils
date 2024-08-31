.PHONY: help session
help:
	@grep -E '^[a-zA-Z0-9_-]+%?:.*?## .*$$' $(MAKEFILE_LIST) | sed -e 's/^Makefile://' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

tests: ## runs all tests and benchmarks with coverage
	go test -bench . -benchmem ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
	go vet ./...
	golangci-lint run
	gosec ./...