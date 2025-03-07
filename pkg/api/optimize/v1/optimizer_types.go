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

package v1

import (
	hpav2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkloadOptimizerSpec struct {
	TargetRef      *TargetRef      `json:"workloadTargetRef,omitempty"`
	PatchTargetRef *PatchTargetRef `json:"patchTargetRef,omitempty"`
	Workload       Workload        `json:"workloadSettings,omitempty"`
	// Schedule dictates the frequency of the recommendations.
	// This can be done using one of the following syntaxes:
	// - Cron; 10 min
	// A schedule for every 10 minutes would look like: `schedule: */10 * * * *`.
	// - ISO-8601 Duration ( https://en.wikipedia.org/wiki/ISO_8601#Durations )
	// A schedule for every 1 hour would look like `schedule: PT1H`.\
	// this overrides the schedule under Schedules.
	// we will set this field as optional until we completely deprecate Schedules
	// +optional
	Schedule string `json:"schedule,omitempty"`
	// +optional
	LearningPeriod string         `json:"learningPeriod,omitempty"`
	Apply          *Apply         `json:"apply,omitempty"`
	Reliability    *Reliability   `json:"reliability,omitempty"`
	AutoDeploy     bool           `json:"autoDeploy,omitempty"`
	Containers     []Container    `json:"containerSettings,omitempty"`
	Autoscaler     *Autoscaler    `json:"hpaSettings,omitempty"`
	PodScheduling  *PodScheduling `json:"podScheduling,omitempty"`

	// +optional
	Thresholds Thresholds `json:"thresholds,omitempty"`
}

// PatchTargetRef specifies the kind and name of the object that should be patched.
type PatchTargetRef struct {
	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name" protobuf:"bytes,2,opt,name=name"`
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,3,opt,name=apiVersion"`
}

// TargetRef specifies the kind and name of the workload object.
type TargetRef struct {
	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// TODO lowercase names are painful; do we handle case changes in code?
	// Restricted to [deployments|statefulsets|daemonsets|rollouts] for initial release
	// +kubebuilder:validation:Enum=deployments;statefulsets;daemonsets;rollouts
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name" protobuf:"bytes,2,opt,name=name"`
	// API version of the referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,3,opt,name=apiVersion"`
}

// Goal is used to configure the aggressiveness of the recommendations.
// +kubebuilder:validation:Enum=savings;Savings;balanced;Balanced;reliability;Reliability;
type Goal string

// Workload allows configuring workload settings.
type Workload map[corev1.ResourceName]WorkloadSettings

// WorkloadSettings provides a way to configure settings for the entire workload.
type WorkloadSettings struct {
	// Goal is used to configure the aggressiveness of the recommendations.
	Goal Goal `json:"optimizationGoal,omitempty"`
}

// Apply Method is used to configure how Optimize Live deploys the complete recommendations
// that are generated when the workload’s learning period is over.
type Apply struct {
	// +kubebuilder:validation:Enum=patchworkloadresources;PatchWorkloadResources;dynamicadmissionwebhook;DynamicAdmissionWebhook;
	Method string `json:"method,omitempty"`
}

type Reliability struct {
	Oom *Oom `json:"oom,omitempty"`
}

type Oom struct {
	MemoryBumpUp *MemoryBumpUp `json:"memoryBumpUp,omitempty"`
}

type MemoryBumpUp struct {
	// +optional
	// +kubebuilder:validation:Enum=IfAutoDeployEnabled;ifautodeployenabled;Always;always;Never;never
	ApplyImmediately string `json:"applyImmediately,omitempty"`
	Period           string `json:"period,omitempty"`
	Percent          string `json:"percent,omitempty"`
	Min              string `json:"min,omitempty"`
	Max              string `json:"max,omitempty"`
}

type Thresholds map[corev1.ResourceName]Threshold

type Threshold struct {
	// +optional
	MinPercentChange *int32 `json:"minPercentChange,omitempty"`
	// +optional
	MinUnitChange resource.Quantity `json:"minUnitChange,omitempty"`
}

// Options configures how recommendations produced by the scheduled run
// should be handled.
type ScheduleOptions struct {
	// Deploy controls if recommendations generated on this schedule
	// should be automatically deployed to the cluster.
	Deploy bool `json:"deploy,omitempty"`
}

// Container is used to configure the workload container settings.
// This should allow for the ability to set container specific resource bounds.
type Container struct {
	Name   string            `json:"name"`
	CPU    *ResourceSettings `json:"cpu,omitempty"`
	Memory *ResourceSettings `json:"memory,omitempty"`
}

// Policy controls how the recommendations should be generated.
// **RequestsAndLimits** which will produce both requests and
// limits recommendations for the container. This is the default.
// **RequestsOnly** will only produce recommendations for requests for the container.
// **DoNotOptimize** will exclude the container from having recommendations generated.
// **RequestsRaiseLimitsIfNeeded** will produce recommendations for requests and raise limits if needed. We will never lower a limit.
// +kubebuilder:validation:Enum=requestsandlimits;RequestsAndLimits;requestsonly;RequestsOnly;donotoptimize;DoNotOptimize;requestsraiselimitsifneeded;RequestsRaiseLimitsIfNeeded
type Policy string

// ResourceSettings configures the specific requests and limits for the container.
type ResourceSettings struct {
	// Policy dictates the OptimizationPolicy that should be used.
	Policy   Policy            `json:"optimizationPolicy,omitempty"`
	Requests *RequestsSettings `json:"requests,omitempty"`
	Limits   *LimitsSettings   `json:"limits,omitempty"`
}

// RequestsSettings configures the container requests.
type RequestsSettings struct {
	CommonResource `json:",inline"`
}

// LimitsSettings configures the container limits.
type LimitsSettings struct {
	CommonResource    `json:",inline"`
	LimitRequestRatio resource.Quantity `json:"limitRequestRatio,omitempty"`
}

// CommonResource handles the bounds and patch path settings for limits and requests.
type CommonResource struct {
	// Min specifies the lower bound of the recommendation.
	// If a recommendation is produced outside of the defined bound it will be clipped.
	Min resource.Quantity `json:"min,omitempty"`
	// Max specifies the upper bound of the recommendation.
	// If a recommendation is produced outside of the defined bound it will be clipped.
	Max resource.Quantity `json:"max,omitempty"`
	// PatchPath controls where the recommendation should be applied/patched.
	// The format for patchPath is yamlPath format.
	PatchPath string `json:"patchPath,omitempty"`
	// PatchFormat specifies a Go template used to format the value for the patched resource.
	PatchFormat string `json:"patchFormat,omitempty"`
}

// Autoscaler handles the configuration for the HPA configured for the workload.
type Autoscaler struct {
	// PatchTargetRef allows specifiying the object that should be patched
	// with HPA settings.
	// This is most useful in situations where a CRD is used to configure a HPA.
	// +optional
	PatchTargetRef *PatchTargetRef `json:"patchTargetRef,omitempty"`
	// Metrics configures the HPA metrics the workload scales on.
	Metrics []AutoscalerMetrics `json:"metrics,omitempty"`
}

// AutoscalerMetrics holds the configuration for HPA metrics.
type AutoscalerMetrics struct {
	// Configuration for a Container metric based HPA metric type.
	ContainerResource *AutoscalerContainerResource `json:"containerResource,omitempty"`
	// Configuration for an External metric based HPA metric type.
	External *AutoscalerExternal `json:"external,omitempty"`
	// Configuration for an Object metric based HPA metric type.
	Object *AutoscalerObject `json:"object,omitempty"`
	// Configuration for a Pod metric based HPA metric type.
	Pods *AutoscalerPods `json:"pods,omitempty"`
	// Configuration for a Resource metric based HPA metric type.
	Resource *AutoscalerResource `json:"resource,omitempty"`

	// Target defines the bounds for the recommended value.
	Target *AutoscalerTarget `json:"target,omitempty"`
}

// AutoscalerContainerResource contains the identifiers for ContainerResource HPA metrics.
type AutoscalerContainerResource struct {
	Name      string `json:"name,omitempty"`
	Container string `json:"container,omitempty"`
}

// Use new type here so we can disregard the methods associated with MetricIdentifier
// // the hpav2.MetricIdentifier.String() was giving me grief

// AutoscalerIdentifier contains the identifiers for metrics.
type AutoscalerIdentifier hpav2.MetricIdentifier

// AutoscalerContainerResource contains the identifiers for ContainerResource metric type.
type AutoscalerExternal struct {
	AutoscalerIdentifier `json:",inline"`
}

// AutoscalerObject contains the identifiers for Object metric type.
type AutoscalerObject struct {
	AutoscalerIdentifier `json:",inline"`
	DescribedObject      TargetRef `json:"describedObject,omitempty"`
}

// AutoscalerPods contains the identifiers for the Pod metric type.
type AutoscalerPods struct {
	AutoscalerIdentifier `json:",inline"`
}

// AutoscalerResource contains the identifiers for the Resource metric type.
type AutoscalerResource struct {
	Name string `json:"name,omitempty"`
}

// AutoscalerTarget contains the bounds and patch location for HPA recommendations.
type AutoscalerTarget struct {
	Type      hpav2.MetricTargetType `json:"type,omitempty"`
	Min       int                    `json:"min,omitempty"`
	Max       int                    `json:"max,omitempty"`
	PatchPath string                 `json:"patchPath,omitempty"`
}

// PodSchedulingOptimizationPolicy is used to configure whether to optimize NodeAffinity.
// +kubebuilder:validation:Enum=instancecategories;InstanceCategories;donotoptimize;DoNotOptimize
// +kubebuilder:default:DoNotOptimize
type PodSchedulingOptimizationPolicy string

// PodScheduling handles the configuration for the behavior of optimizing pod scheduling.
type PodScheduling struct {
	SchedulingOptimizationPolicy PodSchedulingOptimizationPolicy  `json:"schedulingOptimizationPolicy,omitempty"`
	InstanceCategories           *PodSchedulingInstanceCategories `json:"instanceCategories,omitempty"`
}

// SchedulingInstanceCategories describes optimization for instance category affinities.
type PodSchedulingInstanceCategories struct {
	PatchPath    string                     `json:"patchPath,omitempty"`
	PatchFormat  string                     `json:"patchFormat,omitempty"`
	NodeAffinity *PodSchedulingNodeAffinity `json:"nodeAffinity,omitempty"`
}

// +kubebuilder:validation:Enum=soft;Soft;preferred;Preferred;preferredDuringSchedulingIgnoredDuringExecution;preferredduringschedulingignoredduringexecution;hard;Hard;required;Required;requiredDuringSchedulingIgnoredDuringExecution;requiredduringschedulingignoredduringexecution
// +kubebuilder:default:preferredDuringSchedulingIgnoredDuringExecution
type PodSchedulingNodeAffinityType string

const PodSchedulingNodeAffinityDefaultWeight int32 = 70

type PodSchedulingNodeAffinity struct {
	Type PodSchedulingNodeAffinityType `json:"type,omitempty"`
	// +optional
	Weight int32 `json:"weight,omitempty" jsonschema:"minimum=1,maximum=100,default=70"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="WorkloadOptimizer state"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:resource:shortName="wo"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadOptimizer is the Schema for the WorkloadOptimizer API.
type WorkloadOptimizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkloadOptimizerSpec   `json:"spec,omitempty"`
	Status WorkloadOptimizerStatus `json:"status,omitempty"`
}

// WorkloadOptimizerState communicates the current state of the WorkloadOptimizer.
// **NotSynchronized** configuration has not been sent to StormForge.
// **Synchronized** configuration has been sent to StormForge.
// **Ignored** configuration will not be sent to StormForge.
// +kubebuilder:validation:Enum=NotSynchronized;Synchronized;Ignored
// +kubebuilder:default:NotSynchronized
type WorkloadOptimizerState string

const (
	StateNotSynchronized WorkloadOptimizerState = "NotSynchronized"
	StateSynchronized    WorkloadOptimizerState = "Synchronized"
	StateIgnored         WorkloadOptimizerState = "Ignored"
)

// WorkloadOptimizerStatus defines the observed state of Optimizer.
type WorkloadOptimizerStatus struct {
	State WorkloadOptimizerState `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadOptimizerList contains a list of WorkloadOptimizers.
type WorkloadOptimizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadOptimizer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkloadOptimizer{}, &WorkloadOptimizerList{})
}
