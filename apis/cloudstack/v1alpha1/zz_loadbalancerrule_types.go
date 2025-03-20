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

type LoadBalancerRuleInitParameters struct {

	// Load balancer rule algorithm (source, roundrobin,
	// leastconn). Changing this forces a new resource to be created.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The load balancer rule ID.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The description of the load balancer rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Public IP address ID from where the network
	// traffic will be load balanced from. Changing this forces a new resource
	// to be created.
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// List of instance IDs to assign to the load balancer
	// rule. Changing this forces a new resource to be created.
	// +listType=set
	MemberIds []*string `json:"memberIds,omitempty" tf:"member_ids,omitempty"`

	// Name of the loadbalancer rule.
	// Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The network ID this rule will be created for.
	// Required when public IP address is not associated with any network yet
	// (VPC case).
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.Network
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The private port of the private IP address
	// (virtual machine) where the network traffic will be load balanced to.
	// Changing this forces a new resource to be created.
	PrivatePort *float64 `json:"privatePort,omitempty" tf:"private_port,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Load balancer protocol (tcp, udp, tcp-proxy).
	// Changing this forces a new resource to be created.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The public port from where the network traffic
	// will be load balanced from. Changing this forces a new resource to be
	// created.
	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`
}

type LoadBalancerRuleObservation struct {

	// Load balancer rule algorithm (source, roundrobin,
	// leastconn). Changing this forces a new resource to be created.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The load balancer rule ID.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The description of the load balancer rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The load balancer rule ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Public IP address ID from where the network
	// traffic will be load balanced from. Changing this forces a new resource
	// to be created.
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// List of instance IDs to assign to the load balancer
	// rule. Changing this forces a new resource to be created.
	// +listType=set
	MemberIds []*string `json:"memberIds,omitempty" tf:"member_ids,omitempty"`

	// Name of the loadbalancer rule.
	// Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The network ID this rule will be created for.
	// Required when public IP address is not associated with any network yet
	// (VPC case).
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The private port of the private IP address
	// (virtual machine) where the network traffic will be load balanced to.
	// Changing this forces a new resource to be created.
	PrivatePort *float64 `json:"privatePort,omitempty" tf:"private_port,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Load balancer protocol (tcp, udp, tcp-proxy).
	// Changing this forces a new resource to be created.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The public port from where the network traffic
	// will be load balanced from. Changing this forces a new resource to be
	// created.
	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`
}

type LoadBalancerRuleParameters struct {

	// Load balancer rule algorithm (source, roundrobin,
	// leastconn). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The load balancer rule ID.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The description of the load balancer rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Public IP address ID from where the network
	// traffic will be load balanced from. Changing this forces a new resource
	// to be created.
	// +kubebuilder:validation:Optional
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// List of instance IDs to assign to the load balancer
	// rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	MemberIds []*string `json:"memberIds,omitempty" tf:"member_ids,omitempty"`

	// Name of the loadbalancer rule.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The network ID this rule will be created for.
	// Required when public IP address is not associated with any network yet
	// (VPC case).
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The private port of the private IP address
	// (virtual machine) where the network traffic will be load balanced to.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrivatePort *float64 `json:"privatePort,omitempty" tf:"private_port,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Load balancer protocol (tcp, udp, tcp-proxy).
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The public port from where the network traffic
	// will be load balanced from. Changing this forces a new resource to be
	// created.
	// +kubebuilder:validation:Optional
	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`
}

// LoadBalancerRuleSpec defines the desired state of LoadBalancerRule
type LoadBalancerRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerRuleParameters `json:"forProvider"`
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
	InitProvider LoadBalancerRuleInitParameters `json:"initProvider,omitempty"`
}

// LoadBalancerRuleStatus defines the observed state of LoadBalancerRule.
type LoadBalancerRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LoadBalancerRule is the Schema for the LoadBalancerRules API. Creates a load balancer rule.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type LoadBalancerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.algorithm) || (has(self.initProvider) && has(self.initProvider.algorithm))",message="spec.forProvider.algorithm is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddressId) || (has(self.initProvider) && has(self.initProvider.ipAddressId))",message="spec.forProvider.ipAddressId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.memberIds) || (has(self.initProvider) && has(self.initProvider.memberIds))",message="spec.forProvider.memberIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privatePort) || (has(self.initProvider) && has(self.initProvider.privatePort))",message="spec.forProvider.privatePort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicPort) || (has(self.initProvider) && has(self.initProvider.publicPort))",message="spec.forProvider.publicPort is a required parameter"
	Spec   LoadBalancerRuleSpec   `json:"spec"`
	Status LoadBalancerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerRuleList contains a list of LoadBalancerRules
type LoadBalancerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerRule `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerRule_Kind             = "LoadBalancerRule"
	LoadBalancerRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerRule_Kind}.String()
	LoadBalancerRule_KindAPIVersion   = LoadBalancerRule_Kind + "." + CRDGroupVersion.String()
	LoadBalancerRule_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerRule{}, &LoadBalancerRuleList{})
}
