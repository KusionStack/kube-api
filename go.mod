module kusionstack.io/kube-api

go 1.22.0

require (
	k8s.io/api v0.33.2
	k8s.io/apimachinery v0.33.2
	sigs.k8s.io/gateway-api v1.3.0
)

require (
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738 // indirect
	sigs.k8s.io/json v0.0.0-20241010143419-9aa6b5e7a4b3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.7.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.27.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.27.2
	k8s.io/apiserver => k8s.io/apiserver v0.27.2
	k8s.io/client-go => k8s.io/client-go v0.27.2
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.15.1
	sigs.k8s.io/gateway-api => sigs.k8s.io/gateway-api v1.2.0
)
