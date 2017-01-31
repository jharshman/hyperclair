SHELL=/bin/bash
RELEASE=hyperclair
VERSION:=$(shell cat VERSION)
BIN=./out

default: build

.PHONY: build
build:
	go get -d -v ./...
	gox -os="linux darwin" -arch="amd64 386" -output $(BIN)/{{.Dir}}-{{.OS}}-{{.Arch}} -ldflags="-X cmd.Version=$(VERSION)"

.PHONY: lint
lint:
	golint ./...

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(BIN)/$(RELEASE)-*
