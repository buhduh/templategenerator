ifndef GOPATH
$(error Environment variable "GOPATH" is required.)
endif

ifndef VERSION
$(error Environment varialbe "VERSION" is required.)
endif

ifndef APPLICATION
APPLICATION=templategenerator
endif

SRC = $(foreach f,$(shell find . -type f -name '*.go'),$(realpath ${f}))
MAKEFILE = $(abspath $(lastword $(MAKEFILE_LIST)))

BIN_DIR = $(realpath ../../bin/)
EXE = ${BIN_DIR}/${APPLICATION}

CONST := templategenerator/constants
LDXFLAGS +=-X ${CONST}.APPLICATION=${APPLICATION}
LDXFLAGS +=-X ${CONST}.TEMPLATE_DIR=${TEMPLATE_DIR}
LDXFLAGS +=-X ${CONST}.LOG_LEVEL=${LOG_LEVEL}
LDXFLAGS +=-X ${CONST}.VERSION=${VERSION}
LDXFLAGS +=-X ${CONST}.DEF_PORT=${PORT}
LDXFLAGS +=-X ${CONST}.RES_ROOT=${RES_ROOT}

ifndef DEBUG
$(error Only DEBUG mode is supported!)
endif

LDFLAGS :=-ldflags="${LDXFLAGS}"

# TODO, "build" shouldn't be the default target
# the "real" default target will do a bunch of stuff i don't want to do while iterating
# like go getting deps, etc.,
.PHONY: build
build:  ${EXE}

.PHONY: all
all: ${EXE}

#TODO compile for linux and mac
${EXE}: ${SRC} ${MAKEFILE}
	GOPATH=${GOPATH} go build -o $@ ${LDFLAGS}

.PHONY: clean
clean:
	@echo cleaning go ../../bin/${APPLICATION}

.PHONY: deps
deps:
	@GOPATH=${GOPATH} go get -u -v ./...
