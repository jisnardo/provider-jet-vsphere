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

type DatacenterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`
}

type DatacenterParameters struct {

	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DatacenterSpec defines the desired state of Datacenter
type DatacenterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatacenterParameters `json:"forProvider"`
}

// DatacenterStatus defines the observed state of Datacenter.
type DatacenterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatacenterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Datacenter is the Schema for the Datacenters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vspherejet}
type Datacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatacenterSpec   `json:"spec"`
	Status            DatacenterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatacenterList contains a list of Datacenters
type DatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datacenter `json:"items"`
}

// Repository type metadata.
var (
	Datacenter_Kind             = "Datacenter"
	Datacenter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Datacenter_Kind}.String()
	Datacenter_KindAPIVersion   = Datacenter_Kind + "." + CRDGroupVersion.String()
	Datacenter_GroupVersionKind = CRDGroupVersion.WithKind(Datacenter_Kind)
)

func init() {
	SchemeBuilder.Register(&Datacenter{}, &DatacenterList{})
}
