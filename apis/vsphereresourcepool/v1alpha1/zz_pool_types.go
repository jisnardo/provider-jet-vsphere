/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PoolObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PoolParameters struct {

	// Determines if the reservation on a resource pool can grow beyond the specified value, if the parent resource pool has unreserved resources.
	// +kubebuilder:validation:Optional
	CPUExpandable *bool `json:"cpuExpandable,omitempty" tf:"cpu_expandable,omitempty"`

	// The utilization of a resource pool will not exceed this limit, even if there are available resources. Set to -1 for unlimited.
	// +kubebuilder:validation:Optional
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// Amount of CPU (MHz) that is guaranteed available to the resource pool.
	// +kubebuilder:validation:Optional
	CPUReservation *float64 `json:"cpuReservation,omitempty" tf:"cpu_reservation,omitempty"`

	// The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of low, normal, high, or custom.
	// +kubebuilder:validation:Optional
	CPUShareLevel *string `json:"cpuShareLevel,omitempty" tf:"cpu_share_level,omitempty"`

	// The number of shares allocated. Used to determine resource allocation in case of resource contention. If this is set, cpu_share_level must be custom.
	// +kubebuilder:validation:Optional
	CPUShares *float64 `json:"cpuShares,omitempty" tf:"cpu_shares,omitempty"`

	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// Determines if the reservation on a resource pool can grow beyond the specified value, if the parent resource pool has unreserved resources.
	// +kubebuilder:validation:Optional
	MemoryExpandable *bool `json:"memoryExpandable,omitempty" tf:"memory_expandable,omitempty"`

	// The utilization of a resource pool will not exceed this limit, even if there are available resources. Set to -1 for unlimited.
	// +kubebuilder:validation:Optional
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// Amount of memory (MB) that is guaranteed available to the resource pool.
	// +kubebuilder:validation:Optional
	MemoryReservation *float64 `json:"memoryReservation,omitempty" tf:"memory_reservation,omitempty"`

	// The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of low, normal, high, or custom.
	// +kubebuilder:validation:Optional
	MemoryShareLevel *string `json:"memoryShareLevel,omitempty" tf:"memory_share_level,omitempty"`

	// The number of shares allocated. Used to determine resource allocation in case of resource contention. If this is set, memory_share_level must be custom.
	// +kubebuilder:validation:Optional
	MemoryShares *float64 `json:"memoryShares,omitempty" tf:"memory_shares,omitempty"`

	// The ID of the root resource pool of the compute resource the resource pool is in.
	// +kubebuilder:validation:Required
	ParentResourcePoolID *string `json:"parentResourcePoolId" tf:"parent_resource_pool_id,omitempty"`

	// Determines if the shares of all descendants of the resource pool are scaled up or down when the shares of the resource pool are scaled up or down.
	// +kubebuilder:validation:Optional
	ScaleDescendantsShares *string `json:"scaleDescendantsShares,omitempty" tf:"scale_descendants_shares,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PoolSpec defines the desired state of Pool
type PoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolParameters `json:"forProvider"`
}

// PoolStatus defines the observed state of Pool.
type PoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pool is the Schema for the Pools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vspherejet}
type Pool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PoolSpec   `json:"spec"`
	Status            PoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolList contains a list of Pools
type PoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pool `json:"items"`
}

// Repository type metadata.
var (
	Pool_Kind             = "Pool"
	Pool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pool_Kind}.String()
	Pool_KindAPIVersion   = Pool_Kind + "." + CRDGroupVersion.String()
	Pool_GroupVersionKind = CRDGroupVersion.WithKind(Pool_Kind)
)

func init() {
	SchemeBuilder.Register(&Pool{}, &PoolList{})
}
