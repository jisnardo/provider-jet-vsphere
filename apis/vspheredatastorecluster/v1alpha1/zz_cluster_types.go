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

type ClusterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClusterParameters struct {

	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The managed object ID of the datacenter to put the datastore cluster in.
	// +kubebuilder:validation:Required
	DatacenterID *string `json:"datacenterId" tf:"datacenter_id,omitempty"`

	// The name of the folder to locate the datastore cluster in.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// Advanced configuration options for storage DRS.
	// +kubebuilder:validation:Optional
	SdrsAdvancedOptions map[string]*string `json:"sdrsAdvancedOptions,omitempty" tf:"sdrs_advanced_options,omitempty"`

	// The default automation level for all virtual machines in this storage cluster.
	// +kubebuilder:validation:Optional
	SdrsAutomationLevel *string `json:"sdrsAutomationLevel,omitempty" tf:"sdrs_automation_level,omitempty"`

	// When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
	// +kubebuilder:validation:Optional
	SdrsDefaultIntraVMAffinity *bool `json:"sdrsDefaultIntraVmAffinity,omitempty" tf:"sdrs_default_intra_vm_affinity,omitempty"`

	// Enable storage DRS for this datastore cluster.
	// +kubebuilder:validation:Optional
	SdrsEnabled *bool `json:"sdrsEnabled,omitempty" tf:"sdrs_enabled,omitempty"`

	// The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	// +kubebuilder:validation:Optional
	SdrsFreeSpaceThreshold *float64 `json:"sdrsFreeSpaceThreshold,omitempty" tf:"sdrs_free_space_threshold,omitempty"`

	// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to freeSpace, drs_free_space_threshold is used.
	// +kubebuilder:validation:Optional
	SdrsFreeSpaceThresholdMode *string `json:"sdrsFreeSpaceThresholdMode,omitempty" tf:"sdrs_free_space_threshold_mode,omitempty"`

	// The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to balance the space.
	// +kubebuilder:validation:Optional
	SdrsFreeSpaceUtilizationDifference *float64 `json:"sdrsFreeSpaceUtilizationDifference,omitempty" tf:"sdrs_free_space_utilization_difference,omitempty"`

	// Overrides the default automation settings when correcting I/O load imbalances.
	// +kubebuilder:validation:Optional
	SdrsIoBalanceAutomationLevel *string `json:"sdrsIoBalanceAutomationLevel,omitempty" tf:"sdrs_io_balance_automation_level,omitempty"`

	// The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this datastore.
	// +kubebuilder:validation:Optional
	SdrsIoLatencyThreshold *float64 `json:"sdrsIoLatencyThreshold,omitempty" tf:"sdrs_io_latency_threshold,omitempty"`

	// Enable I/O load balancing for this datastore cluster.
	// +kubebuilder:validation:Optional
	SdrsIoLoadBalanceEnabled *bool `json:"sdrsIoLoadBalanceEnabled,omitempty" tf:"sdrs_io_load_balance_enabled,omitempty"`

	// The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
	// +kubebuilder:validation:Optional
	SdrsIoLoadImbalanceThreshold *float64 `json:"sdrsIoLoadImbalanceThreshold,omitempty" tf:"sdrs_io_load_imbalance_threshold,omitempty"`

	// The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to move VMs off of a datastore.
	// +kubebuilder:validation:Optional
	SdrsIoReservableIopsThreshold *float64 `json:"sdrsIoReservableIopsThreshold,omitempty" tf:"sdrs_io_reservable_iops_threshold,omitempty"`

	// The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
	// +kubebuilder:validation:Optional
	SdrsIoReservablePercentThreshold *float64 `json:"sdrsIoReservablePercentThreshold,omitempty" tf:"sdrs_io_reservable_percent_threshold,omitempty"`

	// The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
	// +kubebuilder:validation:Optional
	SdrsIoReservableThresholdMode *string `json:"sdrsIoReservableThresholdMode,omitempty" tf:"sdrs_io_reservable_threshold_mode,omitempty"`

	// The storage DRS poll interval, in minutes.
	// +kubebuilder:validation:Optional
	SdrsLoadBalanceInterval *float64 `json:"sdrsLoadBalanceInterval,omitempty" tf:"sdrs_load_balance_interval,omitempty"`

	// Overrides the default automation settings when correcting storage and VM policy violations.
	// +kubebuilder:validation:Optional
	SdrsPolicyEnforcementAutomationLevel *string `json:"sdrsPolicyEnforcementAutomationLevel,omitempty" tf:"sdrs_policy_enforcement_automation_level,omitempty"`

	// Overrides the default automation settings when correcting affinity rule violations.
	// +kubebuilder:validation:Optional
	SdrsRuleEnforcementAutomationLevel *string `json:"sdrsRuleEnforcementAutomationLevel,omitempty" tf:"sdrs_rule_enforcement_automation_level,omitempty"`

	// Overrides the default automation settings when correcting disk space imbalances.
	// +kubebuilder:validation:Optional
	SdrsSpaceBalanceAutomationLevel *string `json:"sdrsSpaceBalanceAutomationLevel,omitempty" tf:"sdrs_space_balance_automation_level,omitempty"`

	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	// +kubebuilder:validation:Optional
	SdrsSpaceUtilizationThreshold *float64 `json:"sdrsSpaceUtilizationThreshold,omitempty" tf:"sdrs_space_utilization_threshold,omitempty"`

	// Overrides the default automation settings when generating recommendations for datastore evacuation.
	// +kubebuilder:validation:Optional
	SdrsVMEvacuationAutomationLevel *string `json:"sdrsVmEvacuationAutomationLevel,omitempty" tf:"sdrs_vm_evacuation_automation_level,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vspherejet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
