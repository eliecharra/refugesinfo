VERSION := $(shell git describe --always --long --dirty)

build-all:
	@go get github.com/mitchellh/gox
	@gox -ldflags="-X main.version=${VERSION}"

build:
	@go build -ldflags="-X main.version=${VERSION}"
