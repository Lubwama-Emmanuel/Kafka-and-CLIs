GOBASE = $(shell pwd)
LINT_PATH = $(GOBASE)/build/lint
MOCKGEN_PATH = $(GOBASE)/build/mockgen

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Fetch required dependencies
	go mod tidy -compat=1.17
	go mod download
	
generate-mocks:
	rm -r **/mock || :
	$(MOCKGEN_PATH)/mockgen -package mock -source consumers/consumer.go -destination consumers/mock/consumer.go

lint: ## Linter for developers
	$(LINT_PATH)/golangci-lint run --timeout=5m -c .golangci.yml

install-golangci: ## Install the correct version of lint
	GOBIN=$(LINT_PATH) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

install-mockgen: ## Install the correct version of mockgen
	GOBIN=$(MOCKGEN_PATH) go install github.com/golang/mock/mockgen@v1.6.0
