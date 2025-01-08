_DEFAULT: build

GO_SOURCES := go.mod go.sum $(shell find . -name '*.go')
GOOS := $(shell go env GOOS)
EXECUTABLE := bin/halebopp
EXECUTABLE_TESTING := bin/testing/halebopp

########## Utils

.PHONY: all
all: build build-testing

.PHONY: generate
generate: ${GO_SOURCES}
	go generate ./...

.PHONY: build-raspi
build: bin/halebopp

.PHONY: build-testing
build-testing: bin/testing/halebopp

.PHONY: clean
clean:
	rm -rf bin

########## Executable targets

${EXECUTABLE}: generate ${GO_SOURCES}
	go build -o "${EXECUTABLE}" ./cmd/halebopp

${EXECUTABLE_TESTING}: generate ${GO_SOURCES}
	>&2 echo "TESTING EXECUTABLE NOT YET IMPLEMENTED" && exit 1
