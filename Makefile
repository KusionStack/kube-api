SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: fmt
fmt: golangci
	$(GOLANGCI) fmt

.PHONY: lint
lint: fmt
	$(GOLANGCI) run

manifests:
	@bash hack/make-rules/update-manifests.sh

generate:
	@bash hack/make-rules/update-codegen.sh


##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin

GOLANGCI ?= $(LOCALBIN)/golangci-lint

.PHONY: golangci
golangci: 
	@bash hack/make-rules/install-go-tools.sh golangci-lint
