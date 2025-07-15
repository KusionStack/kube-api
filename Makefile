SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

manifests:
	@bash hack/make-rules/update-manifests.sh

generate:
	@bash hack/make-rules/update-codegen.sh

