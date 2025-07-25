// Copyright 2023 The KusionStack Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ttopo
// +kubebuilder:printcolumn:name="TYPE",type="string",JSONPath=".spec.trafficType"
// +kubebuilder:printcolumn:name="SERVICE",type="string",JSONPath=".spec.backend.name"
// +kubebuilder:printcolumn:name="Routes",type="string",JSONPath=".spec.routes[*].name"
// +kubebuilder:printcolumn:name="BACKEND_ROUTINGS",type="string",JSONPath=".status.topologies[*].backendRoutingName"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",format="date-time"

// TrafficTopologies defines the networking traffic relationships between
// workloads, backend services, and routes.
type TrafficTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficTopologySpec   `json:"spec,omitempty"`
	Status TrafficTopologyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TrafficTopologyList is a list of TrafficTopology resources.
type TrafficTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TrafficTopology `json:"items"`
}

// TrafficTopologySpec is the spec for a TrafficTopology resource.
type TrafficTopologySpec struct {
	// WorkloadRef is the reference to a kind of workloads
	WorkloadRef WorkloadRef `json:"workloadRef"`

	// TrafficType defines the type of traffic
	TrafficType TrafficType `json:"trafficType"`

	// Backend defines the reference to a kind of backend
	Backend BackendRef `json:"backend"`

	// Routes defines the list of routes
	Routes []RouteRef `json:"routes,omitempty"`
}

type TrafficType string

const (
	MultiClusterTrafficType TrafficType = "MultiCluster"
	InClusterTrafficType    TrafficType = "InCluster"

	TrafficTopologyConditionReady ConditionType = "Ready"
)

type BackendRef struct {
	// Group is the group of the referent. For example, "gateway.networking.k8s.io".
	// When unspecified or empty string, core API group is inferred.
	//
	// +optional
	// +kubebuilder:default="v1"
	APIVersion *string `json:"apiVersion,omitempty"`

	// Kind is the Kubernetes resource kind of the referent. For example
	// "Service".
	//
	// Defaults to "Service" when not specified.
	//
	// ExternalName services can refer to CNAME DNS records that may live
	// outside of the cluster and as such are difficult to reason about in
	// terms of conformance. They also may not be safe to forward to (see
	// CVE-2021-25740 for more information). Implementations SHOULD NOT
	// support ExternalName Services.
	//
	// Support: Core (Services with a type other than ExternalName)
	//
	// Support: Implementation-specific (Services with type ExternalName)
	//
	// +optional
	// +kubebuilder:default=Service
	Kind *string `json:"kind,omitempty"`

	// Name is the name of the referent.
	Name string `json:"name"`
}

type RouteRef struct {
	// APIVersion is the group/version of the referent. For example, "gateway.networking.k8s.io/v1".
	//
	// Defaults to "gateway.networking.k8s.io/v1" when not specified.
	//
	// +optional
	// +kubebuilder:default="gateway.networking.k8s.io/v1"
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is the Kubernetes resource kind of the referent. For example
	// "HTTPRoute".
	//
	// Defaults to "HTTPRoute" when not specified.
	//
	// +optional
	// +kubebuilder:default=HTTPRoute
	Kind *string `json:"kind,omitempty"`
	// Name is the name of the custom route.
	Name string `json:"name"`
}

type TrafficTopologyStatus struct {
	// ObservedGeneration is the most recent generation observed.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// Conditions is the list of conditions
	Conditions []Condition `json:"conditions,omitempty"`
	// Topologies information aggregated by workload
	Topologies []TopologyInfo `json:"topologies,omitempty"`
}

type TopologyInfo struct {
	// workload reference name and cluster
	WorkloadRef CrossClusterObjectNameReference `json:"workloadRef,omitempty"`
	// backend routing reference
	// The name of the backendRouting referent
	BackendRoutingName string `json:"backendRoutingName,omitempty"`
}

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=bkr
// +kubebuilder:printcolumn:name="TYPE",type="string",JSONPath=".spec.trafficType"
// +kubebuilder:printcolumn:name="BACKEND",type="string",JSONPath=".spec.backend.name"
// +kubebuilder:printcolumn:name="ROUTES",type="string",JSONPath=".spec.routes[*].name"
// +kubebuilder:printcolumn:name="STABLE",type="string",JSONPath=".status.backends.stable.name"
// +kubebuilder:printcolumn:name="CANARY",type="string",JSONPath=".status.backends.canary.name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",format="date-time"

// BackendRouting defines defines the association between frontend routes and
// backend service, and it allows the user to define forwarding rules for canary scenario.
type BackendRouting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackendRoutingSpec   `json:"spec,omitempty"`
	Status BackendRoutingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BackendRoutingList is a list of BackendRouting resources.
type BackendRoutingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []BackendRouting `json:"items"`
}

