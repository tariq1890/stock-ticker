DOCKER_CLI ?= docker

REGISTRY ?= docker.io/tariq181290
TAG_PREFIX = v
VERSION = $(shell cat VERSION)
IMAGE = $(REGISTRY)/stock-ticker
TAG ?= $(TAG_PREFIX)$(VERSION)


.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

##@ Build

.PHONY: build
build: fmt vet ## Build manager binary.
	go build -o bin/ticker main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: container
container:
	${DOCKER_CLI} build --pull -t $(IMAGE):$(TAG) .

.PHONY: container-push
container-push:
	${DOCKER_CLI} push $(IMAGE):$(TAG)
