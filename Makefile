# Go related variables
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOMODULE := $(shell go list -m)
GOOS := $(shell go env GOOS)

CMDS :=\
	bin/$(GOOS)/streamer\

TAG := $(shell git describe --tags --abbre=0 --dirty --always)
COMMIT := $(shell git rev-parse --short HEAD)

# Use linker flags to provide version/build settings to the target
LDFLAGS :=\

# Targets
.DEFAULT_GOAL := help

.PHONE: help
help: ## Display this help screen
	@echo "Available targets:"
	@grep -h -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-24s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: $(CMDS) ## Build applications

bin/$(GOOS)/streamer: CMD_TITLE := GoStreamer Server
bin/$(GOOS)/%:
	@echo "==> Building ${CMD_TITLE}"
	@go build -ldflags "$(LDFLAGS)" -o $@ ./cmd/$(shell basename "$@")

.PHONY: clean
clean: ## Clean the build directory
	@echo "==> Cleaning"
	@go clean ./...
	@rm -rf $(GOBIN)

.PHONY: test
test: ## Run tests
	@echo "==> Running tests"
	@go test -v -race ./...

.PHONY: bench
bench: ## Run benchmark tests
	@echo "==> Running benchmark tests"
	@go test -bench=. -benchmem ./...

.PHONY: cover
cover: ## Display test coverage
	@go test -cover -coverpkg ./... -coverprofile cover.out ./...
	@go tool cover -func cover.out
	@rm -f cover.out
