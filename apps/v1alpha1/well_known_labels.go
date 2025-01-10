/*
Copyright 2023 The KusionStack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

const (
	ControlledByKusionStackLabelKey = "kusionstack.io/control"
)

// PodOpsLifecycle labels
const (
	// PodOperatingLabelPrefix indicates a pod is operating
	PodOperatingLabelPrefix = "operating.podopslifecycle.kusionstack.io"
	// PodOperationTypeLabelPrefix indicates the type of operation
	PodOperationTypeLabelPrefix = "operation-type.podopslifecycle.kusionstack.io"
	// PodOperationPermissionLabelPrefix indicates the permission of operation
	PodOperationPermissionLabelPrefix = "operation-permission.podopslifecycle.kusionstack.io"
	// PodUndoOperationTypeLabelPrefix indicates the type of operation has been canceled
	PodUndoOperationTypeLabelPrefix = "undo-operation-type.podopslifecycle.kusionstack.io"
	// PodDoneOperationTypeLabelPrefix indicates the type of operation has been done
	PodDoneOperationTypeLabelPrefix = "done-operation-type.podopslifecycle.kusionstack.io"

	// PodPreCheckLabelPrefix indicates a pod is in pre-check phase
	PodPreCheckLabelPrefix = "pre-check.podopslifecycle.kusionstack.io"
	// PodPreCheckedLabelPrefix indicates a pod has finished pre-check phase
	PodPreCheckedLabelPrefix = "pre-checked.podopslifecycle.kusionstack.io"
	// PodPreparingLabelPrefix indicates a pod is preparing for operation
	PodPreparingLabelPrefix = "preparing.podopslifecycle.kusionstack.io"
	// PodOperateLabelPrefix indicates a pod is in operate phase
	PodOperateLabelPrefix = "operate.podopslifecycle.kusionstack.io"
	// PodOperatedLabelPrefix indicates a pod has finished operate phase
	PodOperatedLabelPrefix = "operated.podopslifecycle.kusionstack.io"
	// PodPostCheckLabelPrefix indicates a pod is in post-check phase
	PodPostCheckLabelPrefix = "post-check.podopslifecycle.kusionstack.io"
	// PodPostCheckedLabelPrefix indicates a pod has finished post-check phase
	PodPostCheckedLabelPrefix = "post-checked.podopslifecycle.kusionstack.io"
	// PodCompletingLabelPrefix indicates a pod is completing operation
	PodCompletingLabelPrefix = "completing.podopslifecycle.kusionstack.io"

	// PodServiceAvailableLabel indicates a pod is available to serve
	PodServiceAvailableLabel = "podopslifecycle.kusionstack.io/service-available"
	PodPreCheckLabel         = "podopslifecycle.kusionstack.io/pre-checking"
	PodPreparingLabel        = "podopslifecycle.kusionstack.io/preparing"
	PodOperatingLabel        = "podopslifecycle.kusionstack.io/operating"
	PodPostCheckLabel        = "podopslifecycle.kusionstack.io/post-checking"
	PodCompletingLabel       = "podopslifecycle.kusionstack.io/completing"
	PodCreatingLabel         = "podopslifecycle.kusionstack.io/creating"

	// PodStayOfflineLabel indicates a pod is not ready and available to serve
	PodStayOfflineLabel     = "podopslifecycle.kusionstack.io/stay-offline"
	PodPreparingDeleteLabel = "podopslifecycle.kusionstack.io/preparing-to-delete"
)

var WellKnownLabelPrefixesWithID = []string{
	PodOperatingLabelPrefix, PodOperationTypeLabelPrefix, PodPreCheckLabelPrefix, PodPreCheckedLabelPrefix,
	PodPreparingLabelPrefix, PodDoneOperationTypeLabelPrefix, PodUndoOperationTypeLabelPrefix, PodOperateLabelPrefix, PodOperatedLabelPrefix, PodPostCheckLabelPrefix,
	PodPostCheckedLabelPrefix, PodCompletingLabelPrefix,
}

// CollaSet labels
const (
	// PodInstanceIDLabelKey is used to attach pod instance ID on pod
	PodInstanceIDLabelKey = "collaset.kusionstack.io/instance-id"
	// CollaSetUpdateIndicateLabelKey is used to indicate a pod should be updated by label
	CollaSetUpdateIndicateLabelKey = "collaset.kusionstack.io/update-included"

	// PodDeletionIndicationLabelKey indicates a pod will be deleted by collaset
	PodDeletionIndicationLabelKey = "collaset.kusionstack.io/to-delete"
	// PodReplaceIndicationLabelKey indicates a pod will be replaced by collaset
	PodReplaceIndicationLabelKey = "collaset.kusionstack.io/to-replace"
	// PodReplaceByReplaceUpdateLabelKey indicates a pod is replaced by update by collaset
	PodReplaceByReplaceUpdateLabelKey = "collaset.kusionstack.io/replaced-by-replace-update"
	// PodExcludeIndicationLabelKey indicates a pod will be excluded by collaset
	PodExcludeIndicationLabelKey = "collaset.kusionstack.io/to-exclude"

	// PodReplacePairOriginName is used to indicate replace origin pod name on the new created pod
	PodReplacePairOriginName = "collaset.kusionstack.io/replace-pair-origin-name"
	// PodReplacePairNewId is used to indicate the new created pod instance on replace origin pod
	PodReplacePairNewId = "collaset.kusionstack.io/replace-pair-new-id"

	// PvcTemplateHashLabelKey is used to attach hash of pvc template to pvc resource
	PvcTemplateHashLabelKey = "collaset.kusionstack.io/pvc-template-hash"

	// PodOrphanedIndicateLabelKey indicates pod or pvc is orphaned
	PodOrphanedIndicateLabelKey = "collaset.kusionstack.io/orphaned"
)

// PodDecoration labels
const (
	PodDecorationLabelPrefix = "poddecoration.kusionstack.io/pd-"
)

const (
	SwarmNameLabelKey      = "swarm.apps.kusionstack.io/name"
	SwarmGroupNameLabelKey = "swarm.apps.kusionstack.io/group-name"
)
