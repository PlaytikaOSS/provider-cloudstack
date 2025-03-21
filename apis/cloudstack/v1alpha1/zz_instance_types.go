// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceInitParameters struct {

	// List of affinity group IDs to apply to this
	// instance.
	// +listType=set
	AffinityGroupIds []*string `json:"affinityGroupIds,omitempty" tf:"affinity_group_ids,omitempty"`

	// List of affinity group names to apply to
	// this instance.
	// +listType=set
	AffinityGroupNames []*string `json:"affinityGroupNames,omitempty" tf:"affinity_group_names,omitempty"`

	// destination Cluster ID to deploy the VM to - parameter available
	// for root admin only
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +mapType=granular
	Details map[string]*string `json:"details,omitempty" tf:"details,omitempty"`

	// The display name of the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// This determines if the instance is expunged when it is
	// destroyed (defaults false)
	Expunge *bool `json:"expunge,omitempty" tf:"expunge,omitempty"`

	// The group name of the instance.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// destination Host ID to deploy the VM to - parameter available
	// for root admin only
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// The IP address to assign to this instance. Changing
	// this forces a new resource to be created.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name of the SSH key pair that will be used to
	// access this instance.
	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// The name of the instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to connect this instance
	// to. Changing this forces a new resource to be created.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +mapType=granular
	Nicnetworklist map[string]*string `json:"nicnetworklist,omitempty" tf:"nicnetworklist,omitempty"`

	// destination Pod ID to deploy the VM to - parameter available for root admin only
	PodID *string `json:"podId,omitempty" tf:"pod_id,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The size of the root disk in gigabytes. The
	// root disk is resized on deploy. Only applies to template-based deployments.
	// Changing this forces a new resource to be created.
	RootDiskSize *float64 `json:"rootDiskSize,omitempty" tf:"root_disk_size,omitempty"`

	// List of security group IDs to apply to this
	// instance. Changing this forces a new resource to be created.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// List of security group names to apply to
	// this instance. Changing this forces a new resource to be created.
	// +listType=set
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// The name or ID of the service offering used
	// for this instance.
	ServiceOffering *string `json:"serviceOffering,omitempty" tf:"service_offering,omitempty"`

	// This determines if the instances is started after it
	// is created (defaults true)
	StartVM *bool `json:"startVm,omitempty" tf:"start_vm,omitempty"`

	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name or ID of the template used for this
	// instance. Changing this forces a new resource to be created.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// When set, will boot the instance in UEFI/Legacy mode (defaults false)
	Uefi *bool `json:"uefi,omitempty" tf:"uefi,omitempty"`

	// The user data to provide when launching the
	// instance. This can be either plain text or base64 encoded text.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The name or ID of the zone where this instance will be
	// created. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceObservation struct {

	// List of affinity group IDs to apply to this
	// instance.
	// +listType=set
	AffinityGroupIds []*string `json:"affinityGroupIds,omitempty" tf:"affinity_group_ids,omitempty"`

	// List of affinity group names to apply to
	// this instance.
	// +listType=set
	AffinityGroupNames []*string `json:"affinityGroupNames,omitempty" tf:"affinity_group_names,omitempty"`

	// destination Cluster ID to deploy the VM to - parameter available
	// for root admin only
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +mapType=granular
	Details map[string]*string `json:"details,omitempty" tf:"details,omitempty"`

	// The display name of the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// This determines if the instance is expunged when it is
	// destroyed (defaults false)
	Expunge *bool `json:"expunge,omitempty" tf:"expunge,omitempty"`

	// The group name of the instance.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// destination Host ID to deploy the VM to - parameter available
	// for root admin only
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// The instance ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address to assign to this instance. Changing
	// this forces a new resource to be created.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name of the SSH key pair that will be used to
	// access this instance.
	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// The name of the instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to connect this instance
	// to. Changing this forces a new resource to be created.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +mapType=granular
	Nicnetworklist map[string]*string `json:"nicnetworklist,omitempty" tf:"nicnetworklist,omitempty"`

	// destination Pod ID to deploy the VM to - parameter available for root admin only
	PodID *string `json:"podId,omitempty" tf:"pod_id,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The size of the root disk in gigabytes. The
	// root disk is resized on deploy. Only applies to template-based deployments.
	// Changing this forces a new resource to be created.
	RootDiskSize *float64 `json:"rootDiskSize,omitempty" tf:"root_disk_size,omitempty"`

	// List of security group IDs to apply to this
	// instance. Changing this forces a new resource to be created.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// List of security group names to apply to
	// this instance. Changing this forces a new resource to be created.
	// +listType=set
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// The name or ID of the service offering used
	// for this instance.
	ServiceOffering *string `json:"serviceOffering,omitempty" tf:"service_offering,omitempty"`

	// This determines if the instances is started after it
	// is created (defaults true)
	StartVM *bool `json:"startVm,omitempty" tf:"start_vm,omitempty"`

	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name or ID of the template used for this
	// instance. Changing this forces a new resource to be created.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// When set, will boot the instance in UEFI/Legacy mode (defaults false)
	Uefi *bool `json:"uefi,omitempty" tf:"uefi,omitempty"`

	// The user data to provide when launching the
	// instance. This can be either plain text or base64 encoded text.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The name or ID of the zone where this instance will be
	// created. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceParameters struct {

	// List of affinity group IDs to apply to this
	// instance.
	// +kubebuilder:validation:Optional
	// +listType=set
	AffinityGroupIds []*string `json:"affinityGroupIds,omitempty" tf:"affinity_group_ids,omitempty"`

	// List of affinity group names to apply to
	// this instance.
	// +kubebuilder:validation:Optional
	// +listType=set
	AffinityGroupNames []*string `json:"affinityGroupNames,omitempty" tf:"affinity_group_names,omitempty"`

	// destination Cluster ID to deploy the VM to - parameter available
	// for root admin only
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Details map[string]*string `json:"details,omitempty" tf:"details,omitempty"`

	// The display name of the instance.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// This determines if the instance is expunged when it is
	// destroyed (defaults false)
	// +kubebuilder:validation:Optional
	Expunge *bool `json:"expunge,omitempty" tf:"expunge,omitempty"`

	// The group name of the instance.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// destination Host ID to deploy the VM to - parameter available
	// for root admin only
	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// The IP address to assign to this instance. Changing
	// this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name of the SSH key pair that will be used to
	// access this instance.
	// +kubebuilder:validation:Optional
	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// The name of the instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to connect this instance
	// to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Nicnetworklist map[string]*string `json:"nicnetworklist,omitempty" tf:"nicnetworklist,omitempty"`

	// destination Pod ID to deploy the VM to - parameter available for root admin only
	// +kubebuilder:validation:Optional
	PodID *string `json:"podId,omitempty" tf:"pod_id,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The size of the root disk in gigabytes. The
	// root disk is resized on deploy. Only applies to template-based deployments.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RootDiskSize *float64 `json:"rootDiskSize,omitempty" tf:"root_disk_size,omitempty"`

	// List of security group IDs to apply to this
	// instance. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// List of security group names to apply to
	// this instance. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// The name or ID of the service offering used
	// for this instance.
	// +kubebuilder:validation:Optional
	ServiceOffering *string `json:"serviceOffering,omitempty" tf:"service_offering,omitempty"`

	// This determines if the instances is started after it
	// is created (defaults true)
	// +kubebuilder:validation:Optional
	StartVM *bool `json:"startVm,omitempty" tf:"start_vm,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name or ID of the template used for this
	// instance. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// When set, will boot the instance in UEFI/Legacy mode (defaults false)
	// +kubebuilder:validation:Optional
	Uefi *bool `json:"uefi,omitempty" tf:"uefi,omitempty"`

	// The user data to provide when launching the
	// instance. This can be either plain text or base64 encoded text.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The name or ID of the zone where this instance will be
	// created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Instance is the Schema for the Instances API. Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceOffering) || (has(self.initProvider) && has(self.initProvider.serviceOffering))",message="spec.forProvider.serviceOffering is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.template) || (has(self.initProvider) && has(self.initProvider.template))",message="spec.forProvider.template is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
