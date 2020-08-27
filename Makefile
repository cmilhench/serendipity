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

pg-start: ## Start a development postgresql database server
	@docker run --rm --detach \
					--volume "${PWD}/config/data/pg":/var/lib/postgresql/data \
					--name pg-$(DOCKER_NAME) \
					--env POSTGRES_PASSWORD=docker \
					--publish 5432:5432 \
	  			postgres
.PHONY: pg-start

pg-stop: ## Stop a development postgresql database server
	@docker stop pg-$(DOCKER_NAME)
.PHONY: pg-stop

es-start: ## Start a development elasticsearch database server
	@docker run --rm --detach \
					--mount type=bind,src=${PWD}/config/data/es,dst=/usr/share/elasticsearch/data \
					--name es-$(DOCKER_NAME) \
					--env "cluster.name=$(DOCKER_NAME)-cluster" \
					--env "network.host: 0.0.0.0" \
					--env "discovery.type=single-node" \
					--env "bootstrap.memory_lock=true" \
					--env "path.repo=/usr/share/elasticsearch/data/snaps" \
					--env ES_JAVA_OPTS="-Xms2g -Xmx2g" \
					--publish 9200:9200 \
					--publish 9300:9300 \
					docker.elastic.co/elasticsearch/elasticsearch:7.6.2
.PHONY: es-start

es-stop: ## Stop a development elasticsearch database server
	@docker stop es-$(DOCKER_NAME)
.PHONY: es-stop

docker-build: clean deps gen ## Build the project as a docker image
	@go mod vendor
	@docker run --rm \
					--volume "${PWD}":/go/src/github.com/cmilhench/$(DOCKER_NAME) \
					--workdir /go/src/github.com/cmilhench/$(DOCKER_NAME) \
					golang:alpine go build -ldflags '-X main.revision=$(shell git rev-parse HEAD)'
	@docker build --tag $(DOCKER_IMAGE) .
	@go clean
	@rm -fr vendor
.PHONY: docker-build

docker-start: ## Start the project in a created docker image
	@docker run --rm \
		--mount type=bind,src=/etc/ssl/certs,dst=/etc/ssl/certs \
		--name $(DOCKER_NAME) \
		--publish 8080:8080 \
		--env-file=.env \
		--env DEBUG=true \
		$(DOCKER_IMAGE)
.PHONY: docker-run

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
