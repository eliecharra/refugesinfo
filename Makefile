VERSION := $(shell git describe --always --long --dirty)

build:
	@go build -ldflags="-X main.version=${VERSION}"
