.PHONY: dev build install image release clean

CGO_ENABLED=0
VERSION=$(shell git describe --abbrev=0 --tags)
COMMIT=$(shell git rev-parse --short HEAD)

all: dev

dev: build
	./petstore

build: clean
	@go build \
		-tags "netgo static_build" -installsuffix netgo \
		-ldflags "-w -X $(shell go list)/internal.Version=$(VERSION) -X $(shell go list)/internal.Commit=$(COMMIT)" \
		./cmd/petstore/...

install: build
	@go install ./cmd/petstore/...

image:
	@docker build -t prologic/petstore .

release:
	@./tools/release.sh

clean:
	@git clean -f -d -X
