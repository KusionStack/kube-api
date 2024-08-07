LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)
CONTROLLER_TOOLS_VERSION ?= v0.15.0
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
CODEGEN = $(LOCALBIN)/kube-codegen

controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary. If wrong version is installed, it will be overwritten.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen && $(LOCALBIN)/controller-gen --version | grep -q $(CONTROLLER_TOOLS_VERSION) || \
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

manifests: controller-gen
	@for dir in apps cluster; do \
  		mkdir -p "config/crd/$$dir"; \
  	    $(CONTROLLER_GEN) rbac:roleName=manager-role crd:generateEmbeddedObjectMeta=true webhook paths="./$$dir/..." output:crd:artifacts:config="config/crd/$$dir"; \
   	done

generate: codegen controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	#@scripts/generate_client.sh
	# $(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	$(CODEGEN) code-gen --go-header-file=./hack/boilerplate.go.txt --code-generator-version=v0.27.16 --apis-path=./ --generators=deepcopy,register

codegen: $(CODEGEN) # Download kube-codegen locally if necessary. If wrong version is installed, it will be overwritten.
$(CODEGEN):
ifeq (, $(shell command -v $(CODEGEN)))
	@cd $(shell mktemp -d) && \
		git clone https://github.com/zoumo/kube-codegen.git && \
		cd kube-codegen && \
		GOBIN=$(LOCALBIN) go install ./cmd/kube-codegen
endif
