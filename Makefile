# This is a generic Makefile suitable for agylia projects writen
# in *GO*. Type `make` to get additional information on commands.
DOCKER_NAME  ?= "$(notdir $(shell pwd))"
DOCKER_IMAGE ?= "agylia/$(DOCKER_NAME)"

help: # This
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
.PHONY: help

deps: ## Install dependencies
	@echo "==> Installing dependencies"
	@go mod tidy
.PHONY: deps

gen:
	@go generate ./...
.PHONY: gen

dev: ## Start the project and and monitor source for changes (requires entr)
	@find . -name '*.go' -not -path './vendor/*' -not -path '*.gen.go' | entr -rcs 'make test && make start'
.PHONY: dev

test: ## Start the project tests
	@bash -c "set -a && source .env && set +a && \
		DEBUG=true \
		go test -cover ./..."
.PHONY: test

bench:
	@bash -c "set -a && source .env && set +a && \
		go test -run=^$$ -benchtime=.1s -bench=. ./..."
.PHONY: bench

start: gen ## Start the project
	@bash -c "set -a && source .env && set +a && \
		DEBUG=true \
		go run -ldflags '-X main.revision=`git rev-parse HEAD`' ./cmd/main.go"
.PHONY: run

clean: # Clean the project
	@go clean
	@git clean -f
.PHONY: clean
