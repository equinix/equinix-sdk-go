.PHONY: all pull fetch patch generate clean codegen mod docs move-other patch-post fmt test stage

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

GIT_ORG=equinix
GIT_REPO=equinix-sdk-go
PACKAGE_VERSION=$(shell cat version)
USER_AGENT=${GIT_REPO}/${PACKAGE_VERSION}

OPENAPI_IMAGE_TAG=v7.12.0
OPENAPI_IMAGE=openapitools/openapi-generator-cli:${OPENAPI_IMAGE_TAG}
CRI=docker # nerdctl
CRI_COMMAND_BASE=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} $(DOCKER_EXTRA_ARGS)
OPENAPI_GENERATOR=${CRI_COMMAND_BASE} -v $(CURDIR):/workdir -w /workdir ${OPENAPI_IMAGE}
SPEC_FETCHER=${CRI_COMMAND_BASE} -v $(CURDIR):/workdir --entrypoint sh mikefarah/yq:4.30.8 script/download_spec.sh
MIN_GO_VERSION=1.23.0
GO_CMD=${CRI_COMMAND_BASE} -v $(CURDIR):/workdir -w /workdir -e GOCACHE=/tmp/.cache golang:${MIN_GO_VERSION}
GOLANGCI_LINT_VERSION=v2.3.0
GOLANGCI_LINT_IMAGE=golangci/golangci-lint:${GOLANGCI_LINT_VERSION}
GOLANGCI_LINT=${CRI_COMMAND_BASE} -v $(CURDIR):/app -w /app -e GOLANGCI_LINT_CACHE=/tmp/.cache -e GOCACHE=/tmp/.cache ${GOLANGCI_LINT_IMAGE} golangci-lint

SPEC_BASE_DIR=spec/services
TEMPLATE_BASE_DIR=templates/services
CODE_BASE_DIR=services

onboard-service:
	script/onboard_service.sh

patch-all:
	for makefile in $(shell set -x; ls -1 Makefile.* | sort -n); do \
		make -f $$makefile patch;\
		RESULT=$$(($${RESULT} + $$?)); \
	done; \
	exit $$RESULT;

generate-all:
	for makefile in $(shell set -x; ls -1 Makefile.* | sort -n); do \
		make -f $$makefile generate;\
	done

mod:
	for goModOrSum in $(shell set -x; find . -not \( -path ./examples -prune \) -name go.mod -o -name go.sum | sort -n); do \
		rm -f $$goModOrSum;\
	done
	${GO_CMD} go mod init github.com/${GIT_ORG}/${GIT_REPO}
	${GO_CMD} go mod tidy

test:
	${GO_CMD} go test -v ./...

lint:
	${GOLANGCI_LINT} run -v

fmt:
	${GO_CMD} go run mvdan.cc/gofumpt@v0.3.1 -l -w .

stage:
	test -d .git && git add --intent-to-add .
