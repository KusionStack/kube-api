package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// Swarm is the Schema for the multi pod groups API
type Swarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SwarmSpec   `json:"spec,omitempty"`
	Status SwarmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SwarmList contains a list of Swarm
type SwarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Swarm `json:"items"`
}

type SwarmSpec struct {
	// Indicate the number of histories to be conserved
	// If unspecified, defaults to 5
	//
	// +kubebuilder:default=5
	// +optional
	HistoryLimit int32               `json:"historyLimit,omitempty"`
	PodGroups    []SwarmPodGroupSpec `json:"podGroups,omitempty"`
}

type SwarmPodGroupSpec struct {
	// Name must be unique within a swarm.
	Name string `json:"name"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Replicas is the desired number of replicas of the given template.
	Replicas *int32 `json:"replicas,omitempty"`

	// Template defines the pods that will be created from this pod template.
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Template corev1.PodTemplateSpec `json:"template"`

	// VolumeClaimTemplates is a list of claims that pods are allowed to reference.
	// The StatefulSet controller is responsible for mapping network identities to
	// claims in a way that maintains the identity of a pod. Every claim in
	// this list must have at least one matching (by name) volumeMount in one
	// container in the template. A claim in this list takes precedence over
	// any volumes in the template, with the same name.
	// +optional
	VolumeClaimTemplates []corev1.PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"`

	// +optional
	TopologyAware *SwarmTopologyAware `json:"topologyAware,omitempty"`
}

type SwarmTopologyAware struct {
	// DependsOn is a list of pod group names that this pod group depends on.
	DependsOn []string `json:"dependsOn,omitempty"`

	// TopologyInjection defines the policy for injecting pod IPs and FQDNs into
	// this pod group template.
	TopologyInjection *SwarmTopologyInjection `json:"topologyInjection,omitempty"`
}

type SwarmTopologyInjection struct {
	PodIPs   bool `json:"podIPs,omitempty"`
	PodFQDNs bool `json:"podFQDNs,omitempty"`
	// ServiceDomain bool `json:"serviceDomain,omitempty"`
}

type SwarmStatus struct {
	// ObservedGeneration is the most recent generation observed for this Object. It corresponds to the
	// Object's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Count of hash collisions for the Object. The Object controller
	// uses this field as a collision avoidance mechanism when it needs to
	// create the name for the newest ControllerRevision.
	// +optional
	CollisionCount *int32 `json:"collisionCount,omitempty"`

	// CurrentRevision, if not empty, indicates the version of the Object.
	// +optional
	CurrentRevision string `json:"currentRevision,omitempty"`

	// UpdatedRevision, if not empty, indicates the version of the Object currently updated.
	// +optional
	UpdatedRevision string `json:"updatedRevision,omitempty"`

	Conditions       []metav1.Condition    `json:"conditions,omitempty"`
	PodGroupStatuses []SwarmPodGroupStatus `json:"podGroupStatuses,omitempty"`
}

type SwarmPodGroupStatus struct {
	// Name is the name of the pod group.
	Name string `json:"name,omitempty"`

	// Replicas is the most recently observed number of replicas.
	Replicas int32 `json:"replicas,omitempty"`

	// The number of pods in updated version.
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`

	// ReadyReplicas indicates the number of the pod with ready condition
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// The number of available replicas (ready for at least minReadySeconds) for this replica set.
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`
}

type SwarmTopologyAwareInjection struct {
	Groups []SwarmTopologyAwareInjectionGroup `json:"podGroups,omitempty"`
}

type SwarmTopologyAwareInjectionGroup struct {
	GroupName        string            `json:"groupName"`
	Replicas         int32             `json:"replicas"`
	Labels           map[string]string `json:"labels,omitempty"`
	Annotations      map[string]string `json:"annotations,omitempty"`
	PodAddresses     []string          `json:"podAddresses,omitempty"`
	ServiceAddresses []string          `json:"serviceAddresses,omitempty"`
}
