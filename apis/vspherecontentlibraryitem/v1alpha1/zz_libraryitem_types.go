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

type LibraryItemObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LibraryItemParameters struct {

	// Optional description of the content library item.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of source VM of content library item.
	// +kubebuilder:validation:Optional
	FileURL *string `json:"fileUrl,omitempty" tf:"file_url,omitempty"`

	// ID of the content library to contain item
	// +kubebuilder:validation:Required
	LibraryID *string `json:"libraryId" tf:"library_id,omitempty"`

	// The managed object ID of an existing VM to be cloned to the content library.
	// +kubebuilder:validation:Optional
	SourceUUID *string `json:"sourceUuid,omitempty" tf:"source_uuid,omitempty"`

	// Type of content library item.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// LibraryItemSpec defines the desired state of LibraryItem
type LibraryItemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LibraryItemParameters `json:"forProvider"`
}

// LibraryItemStatus defines the observed state of LibraryItem.
type LibraryItemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LibraryItemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LibraryItem is the Schema for the LibraryItems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vspherejet}
type LibraryItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LibraryItemSpec   `json:"spec"`
	Status            LibraryItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LibraryItemList contains a list of LibraryItems
type LibraryItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LibraryItem `json:"items"`
}

// Repository type metadata.
var (
	LibraryItem_Kind             = "LibraryItem"
	LibraryItem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LibraryItem_Kind}.String()
	LibraryItem_KindAPIVersion   = LibraryItem_Kind + "." + CRDGroupVersion.String()
	LibraryItem_GroupVersionKind = CRDGroupVersion.WithKind(LibraryItem_Kind)
)

func init() {
	SchemeBuilder.Register(&LibraryItem{}, &LibraryItemList{})
}
