
CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

GIT_ORG=equinix
GIT_REPO=equinix-sdk-go
OPENAPI_IMAGE_TAG=v7.1.0
OPENAPI_IMAGE=openapitools/openapi-generator-cli:${OPENAPI_IMAGE_TAG}
PACKAGE_VERSION=$(shell cat version)
CRI=docker # nerdctl

mod:
	find services -regex ".*\/go\.mod" -exec rm {} \;
	find services -regex ".*\/go\.sum" -exec rm {} \;
	go mod tidy
