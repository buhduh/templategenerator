DIR = $(realpath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

export GOPATH=${DIR}/gopath

export APPLICATION := templategenerator

BIN := ${DIR}/bin
EXE := ${BIN}/${APPLICATION}

BUILD := ${DIR}/build
BUILD_JS := ${BUILD}/js

# source variables
SRC := ${DIR}/src
GO := ${SRC}/go
FRONTEND := ${SRC}/frontend
SCRIPTS := ${DIR}/scripts

export DEBUG := true
ifdef DEBUG

# general application "constants"
export LOG_LEVEL = TRACE
export VERSION := $(shell cat VERSION)-$(shell date +"%m%d%Y%H%M%S")-debug
export PORT := 8080

# env dependent pathing "constants"
export TEMPLATE_DIR = ${GO}/templates
export RES := ${BUILD}/res
export RES_ROOT := ${RES}

# JS
export RES_JS := ${RES}/js
export MAIN_JS := ${RES_JS}/main.bundle.js

BUILD_DIRS := ${BUILD} ${BIN} ${RES} ${RES_JS}
else
$(error Only DEBUG mode is supported!)
endif


.PHONY: all
all: ${BUILD_DIRS} ${GOPATH} ${MAIN_JS}
	@${MAKE} -C ${GO}

.PHONY: ${MAIN_JS}
${MAIN_JS}:
	@${MAKE} -C ${FRONTEND}
	@cp ${FRONTEND}/build/main.bundle.js $@

${GOPATH}:
	@mkdir -p $@

.PHONY: deps
deps: ${GOPATH}
	@${MAKE} -C ${GO} deps

${BUILD_DIRS}:
	@mkdir -p $@

.PHONY: run
run: all
	${EXE}

.PHONY: clean
clean:
	@rm -rf ${BUILD_DIRS}
	@${MAKE} -C ${GO} clean
	@${MAKE} -C ${FRONTEND} clean

.PHONY: purge
purge: clean
	@echo "Purging ALL files generated by the Make system, including ${GOPATH}"
	@chmod -R u+w ${GOPATH}/
	@rm -rf ${GOPATH}
	@${MAKE} -C ${FRONTEND} purge