type BackendRoutingSpec struct {
	// TrafficType defines the type of traffic
	TrafficType TrafficType `json:"trafficType"`
	// Backend defines the reference to a kind of backend
	Backend CrossClusterObjectReference `json:"backend"`
	// Routes defines the list of routes
	Routes []CrossClusterObjectReference `json:"routes,omitempty"`
	// ForkedBackends
	ForkedBackends *ForkedBackends `json:"forkedBackends,omitempty"`
	// Forwarding defines the forwarding rules for canary scenario
	Forwarding *BackendForwarding `json:"forwarding,omitempty"`
}

type ForkedBackends struct {
	// the temporary stable backend service name, generally it is the {originServiceName}-stable
	Stable ForkedBackend `json:"stable"`
	// the temporary canary backend service name, generally it is the {originServiceName}-canary
	Canary ForkedBackend `json:"canary"`
}

type ForkedBackend struct {
	// the temporary backend name
	Name string `json:"name"`
	// ExtraLabelSelector defines the extra label selector for the temporary backend to select specific pods
	ExtraLabelSelector map[string]string `json:"extraLabelSelector,omitempty"`
}

type BackendForwarding struct {
	HTTP *HTTPForwarding `json:"http,omitempty"`
}

type HTTPForwarding struct {
	Origin *OriginHTTPForwarding `json:"origin,omitempty"`
	Stable *StableHTTPForwarding `json:"stable,omitempty"`
	Canary *CanaryHTTPForwarding `json:"canary,omitempty"`
}

type OriginHTTPForwarding struct {
	BackendName string `json:"backendName,omitempty"`
}

type StableHTTPForwarding struct {
	BackendName string `json:"backendName,omitempty"`
	// stable traffic rule
	HTTPRouteRule `json:",inline"`
}

type CanaryHTTPForwarding struct {
	BackendName string `json:"backendName,omitempty"`
	// Canary traffic rule
	CanaryHTTPRouteRule `json:",inline"`
}

type TrafficStrategy struct {
	HTTP *HTTPTrafficStrategy `json:"http,omitempty"`
}

type HTTPTrafficStrategy struct {
	CanaryHTTPRouteRule `json:",inline"`
	// StableTraffic indicate the base traffic rule
	StableTraffic *HTTPRouteRule `json:"stableTraffic,omitempty"`
}

type BackendRoutingStatus struct {
	// ObservedGeneration is the most recent generation observed.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// Conditions is the list of conditions
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Phase indicates the current phase of this object.
	// Phase BackendRoutingPhase `json:"phase,omitempty"`
	// current backends routing
	Backends BackendStatuses `json:"backends,omitempty"`
	// route statuses
	Routes []BackendRouteStatus `json:"routes,omitempty"`
}

type BackendStatuses struct {
	// Origin backend status
	Origin BackendStatus `json:"origin,omitempty"`
	// Stable backend status
	Stable BackendStatus `json:"stable,omitempty"`
	// Canary backend status
	Canary BackendStatus `json:"canary,omitempty"`
}

type BackendStatus struct {
	// Name is the name of the referent.
	Name string `json:"name"`
	// Conditions represents the current condition of an backend.
	Conditions BackendConditions `json:"conditions,omitempty"`
}

// Backendonditions represents the current condition of an backend.
type BackendConditions struct {
	// ready indicates that this endpoint is prepared to receive traffic,
	// according to whatever system is managing the endpoint. A nil value
	// indicates an unknown state. In most cases consumers should interpret this
	// unknown state as ready. For compatibility reasons, ready should never be
	// "true" for terminating endpoints.
	// +optional
	Ready *bool `json:"ready,omitempty" protobuf:"bytes,1,name=ready"`

	// terminating indicates that this endpoint is terminating. A nil value
	// indicates an unknown state. Consumers should interpret this unknown state
	// to mean that the endpoint is not terminating.
	// +optional
	Terminating *bool `json:"terminating,omitempty" protobuf:"bytes,3,name=terminating"`
}

// condition type
const (
	BackendRoutingReady        string = "Ready"
	BackendRoutingBackendReady string = "BackendReady"
	BackendRoutingRouteReady   string = "RouteReady"
)

// BackendRouteStatus defines the status of a backend route.
type BackendRouteStatus struct {
	// CrossClusterObjectReference defines the reference to a kind of route resource.
	CrossClusterObjectReference `json:",inline"`
	// Forwarding statuses
	// +optional
	Forwarding *BackendRouteForwardingStatuses `json:"forwarding,omitempty"`
	// Route condition
	// +optional
	Condition *metav1.Condition `json:"condition,omitempty"`
}

type BackendRouteForwardingStatuses struct {
	Origin *BackendRouteForwardingStatus `json:"origin,omitempty"`
	Stable *BackendRouteForwardingStatus `json:"stable,omitempty"`
	Canary *BackendRouteForwardingStatus `json:"canary,omitempty"`
}

type BackendRouteForwardingStatus struct {
	BackendName string            `json:"backendName"`
	Conditions  BackendConditions `json:"conditions,omitempty"`
}
