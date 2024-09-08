PROJECT_NAME := "golang-clean-architecture"
PKG := "github.com/rakib-09/$(PROJECT_NAME)"

GO ?= $(shell command -v go 2> /dev/null)
PACKAGES=$(shell go list ./...)

TOOLS_BIN_DIR := $(abspath bin)
GO_INSTALL = ./scripts/go_install.sh

GOIMPORTS_VER := master
GOIMPORTS_BIN := goimports
GOIMPORTS := $(TOOLS_BIN_DIR)/$(GOIMPORTS_BIN)

GODOC_VER := master
GODOC_BIN := gomarkdoc
GODOC := $(TOOLS_BIN_DIR)/$(GODOC_BIN)

GOLANGCILINT_VER := v1.56.2
GOLANGCILINT_BIN := golangci-lint
GOLANGCILINT := $(TOOLS_BIN_DIR)/$(GOLANGCILINT_BIN)


.PHONY: all dep test build check goimports goformat prepare docker swag

all: build

run:
	@go run main.go serve

test: ## Run unittests
	@rm -f ${DB_LIST}
	@go clean -testcache
	@go test -p 20 -timeout 1800s -failfast -short ${PACKAGES}

dep: ## Get the dependencies
	@go mod vendor
	@go mod tidy

build: ## Build the binary file
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -a -installsuffix cgo -o binary/app  .
lint: $(GOLANGCILINT)
	@echo Running golangci-lint
	$(GOLANGCILINT) run

govet:
	@echo Running govet
	$(GO) vet ./...
	@echo Govet success