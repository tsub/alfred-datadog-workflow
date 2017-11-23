PROJECT  = alfred-datadog-workflow
SRC      ?= $(shell go list ./... | grep -v vendor)
TESTARGS ?= -v

deps:
	dep ensure
.PHONY: deps

build:
	go build -o build/$(PROJECT)
	cp assets/* build/
.PHONY: build

test:
	go test $(SRC) $(TESTARGS)
.PHONY: test

fmt:
	go fmt $(SRC)
.PHONY: fmt

vet:
	go vet $(SRC)
.PHONY: vet
