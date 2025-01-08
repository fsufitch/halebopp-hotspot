_DEFAULT: build

GO_SOURCES := go.mod go.sum $(shell find . -name '*.go')
GOOS := $(shell go env GOOS)
EXECUTABLE := bin/halebopp
EXECUTABLE_TESTING := bin/testing/halebopp

########## Utils

.PHONY: all
all: build build-testing

.PHONY: build-raspi
build: wire bin/halebopp

.PHONY: build-testing
build-testing: bin/testing/halebopp

.PHONY: install-wire
install-wire:
	go install github.com/google/wire/cmd/wire@v0.6.0

.PHONY: wire
wire: install-wire ${GO_SOURCES}
	wire ./cmd/halebopp/

.PHONY: clean
clean:
	rm -rf bin

.PHONY: clean-all
clean-all: clean
	chmod -R u+w dev 
	git clean -d -f -x dev
	git checkout dev

########## Executable targets

${EXECUTABLE}: wire ${GO_SOURCES}
	go build -o "${EXECUTABLE}" ./cmd/halebopp

${EXECUTABLE_TESTING}: wire ${GO_SOURCES}
	>&2 echo "TESTING EXECUTABLE NOT YET IMPLEMENTED" && exit 1
