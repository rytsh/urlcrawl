BINARY    := urlcrawl
MAIN_FILE := cmd/$(BINARY)/main.go

PKG       := $(shell go list -m)
VERSION   := $(or $(IMAGE_TAG),$(shell git describe --tags --first-parent --match "v*" 2> /dev/null || echo v0.0.0))

LOCAL_BIN_DIR := $(PWD)/bin

BUILD_DATE := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(or $(IMAGE_TAG),$(shell git describe --tags --first-parent --match "v*" 2> /dev/null || echo v0.0.0))

.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the application
	go run $(MAIN_FILE)

# go build -trimpath -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.date=$(BUILD_DATE)" -o bin/$(BINARY_NAME) $(BINARY_PATH)
.PHONY: build
build: ## Build the binary
	goreleaser build --snapshot --clean --single-target

.PHONY: test
test: ## Run the tests
	go test -v -race ./...

.PHONY: tools
tools: ## Download tools (mockgen)
	go install go.uber.org/mock/mockgen@latest

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
