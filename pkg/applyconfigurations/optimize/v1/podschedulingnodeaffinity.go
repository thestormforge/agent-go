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

// PodSchedulingNodeAffinityApplyConfiguration represents a declarative configuration of the PodSchedulingNodeAffinity type for use
// with apply.
type PodSchedulingNodeAffinityApplyConfiguration struct {
	Type   *optimizev1.PodSchedulingNodeAffinityType `json:"type,omitempty"`
	Weight *int32                                    `json:"weight,omitempty"`
}

// PodSchedulingNodeAffinityApplyConfiguration constructs a declarative configuration of the PodSchedulingNodeAffinity type for use with
// apply.
func PodSchedulingNodeAffinity() *PodSchedulingNodeAffinityApplyConfiguration {
	return &PodSchedulingNodeAffinityApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *PodSchedulingNodeAffinityApplyConfiguration) WithType(value optimizev1.PodSchedulingNodeAffinityType) *PodSchedulingNodeAffinityApplyConfiguration {
	b.Type = &value
	return b
}

// WithWeight sets the Weight field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Weight field is set to the value of the last call.
func (b *PodSchedulingNodeAffinityApplyConfiguration) WithWeight(value int32) *PodSchedulingNodeAffinityApplyConfiguration {
	b.Weight = &value
	return b
}
