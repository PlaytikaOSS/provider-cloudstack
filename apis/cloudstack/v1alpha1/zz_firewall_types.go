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

type FirewallInitParameters struct {

	// The IP address ID for which to create the
	// firewall rules. Changing this forces a new resource to be created.
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// USE WITH CAUTION! If enabled all the firewall rules for
	// this IP address will be managed by this resource. This means it will delete
	// all firewall rules that are not in your config! (defaults false)
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// Specifies how much rules will be created or deleted
	// concurrently. (defaults 2)
	Parallelism *float64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// Can be specified multiple times. Each rule block supports
	// fields documented below. If managed = false at least one rule is required!
	Rule []FirewallRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FirewallObservation struct {

	// The IP address ID for which the firewall rules are created.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address ID for which to create the
	// firewall rules. Changing this forces a new resource to be created.
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// USE WITH CAUTION! If enabled all the firewall rules for
	// this IP address will be managed by this resource. This means it will delete
	// all firewall rules that are not in your config! (defaults false)
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// Specifies how much rules will be created or deleted
	// concurrently. (defaults 2)
	Parallelism *float64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// Can be specified multiple times. Each rule block supports
	// fields documented below. If managed = false at least one rule is required!
	Rule []FirewallRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FirewallParameters struct {

	// The IP address ID for which to create the
	// firewall rules. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// USE WITH CAUTION! If enabled all the firewall rules for
	// this IP address will be managed by this resource. This means it will delete
	// all firewall rules that are not in your config! (defaults false)
	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// Specifies how much rules will be created or deleted
	// concurrently. (defaults 2)
	// +kubebuilder:validation:Optional
	Parallelism *float64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// Can be specified multiple times. Each rule block supports
	// fields documented below. If managed = false at least one rule is required!
	// +kubebuilder:validation:Optional
	Rule []FirewallRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FirewallRuleInitParameters struct {

	// A CIDR list to allow access to the given ports.
	// +listType=set
	CidrList []*string `json:"cidrList,omitempty" tf:"cidr_list,omitempty"`

	// The ICMP code to allow. This can only be specified if
	// the protocol is ICMP.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to allow. This can only be specified if
	// the protocol is ICMP.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// List of ports and/or port ranges to allow. This can only
	// be specified if the protocol is TCP or UDP.
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The name of the protocol to allow. Valid options are:
	// tcp, udp and icmp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type FirewallRuleObservation struct {

	// A CIDR list to allow access to the given ports.
	// +listType=set
	CidrList []*string `json:"cidrList,omitempty" tf:"cidr_list,omitempty"`

	// The ICMP code to allow. This can only be specified if
	// the protocol is ICMP.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to allow. This can only be specified if
	// the protocol is ICMP.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// List of ports and/or port ranges to allow. This can only
	// be specified if the protocol is TCP or UDP.
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The name of the protocol to allow. Valid options are:
	// tcp, udp and icmp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +mapType=granular
	Uuids map[string]*string `json:"uuids,omitempty" tf:"uuids,omitempty"`
}

type FirewallRuleParameters struct {

	// A CIDR list to allow access to the given ports.
	// +kubebuilder:validation:Optional
	// +listType=set
	CidrList []*string `json:"cidrList" tf:"cidr_list,omitempty"`

	// The ICMP code to allow. This can only be specified if
	// the protocol is ICMP.
	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to allow. This can only be specified if
	// the protocol is ICMP.
	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// List of ports and/or port ranges to allow. This can only
	// be specified if the protocol is TCP or UDP.
	// +kubebuilder:validation:Optional
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The name of the protocol to allow. Valid options are:
	// tcp, udp and icmp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
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
	InitProvider FirewallInitParameters `json:"initProvider,omitempty"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Firewall is the Schema for the Firewalls API. Creates firewall rules for a given IP address.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddressId) || (has(self.initProvider) && has(self.initProvider.ipAddressId))",message="spec.forProvider.ipAddressId is a required parameter"
	Spec   FirewallSpec   `json:"spec"`
	Status FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
