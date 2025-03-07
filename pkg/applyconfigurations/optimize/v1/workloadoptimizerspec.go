/*
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	optimizev1 "github.com/thestormforge/agent-go/pkg/api/optimize/v1"
)

// WorkloadOptimizerSpecApplyConfiguration represents a declarative configuration of the WorkloadOptimizerSpec type for use
// with apply.
type WorkloadOptimizerSpecApplyConfiguration struct {
	TargetRef      *TargetRefApplyConfiguration      `json:"workloadTargetRef,omitempty"`
	PatchTargetRef *PatchTargetRefApplyConfiguration `json:"patchTargetRef,omitempty"`
	Workload       *optimizev1.Workload              `json:"workloadSettings,omitempty"`
	Schedule       *string                           `json:"schedule,omitempty"`
	LearningPeriod *string                           `json:"learningPeriod,omitempty"`
	Apply          *ApplyApplyConfiguration          `json:"apply,omitempty"`
	Reliability    *ReliabilityApplyConfiguration    `json:"reliability,omitempty"`
	AutoDeploy     *bool                             `json:"autoDeploy,omitempty"`
	Containers     []ContainerApplyConfiguration     `json:"containerSettings,omitempty"`
	Autoscaler     *AutoscalerApplyConfiguration     `json:"hpaSettings,omitempty"`
	PodScheduling  *PodSchedulingApplyConfiguration  `json:"podScheduling,omitempty"`
	Thresholds     *optimizev1.Thresholds            `json:"thresholds,omitempty"`
}

// WorkloadOptimizerSpecApplyConfiguration constructs a declarative configuration of the WorkloadOptimizerSpec type for use with
// apply.
func WorkloadOptimizerSpec() *WorkloadOptimizerSpecApplyConfiguration {
	return &WorkloadOptimizerSpecApplyConfiguration{}
}

// WithTargetRef sets the TargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetRef field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithTargetRef(value *TargetRefApplyConfiguration) *WorkloadOptimizerSpecApplyConfiguration {
	b.TargetRef = value
	return b
}

// WithPatchTargetRef sets the PatchTargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PatchTargetRef field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithPatchTargetRef(value *PatchTargetRefApplyConfiguration) *WorkloadOptimizerSpecApplyConfiguration {
	b.PatchTargetRef = value
	return b
}

// WithWorkload sets the Workload field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Workload field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithWorkload(value optimizev1.Workload) *WorkloadOptimizerSpecApplyConfiguration {
	b.Workload = &value
	return b
}

// WithSchedule sets the Schedule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Schedule field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithSchedule(value string) *WorkloadOptimizerSpecApplyConfiguration {
	b.Schedule = &value
	return b
}

// WithLearningPeriod sets the LearningPeriod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LearningPeriod field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithLearningPeriod(value string) *WorkloadOptimizerSpecApplyConfiguration {
	b.LearningPeriod = &value
	return b
}

// WithApply sets the Apply field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Apply field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithApply(value *ApplyApplyConfiguration) *WorkloadOptimizerSpecApplyConfiguration {
	b.Apply = value
	return b
}

// WithReliability sets the Reliability field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reliability field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithReliability(value *ReliabilityApplyConfiguration) *WorkloadOptimizerSpecApplyConfiguration {
	b.Reliability = value
	return b
}

// WithAutoDeploy sets the AutoDeploy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AutoDeploy field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithAutoDeploy(value bool) *WorkloadOptimizerSpecApplyConfiguration {
	b.AutoDeploy = &value
	return b
}

// WithContainers adds the given value to the Containers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Containers field.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithContainers(values ...*ContainerApplyConfiguration) *WorkloadOptimizerSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithContainers")
		}
		b.Containers = append(b.Containers, *values[i])
	}
	return b
}

// WithAutoscaler sets the Autoscaler field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Autoscaler field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithAutoscaler(value *AutoscalerApplyConfiguration) *WorkloadOptimizerSpecApplyConfiguration {
	b.Autoscaler = value
	return b
}

// WithPodScheduling sets the PodScheduling field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodScheduling field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithPodScheduling(value *PodSchedulingApplyConfiguration) *WorkloadOptimizerSpecApplyConfiguration {
	b.PodScheduling = value
	return b
}

// WithThresholds sets the Thresholds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Thresholds field is set to the value of the last call.
func (b *WorkloadOptimizerSpecApplyConfiguration) WithThresholds(value optimizev1.Thresholds) *WorkloadOptimizerSpecApplyConfiguration {
	b.Thresholds = &value
	return b
}
