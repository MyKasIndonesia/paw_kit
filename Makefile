PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: build test race msan

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

all: build test

build: ## Build each packages
	@go build -race ${PKG_LIST}

test: ## Run unittests
	@go test -short -cover ${PKG_LIST}

race: ## Run data race detector
	@go test -race -short -cover ${PKG_LIST}

msan: ## Run memory sanitizer
	@go test -msan -short -cover ${PKG_LIST}

lint: ## Run golang ci linter
	@golangci-lint run