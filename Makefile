# default
.PHONY: build
build:

########## Useful variables for paths, etcs

GO_CMD_WIRE_GEN := \
	$(shell find . -path './cmd/*' -type d -prune | xargs -I{} echo {}/wire_gen.go)
GO_SOURCES := \
	go.mod go.sum \
	$(shell find . -name '*.go' -not -name 'wire_gen.go')
GOOS := $(shell go env GOOS)
EXECUTABLE := bin/halebopp
EXEC_DUMMY := bin/dummy/halebopp

########## Utils

.PHONY: all
all: build dummy

.PHONY: generate
generate: ${GO_CMD_WIRE_GEN}

.PHONY: build
build: ${EXECUTABLE}

.PHONY: dummy
dummy: ${EXEC_DUMMY}

.PHONY: clean
clean:
	rm -rf bin ${GO_CMD_WIRE_GEN}

.PHONY: clean-generated
clean-generated:
	rm -f ${GO_CMD_WIRE_GEN}

.PHONY: upload
upload:
	if [[ ! -f "${EXECUTABLE}" ]]; then echo "not found: ${EXECUTABLE}" && exit 1; fi
	if [[ -z "$$HALEBOPP_REMOTE" ]]; then echo "HALEBOPP_REMOTE not set" && exit 1; fi
	set -ex && scp "${EXECUTABLE}" $${HALEBOPP_SSH_KEY:+-i "$$HALEBOPP_SSH_KEY"} "$$HALEBOPP_REMOTE:$${HALEBOPP_UPLOAD_DIR:-.}/"


########## Code generation

./cmd/%/wire_gen.go: ${GO_SOURCES}
	go run -mod=mod github.com/google/wire/cmd/wire ./$(@D)

########## Executable targets

${EXECUTABLE}: ${GO_CMD_WIRE_GEN} ${GO_SOURCES}
	GOOS=linux GOARCH=arm64 go build -o "${EXECUTABLE}" ./cmd/halebopp

${EXEC_DUMMY}: ${GO_CMD_WIRE_GEN} ${GO_SOURCES}
	go build -o "${EXEC_DUMMY}" ./cmd/halebopp-dummy
