VERSION ?= dev
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
LDFLAGS := -X main.version=$(VERSION) -X main.commit=$(COMMIT)

.PHONY: all build run test lint clean

all: lint test build

build:
	go build -ldflags "$(LDFLAGS)" -o bin/k-mini ./

run: build
	./bin/k-mini

test:
	go test -race ./...

lint:
	go vet ./...
	test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)

clean:
	rm -rf bin/