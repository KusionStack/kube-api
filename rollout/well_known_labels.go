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

package rollout

const (
	// This label is added to objects to reference their controller resource.
	LabelControlledBy = "rollout.kusionstack.io/controlled-by"
	// This label is added to workload object to identify the workload type.
	LabelWorkload = "rollout.kusionstack.io/workload"
)

// canary labels
const (
	// This label will be set on objects to indicate that they are temporary resources.
	CanaryResourceLabelKey = "rollout.kusionstack.io/canary-resource"

	// This label indicates the revision of pods controlled by workload.
	TrafficLaneLabelKey = "rollout.kusionstack.io/traffic-lane"
	StableTrafficLane   = "stable"
	CanaryTrafficLane   = "canary"
	UnknownTrafficLane  = "__unknown__"
)

// rollout class label
const (
	LabelRolloutClass = "rollout.kusionstack.io/rollout-class"
)
