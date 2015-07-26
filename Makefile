.PHONY: clean fmt build test run

all: clean test build

clean:
	go clean

fmt:
	go fmt

test:
	go test -v ./...

build: test
	go build -o bin/rift

run: build
	./bin/rift
