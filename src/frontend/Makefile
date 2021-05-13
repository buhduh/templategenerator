ifndef VERSION
$(error VERSION is required)
else
NPM_TGT := dev
endif

DIR := $(realpath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
SCRIPTS := ${DIR}/scripts
BUILD := ${DIR}/build
BUILD_DIRS := ${BUILD}
VER_BUILD := ${BUILD}/version

PKG_JSON := ${DIR}/package.json
WEBPACK_JS := ${DIR}/webpack.config.js

MAIN := ${BUILD}/main.bundle.js
SRC := $(wildcard js/*.js js/components/*.js styles/*.css)

NODE_MODULES := node_modules

.PHONY: all
all: ${BUILD_DIRS} ${NODE_MODULES} ${MAIN}

${NODE_MODULES}:
	@npm install

${MAIN}: ${SRC} ${PKG_JSON}
	@npm run ${NPM_TGT}

${PKG_JSON}: ${VER_BUILD}
	@${SCRIPTS}/updateVersion.py ${PKG_JSON} ${VERSION}

${VER_BUILD}:
	@touch $@

${BUILD_DIRS}:
	@mkdir -p $@

.PHONY: purge
purge: clean
	@echo "purging frontend"
	@rm -rf ${DIR}/node_modules

.PHONY: clean
clean:
	@echo "cleaning frontend"
	@rm -rf ${BUILD_DIRS}
