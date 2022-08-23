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

type DatastoreObservation struct {
	Accessible *bool `json:"accessible,omitempty" tf:"accessible,omitempty"`

	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	FreeSpace *float64 `json:"freeSpace,omitempty" tf:"free_space,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MaintenanceMode *string `json:"maintenanceMode,omitempty" tf:"maintenance_mode,omitempty"`

	MultipleHostAccess *bool `json:"multipleHostAccess,omitempty" tf:"multiple_host_access,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	UncommittedSpace *float64 `json:"uncommittedSpace,omitempty" tf:"uncommitted_space,omitempty"`
}

type DatastoreParameters struct {

	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The managed object ID of the datastore cluster to place the datastore in.
	// +kubebuilder:validation:Optional
	DatastoreClusterID *string `json:"datastoreClusterId,omitempty" tf:"datastore_cluster_id,omitempty"`

	// The disks to add to the datastore.
	// +kubebuilder:validation:Required
	Disks []*string `json:"disks" tf:"disks,omitempty"`

	// The path to the datastore folder to put the datastore in.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The managed object ID of the host to set up the datastore on.
	// +kubebuilder:validation:Required
	HostSystemID *string `json:"hostSystemId" tf:"host_system_id,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DatastoreSpec defines the desired state of Datastore
type DatastoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatastoreParameters `json:"forProvider"`
}

// DatastoreStatus defines the observed state of Datastore.
type DatastoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatastoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Datastore is the Schema for the Datastores API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vspherejet}
type Datastore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatastoreSpec   `json:"spec"`
	Status            DatastoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatastoreList contains a list of Datastores
type DatastoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datastore `json:"items"`
}

// Repository type metadata.
var (
	Datastore_Kind             = "Datastore"
	Datastore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Datastore_Kind}.String()
	Datastore_KindAPIVersion   = Datastore_Kind + "." + CRDGroupVersion.String()
	Datastore_GroupVersionKind = CRDGroupVersion.WithKind(Datastore_Kind)
)

func init() {
	SchemeBuilder.Register(&Datastore{}, &DatastoreList{})
}
