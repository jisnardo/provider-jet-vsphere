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

type ContainerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ContainerParameters struct {

	// Determines if the reservation on a vApp container can grow beyond the specified value, if the parent resource pool has unreserved resources.
	// +kubebuilder:validation:Optional
	CPUExpandable *bool `json:"cpuExpandable,omitempty" tf:"cpu_expandable,omitempty"`

	// The utilization of a vApp container will not exceed this limit, even if there are available resources. Set to -1 for unlimited.
	// +kubebuilder:validation:Optional
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// Amount of CPU (MHz) that is guaranteed available to the vApp container.
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

	// Determines if the reservation on a vApp container can grow beyond the specified value, if the parent resource pool has unreserved resources.
	// +kubebuilder:validation:Optional
	MemoryExpandable *bool `json:"memoryExpandable,omitempty" tf:"memory_expandable,omitempty"`

	// The utilization of a vApp container will not exceed this limit, even if there are available resources. Set to -1 for unlimited.
	// +kubebuilder:validation:Optional
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// Amount of memory (MB) that is guaranteed available to the vApp container.
	// +kubebuilder:validation:Optional
	MemoryReservation *float64 `json:"memoryReservation,omitempty" tf:"memory_reservation,omitempty"`

	// The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of low, normal, high, or custom.
	// +kubebuilder:validation:Optional
	MemoryShareLevel *string `json:"memoryShareLevel,omitempty" tf:"memory_share_level,omitempty"`

	// The number of shares allocated. Used to determine resource allocation in case of resource contention. If this is set, memory_share_level must be custom.
	// +kubebuilder:validation:Optional
	MemoryShares *float64 `json:"memoryShares,omitempty" tf:"memory_shares,omitempty"`

	// The ID of the parent VM folder.
	// +kubebuilder:validation:Optional
	ParentFolderID *string `json:"parentFolderId,omitempty" tf:"parent_folder_id,omitempty"`

	// The managed object ID of the parent resource pool or the compute resource the vApp container is in.
	// +kubebuilder:validation:Required
	ParentResourcePoolID *string `json:"parentResourcePoolId" tf:"parent_resource_pool_id,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ContainerSpec defines the desired state of Container
type ContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerParameters `json:"forProvider"`
}

// ContainerStatus defines the observed state of Container.
type ContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Container is the Schema for the Containers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vspherejet}
type Container struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerSpec   `json:"spec"`
	Status            ContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerList contains a list of Containers
type ContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Container `json:"items"`
}

// Repository type metadata.
var (
	Container_Kind             = "Container"
	Container_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Container_Kind}.String()
	Container_KindAPIVersion   = Container_Kind + "." + CRDGroupVersion.String()
	Container_GroupVersionKind = CRDGroupVersion.WithKind(Container_Kind)
)

func init() {
	SchemeBuilder.Register(&Container{}, &ContainerList{})
}
