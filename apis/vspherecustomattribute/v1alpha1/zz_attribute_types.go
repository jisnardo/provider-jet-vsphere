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

type AttributeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AttributeParameters struct {

	// Object type for which the custom attribute is valid. If not specified, the attribute is valid for all managed object types.
	// +kubebuilder:validation:Optional
	ManagedObjectType *string `json:"managedObjectType,omitempty" tf:"managed_object_type,omitempty"`
}

// AttributeSpec defines the desired state of Attribute
type AttributeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttributeParameters `json:"forProvider"`
}

// AttributeStatus defines the observed state of Attribute.
type AttributeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttributeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Attribute is the Schema for the Attributes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vspherejet}
type Attribute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttributeSpec   `json:"spec"`
	Status            AttributeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttributeList contains a list of Attributes
type AttributeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attribute `json:"items"`
}

// Repository type metadata.
var (
	Attribute_Kind             = "Attribute"
	Attribute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attribute_Kind}.String()
	Attribute_KindAPIVersion   = Attribute_Kind + "." + CRDGroupVersion.String()
	Attribute_GroupVersionKind = CRDGroupVersion.WithKind(Attribute_Kind)
)

func init() {
	SchemeBuilder.Register(&Attribute{}, &AttributeList{})
}
