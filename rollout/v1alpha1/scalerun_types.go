// Copyright 2023 The KusionStack Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
// +kubebuilder:resource:shortName=scr
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Batch Index",type="string",JSONPath=".status.batchStatus.currentBatchIndex"
// +kubebuilder:printcolumn:name="Batch State",type="string",JSONPath=".status.batchStatus.currentBatchState"
// +kubebuilder:printcolumn:name="Error",type="string",JSONPath=".status.error.code"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",format="date-time"

type ScaleRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScaleRunSpec   `json:"spec,omitempty"`
	Status ScaleRunStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ScaleRunList contains a list of ScaleRun
type ScaleRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScaleRun `json:"items"`
}

type ScaleRunSpec struct {
	// TargetType defines the GroupVersionKind of target resource
	TargetType ObjectTypeRef `json:"targetType,omitempty"`

	// Scale Batch Strategy
	Batch *ScaleRunBatchStrategy `json:"batch,omitempty"`

	// Rollback Batch Strategy
	Rollback *ScaleRunBatchStrategy `json:"rollback,omitempty"`

	// Webhooks defines scale webhook configuration
	// +optional
	Webhooks []RolloutWebhook `json:"webhooks,omitempty"`
}

type ScaleRunBatchStrategy struct {
	// Batches define the order of phases to execute scale in batch
	Batches []ScaleRunStep `json:"batches,omitempty"`
}

type ScaleRunStep struct {
	// desired target replicas
	Targets []ScaleRunStepTarget `json:"targets"`

	// If set to true, the scale will be paused before the step starts.
	// +optional
	Breakpoint bool `json:"breakpoint,omitempty"`

	// Properties contains additional information for step
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
}

// ScaleRunStepTarget defines the target resource of scale
// Either replicas or multiReplicas must be set, but not both
type ScaleRunStepTarget struct {
	CrossClusterObjectNameReference `json:",inline"`

	// Replicas is the replicas of the scale task, which represents the replicas of the target resource
	Replicas int32 `json:"replicas"`

	// MultipleReplias is the replicas of the scale task, which represents the replicas of different target group resource
	MultipleReplias []MultipleReplia `json:"multiReplicas,omitempty"`
}

type MultipleReplia struct {
	// GroupName is the name of current group
	GroupName string `json:"groupName,omitempty"`
	// Value is the replicas of current group
	Value int32 `json:"value"`
}

type ScaleRunStatus struct {
	// ObservedGeneration is the most recent generation observed for this ScaleRun. It corresponds to the
	// ScaleRun's generation, which is updated on mutation by the API Server.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// Conditions is the list of conditions
	Conditions []Condition `json:"conditions,omitempty"`
	// Phase indecates the current phase of scale
	Phase RolloutRunPhase `json:"phase,omitempty"`
	// The last time this status was updated.
	// +optional
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
	// Error indicates the error info of progressing.
	Error *CodeReasonMessage `json:"error,omitempty"`
	// Batches describes the state of the active scale batch.
	// +optional
	Batches *ScaleRunBatchStatus `json:"batches,omitempty"`
	// Rollback describes the state of the active rollback batch.
	// +optional
	Rollback *ScaleRunBatchStatus `json:"rollback,omitempty"`
}

type ScaleRunBatchStatus struct {
	// RolloutBatchStatus contains status of current batch
	RolloutBatchStatus `json:",inline"`
	// Records contains all batches status details.
	Records []ScaleRunStepStatus `json:"records,omitempty"`
}

type ScaleRunStepStatus struct {
	// Index is the id of the batch
	Index *int32 `json:"index,omitempty"`
	// State is Rollout step state
	State RolloutStepState `json:"state,omitempty"`
	// StartTime is the time when the stage started
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// FinishTime is the time when the stage finished
	// +optional
	FinishTime *metav1.Time `json:"finishTime,omitempty"`
	// Targets contains release details for each workload
	// +optional
	Targets []ScaleWorkloadStatus `json:"targets,omitempty"`
	// Webhooks contains webhook status
	// +optional
	Webhooks []RolloutWebhookStatus `json:"webhooks,omitempty"`
}

type ScaleWorkloadStatus struct {
	// Name is the workload name
	Name string `json:"name,omitempty"`
	// Cluster defines which cluster the workload is in.
	Cluster string `json:"cluster,omitempty"`
	// Replicas is the desired number of pods targeted by workload
	Replicas int32 `json:"replicas"`
	// CurrentReplicas is the number of current existed pods targeted by workload
	CurrentReplicas int32 `json:"currentReplicas"`
	// AvailableReplicas is the available number of pods targeted by workload
	AvailableReplicas int32 `json:"availableReplicas"`
	// ScaleFrom is the number of replicas to scale from.
	ScaleFrom int32 `json:"scaleFrom"`
	// ScaleTo is the number of replicas to scale to.
	// It is the final number of replicas after scale operation.
	ScaleTo int32 `json:"scaleTo"`
}

func (r *ScaleRun) IsCompleted() bool {
	if r == nil {
		return false
	}
	return r.Status.Phase == RolloutRunPhaseSucceeded || r.Status.Phase == RolloutRunPhaseCanceled
}
