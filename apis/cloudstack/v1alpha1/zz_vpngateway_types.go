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

type VPNGatewayInitParameters struct {

	// The ID of the VPC for which to create the VPN Gateway.
	// Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type VPNGatewayObservation struct {

	// The ID of the VPN Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The public IP address associated with the VPN Gateway.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// The ID of the VPC for which to create the VPN Gateway.
	// Changing this forces a new resource to be created.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPNGatewayParameters struct {

	// The ID of the VPC for which to create the VPN Gateway.
	// Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/playtika/provider-cloudstack/apis/cloudstack/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in cloudstack to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPNGatewaySpec defines the desired state of VPNGateway
type VPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayParameters `json:"forProvider"`
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
	InitProvider VPNGatewayInitParameters `json:"initProvider,omitempty"`
}

// VPNGatewayStatus defines the observed state of VPNGateway.
type VPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPNGateway is the Schema for the VPNGateways API. Creates a site to site VPN local gateway.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type VPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNGatewaySpec   `json:"spec"`
	Status            VPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayList contains a list of VPNGateways
type VPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGateway `json:"items"`
}

// Repository type metadata.
var (
	VPNGateway_Kind             = "VPNGateway"
	VPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGateway_Kind}.String()
	VPNGateway_KindAPIVersion   = VPNGateway_Kind + "." + CRDGroupVersion.String()
	VPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(VPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGateway{}, &VPNGatewayList{})
}
