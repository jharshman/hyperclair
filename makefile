SHELL=/bin/bash
RELEASE=hyperclair
VERSION:=$(shell cat VERSION)
BIN=./out

default: build

.PHONY: build
build: test
	gox -os="linux darwin" -arch="amd64 386" -output $(BIN)/{{.Dir}}-{{.OS}}-{{.Arch}} -ldflags="-X cmd.Version=$(VERSION)"

.PHONY: lint
lint:
	golint ./...

.PHONY: test
test: get-deps
	go test ./...

.PHONY: get-deps
get-deps:
	go get -d -v ./...

.PHONY: clean
clean:
	rm -f $(BIN)/$(RELEASE)-*
