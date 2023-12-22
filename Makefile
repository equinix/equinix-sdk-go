.PHONY: all pull fetch patch generate clean codegen mod docs move-other patch-post fmt test stage

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

GIT_ORG=equinix
GIT_REPO=equinix-sdk-go
PACKAGE_VERSION=$(shell cat version)
USER_AGENT=${GIT_REPO}/${PACKAGE_VERSION}

OPENAPI_IMAGE_TAG=v7.2.0
OPENAPI_IMAGE=openapitools/openapi-generator-cli:${OPENAPI_IMAGE_TAG}
CRI=docker # nerdctl
OPENAPI_GENERATOR=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/local ${OPENAPI_IMAGE}
SPEC_FETCHER=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/workdir --entrypoint sh mikefarah/yq:4.30.8 script/download_spec.sh
GOLANGCI_LINT=golangci-lint

SPEC_BASE_DIR=spec/services
TEMPLATE_BASE_DIR=templates/services
CODE_BASE_DIR=services

onboard-service:
	script/onboard_service.sh

generate-all:
	for makefile in $(shell set -x; ls -1 Makefile.* | sort -n); do \
		make -f $$makefile generate;\
	done

mod:
	rm -f go.mod go.sum ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/go.mod ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/go.sum
	go mod init github.com/${GIT_ORG}/${GIT_REPO}
	go mod tidy

test:
	go test -v ./...

lint:
	@$(GOLANGCI_LINT) run -v --no-config --fast=false --fix --disable-all --enable goimports $(PACKAGE_PREFIX)

fmt:
	go run mvdan.cc/gofumpt@v0.3.1 -l -w .

stage:
	test -d .git && git add --intent-to-add .
