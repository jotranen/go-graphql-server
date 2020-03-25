.PHONY: clean

BINARY        ?= go-graphql-server
VERSION       ?= $(shell git describe --tags --always --dirty)
SOURCES       = $(shell find . -name '*.go')
DOCKERFILE    ?= Dockerfile
BUILD_FLAGS   ?= -v
LDFLAGS       ?= -X main.version=$(VERSION) -w -s

default: build.local

clean:
	rm -rf build

build.local: build/$(BINARY)

build/$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" ./cmd/$(BINARY)

