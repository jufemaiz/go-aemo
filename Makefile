# CONSTANT ENVS
FILES		  ?= $(shell find . -type f -name '*.go' -not -path "./proto/*")
PACKAGES	?= $(shell go list ./...)

# DIRECTORIES
BINARY_DIR=bin
CMD_DIR=cmd
FUNCTION_DIR=functions
ZIP_DIR=zip

# Module Root
MODULE_ROOT=github.com/enosi/billing-adjustment

# VERSIONING
# Make use of lazy evaluation ala http://cakoose.com/wiki/gnu_make_thunks
VERSION_GEN=$(shell git fetch --tags && git describe --tags --dirty --always)
VERSION?=$(eval VERSION := $(VERSION_GEN))$(VERSION)
BUILD_TIME_GEN=$(shell date +%FT%T%z)
BUILD_TIME?=$(eval BUILD_TIME := $(BUILD_TIME_GEN))$(BUILD_TIME)
BUILD_VERSION=-ldflags '-X $(MODULE_ROOT)/pkg/version.Version=$(VERSION) -X $(MODULE_ROOT)/pkg/version.BuildTime=$(BUILD_TIME)'

.PHONY: help
default: help

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean: ## go clean
	go clean -i ./...
	rm -rf $(BINARY_DIR)
	mkdir -p $(BINARY_DIR)
	rm -rf $(ZIP_DIR)
	mkdir -p $(ZIP_DIR)

.PHONY: clean-goimports
clean-goimports: ## needed to deal with <https://github.com/facebook/ent/issues/383>
	rm -f go.sum
	rm -rf vendor
	go clean -modcache

.PHONY: clean-all
clean-all: clean

.PHONY: dependencies
dependencies:
	go mod tidy
	go mod verify

.PHONY: dependencies-download
dependencies-download:
	go mod download

.PHONY: fmt
fmt: ## format the go source files
	go fmt ./...
	goimports -local "${MODULE_ROOT}" -w $(FILES)

.PHONY: generate
generate: ## Generate boilerplate and mock code
	go generate ./...

	# MANUALLY run the following on CLI, in the folder where the mock should be generated, to create a mock implementation
	# moq -out s3_uploader_mock.go -pkg PACKAGE_NAME $(go list -f '{{.Dir}}' github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface) UploaderAPI
	# moq -out sns_mock.go -pkg PACKAGE_NAME $(go list -f '{{.Dir}}' github.com/aws/aws-sdk-go/service/sns/snsiface) SNSAPI
	# moq -out sqs_mock.go -pkg PACKAGE_NAME $(go list -f '{{.Dir}}' github.com/aws/aws-sdk-go/service/sqs/sqsiface) SQSAPI

.PHONY: help
help: ## Show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

.PHONY: lint
lint: ## run go lint on the source files
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed. Please see https://github.com/golangci/golangci-lint#install for installation instructions."; \
		exit 1; \
	fi; \
	golangci-lint run ./... --config ./.golangci.yml --timeout 60m --max-issues-per-linter 50 --max-same-issues 50 --exclude-use-default=false --verbose

.PHONY: lint-verbose
lint-verbose: ## run go lint on the source files
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed. Please see https://github.com/golangci/golangci-lint#install for installation instructions."; \
		exit 1; \
	fi; \
	golangci-lint run ./... --config ./.golangci.yml --timeout 60m --max-issues-per-linter 5000 --max-same-issues 5000 --exclude-use-default=false --verbose

.PHONY: list
list:
	go list ./...

.PHONY: test
test: ## run short tests
	mkdir -p .coverage
	go test -race -v ./... -timeout 60m -short -coverprofile .coverage/coverage.out
	go tool cover -func .coverage/coverage.out | grep ^total:
	go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html

.PHONY: tools tools-update
tools: ## fetch and install all required tools
	go get golang.org/x/tools/cmd/goimports
	go get github.com/smartystreets/goconvey
	go get github.com/matryer/moq
tools-update: ## fetch, update and install all required tools
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/smartystreets/goconvey
	go get -u github.com/matryer/moq

.PHONY: vet
vet: ## run go vet on the source files
	go vet ./...
