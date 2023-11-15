.PHONY: all pull fetch patch generate clean codegen mod docs move-other patch-post fmt test stage

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

SPEC_BASE_URL:=https://api.equinix.com/metal/v1/api-docs
SPEC_ROOT_FILE:=openapi3.yaml
SPEC_FETCHED_DIR=spec/oas3.fetched
SPEC_PATCH_DIR=patches/spec.fetched.json
SPEC_PATCHED_DIR=spec/oas3.patched

SPEC_FETCHED_FILE:=spec.fetched.json
SPEC_PATCHED_FILE:=spec.patched.json
OPENAPI_IMAGE_TAG=v7.1.0
OPENAPI_IMAGE=openapitools/openapi-generator-cli:${OPENAPI_IMAGE_TAG}
GIT_ORG=equinix-labs
GIT_REPO=metal-go
PACKAGE_PREFIX=metal
PACKAGE_MAJOR=v1
PACKAGE_VERSION=$(shell cat version)
CRI=docker # nerdctl

OPENAPI_GENERATOR=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/local ${OPENAPI_IMAGE}
SPEC_FETCHER=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/workdir --entrypoint sh mikefarah/yq:4.30.8 script/download_spec.sh
GOLANGCI_LINT=golangci-lint

all: pull fetch patch generate stage

generate: clean codegen mod docs move-other patch-post fmt test

pull:
	${CRI} pull ${OPENAPI_IMAGE}

fetch:
	${SPEC_FETCHER} ${SPEC_BASE_URL} ${SPEC_FETCHED_DIR} ${SPEC_ROOT_FILE}

patch:
	rm -rf ${SPEC_PATCHED_DIR}
	cp -r ${SPEC_FETCHED_DIR} ${SPEC_PATCHED_DIR}

	for diff in $(shell set -x; find ${SPEC_PATCH_DIR} -name '*.patch' | sort -n); do \
		patch --no-backup-if-mismatch -N -t -p1 -i $$diff; \
	done

patch-post:
	# patch is idempotent, always starting with the generated files
	for diff in $(shell find patches/post -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t -p1 -i $$diff; \
	done

clean:
	rm -rf $(PACKAGE_PREFIX)

codegen:
	${OPENAPI_GENERATOR} generate -g go \
		--package-name ${PACKAGE_PREFIX}${PACKAGE_MAJOR} \
		--http-user-agent "${GIT_REPO}/${PACKAGE_VERSION}" \
		-p packageVersion=${PACKAGE_VERSION} \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO}/services/${PACKAGE_PREFIX}/${PACKAGE_MAJOR} \
		-c /local/config/openapi-generator.json \
		-o /local/ \
		-i /local/${SPEC_PATCHED_DIR}/${SPEC_ROOT_FILE}

validate:
	${OPENAPI_GENERATOR} validate \
		--recommend \
		-i /local/${SPEC_PATCHED_DIR}

mod:
	go mod tidy

test:
	go test -v ./...

clean-docs:
	rm -rf docs API.md

move-docs:
	mv README.md API.md

docs: clean-docs move-docs

move-other:
	rm -rf api .travis.yml git_push.sh

lint:
	@$(GOLANGCI_LINT) run -v --no-config --fast=false --fix --disable-all --enable goimports $(PACKAGE_PREFIX)

fmt:
	go run mvdan.cc/gofumpt@v0.3.1 -l -w .

stage:
	test -d .git && git add --intent-to-add API.md docs ${PACKAGE_PREFIX} go.mod go.sum
