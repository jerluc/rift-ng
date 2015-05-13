.PHONY: clean fmt build test

all: clean fmt build

clean:
	go clean

fmt:
	go fmt

build: test
	go build -o bin/rift

test:
	go test -v ./...
