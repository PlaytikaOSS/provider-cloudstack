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

type IPAddressInitParameters struct {

	// This determines if the IP address should be transferable
	// across zones (defaults false)
	IsPortable *bool `json:"isPortable,omitempty" tf:"is_portable,omitempty"`

	// The ID of the network for which an IP address should
	// be acquired and associated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.Network
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the VPC for which an IP address should be
	// acquired and associated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The name or ID of the zone for which an IP address should be
	// acquired and associated. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type IPAddressObservation struct {

	// The ID of the acquired and associated IP address.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address that was acquired and associated.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// This determines if the IP address should be transferable
	// across zones (defaults false)
	IsPortable *bool `json:"isPortable,omitempty" tf:"is_portable,omitempty"`

	IsSourceNAT *bool `json:"isSourceNat,omitempty" tf:"is_source_nat,omitempty"`

	// The ID of the network for which an IP address should
	// be acquired and associated. Changing this forces a new resource to be created.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the VPC for which an IP address should be
	// acquired and associated. Changing this forces a new resource to be created.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The name or ID of the zone for which an IP address should be
	// acquired and associated. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type IPAddressParameters struct {

	// This determines if the IP address should be transferable
	// across zones (defaults false)
	// +kubebuilder:validation:Optional
	IsPortable *bool `json:"isPortable,omitempty" tf:"is_portable,omitempty"`

	// The ID of the network for which an IP address should
	// be acquired and associated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in cloudstack to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the VPC for which an IP address should be
	// acquired and associated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The name or ID of the zone for which an IP address should be
	// acquired and associated. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// IPAddressSpec defines the desired state of IPAddress
type IPAddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPAddressParameters `json:"forProvider"`
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
	InitProvider IPAddressInitParameters `json:"initProvider,omitempty"`
}

// IPAddressStatus defines the observed state of IPAddress.
type IPAddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPAddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IPAddress is the Schema for the IPAddresss API. Acquires and associates a public IP.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type IPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPAddressSpec   `json:"spec"`
	Status            IPAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPAddressList contains a list of IPAddresss
type IPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPAddress `json:"items"`
}

// Repository type metadata.
var (
	IPAddress_Kind             = "IPAddress"
	IPAddress_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPAddress_Kind}.String()
	IPAddress_KindAPIVersion   = IPAddress_Kind + "." + CRDGroupVersion.String()
	IPAddress_GroupVersionKind = CRDGroupVersion.WithKind(IPAddress_Kind)
)

func init() {
	SchemeBuilder.Register(&IPAddress{}, &IPAddressList{})
}
