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

.PHONY: upload
upload:
	if [[ ! -f "${EXECUTABLE}" ]]; then echo "not found: ${EXECUTABLE}" && exit 1; fi
	if [[ -z "$$HALEBOPP_REMOTE" ]]; then echo "HALEBOPP_REMOTE not set" && exit 1; fi
	set -ex && scp bin/halebopp $${HALEBOPP_SSH_KEY:+-i "$$HALEBOPP_SSH_KEY"} "$$HALEBOPP_REMOTE:$${HALEBOPP_UPLOAD_DIR:-.}/"

########## Executable targets

${EXECUTABLE}: generate ${GO_SOURCES}
	go build -o "${EXECUTABLE}" ./cmd/halebopp

${EXECUTABLE_TESTING}: generate ${GO_SOURCES}
	>&2 echo "TESTING EXECUTABLE NOT YET IMPLEMENTED" && exit 1
