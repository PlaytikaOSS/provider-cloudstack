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

type VPNCustomerGatewayInitParameters struct {

	// The CIDR block that needs to be routed through this gateway.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// If DPD is enabled for the related VPN connection (defaults false)
	Dpd *bool `json:"dpd,omitempty" tf:"dpd,omitempty"`

	// The ESP lifetime of phase 2 VPN connection to this
	// VPN Customer Gateway in seconds (defaults 86400)
	EspLifetime *float64 `json:"espLifetime,omitempty" tf:"esp_lifetime,omitempty"`

	// The ESP policy to use for this VPN Customer Gateway.
	EspPolicy *string `json:"espPolicy,omitempty" tf:"esp_policy,omitempty"`

	// The public IP address of the related VPN Gateway.
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// The IKE lifetime of phase 2 VPN connection to this
	// VPN Customer Gateway in seconds (defaults 86400)
	IkeLifetime *float64 `json:"ikeLifetime,omitempty" tf:"ike_lifetime,omitempty"`

	// The IKE policy to use for this VPN Customer Gateway.
	IkePolicy *string `json:"ikePolicy,omitempty" tf:"ike_policy,omitempty"`

	// The IPSEC pre-shared key used for this gateway.
	IpsecPsk *string `json:"ipsecPsk,omitempty" tf:"ipsec_psk,omitempty"`

	// The name of the VPN Customer Gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name or ID of the project to create this VPN Customer
	// Gateway in. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type VPNCustomerGatewayObservation struct {

	// The CIDR block that needs to be routed through this gateway.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// If DPD is enabled for the related VPN connection (defaults false)
	Dpd *bool `json:"dpd,omitempty" tf:"dpd,omitempty"`

	// The ESP lifetime of phase 2 VPN connection to this
	// VPN Customer Gateway in seconds (defaults 86400)
	EspLifetime *float64 `json:"espLifetime,omitempty" tf:"esp_lifetime,omitempty"`

	// The ESP policy to use for this VPN Customer Gateway.
	EspPolicy *string `json:"espPolicy,omitempty" tf:"esp_policy,omitempty"`

	// The public IP address of the related VPN Gateway.
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// The ID of the VPN Customer Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IKE lifetime of phase 2 VPN connection to this
	// VPN Customer Gateway in seconds (defaults 86400)
	IkeLifetime *float64 `json:"ikeLifetime,omitempty" tf:"ike_lifetime,omitempty"`

	// The IKE policy to use for this VPN Customer Gateway.
	IkePolicy *string `json:"ikePolicy,omitempty" tf:"ike_policy,omitempty"`

	// The IPSEC pre-shared key used for this gateway.
	IpsecPsk *string `json:"ipsecPsk,omitempty" tf:"ipsec_psk,omitempty"`

	// The name of the VPN Customer Gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name or ID of the project to create this VPN Customer
	// Gateway in. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type VPNCustomerGatewayParameters struct {

	// The CIDR block that needs to be routed through this gateway.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// If DPD is enabled for the related VPN connection (defaults false)
	// +kubebuilder:validation:Optional
	Dpd *bool `json:"dpd,omitempty" tf:"dpd,omitempty"`

	// The ESP lifetime of phase 2 VPN connection to this
	// VPN Customer Gateway in seconds (defaults 86400)
	// +kubebuilder:validation:Optional
	EspLifetime *float64 `json:"espLifetime,omitempty" tf:"esp_lifetime,omitempty"`

	// The ESP policy to use for this VPN Customer Gateway.
	// +kubebuilder:validation:Optional
	EspPolicy *string `json:"espPolicy,omitempty" tf:"esp_policy,omitempty"`

	// The public IP address of the related VPN Gateway.
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// The IKE lifetime of phase 2 VPN connection to this
	// VPN Customer Gateway in seconds (defaults 86400)
	// +kubebuilder:validation:Optional
	IkeLifetime *float64 `json:"ikeLifetime,omitempty" tf:"ike_lifetime,omitempty"`

	// The IKE policy to use for this VPN Customer Gateway.
	// +kubebuilder:validation:Optional
	IkePolicy *string `json:"ikePolicy,omitempty" tf:"ike_policy,omitempty"`

	// The IPSEC pre-shared key used for this gateway.
	// +kubebuilder:validation:Optional
	IpsecPsk *string `json:"ipsecPsk,omitempty" tf:"ipsec_psk,omitempty"`

	// The name of the VPN Customer Gateway.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name or ID of the project to create this VPN Customer
	// Gateway in. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// VPNCustomerGatewaySpec defines the desired state of VPNCustomerGateway
type VPNCustomerGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNCustomerGatewayParameters `json:"forProvider"`
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
	InitProvider VPNCustomerGatewayInitParameters `json:"initProvider,omitempty"`
}

// VPNCustomerGatewayStatus defines the observed state of VPNCustomerGateway.
type VPNCustomerGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNCustomerGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPNCustomerGateway is the Schema for the VPNCustomerGateways API. Creates a site to site VPN local customer gateway.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type VPNCustomerGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidr) || (has(self.initProvider) && has(self.initProvider.cidr))",message="spec.forProvider.cidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.espPolicy) || (has(self.initProvider) && has(self.initProvider.espPolicy))",message="spec.forProvider.espPolicy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.gateway) || (has(self.initProvider) && has(self.initProvider.gateway))",message="spec.forProvider.gateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ikePolicy) || (has(self.initProvider) && has(self.initProvider.ikePolicy))",message="spec.forProvider.ikePolicy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipsecPsk) || (has(self.initProvider) && has(self.initProvider.ipsecPsk))",message="spec.forProvider.ipsecPsk is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VPNCustomerGatewaySpec   `json:"spec"`
	Status VPNCustomerGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNCustomerGatewayList contains a list of VPNCustomerGateways
type VPNCustomerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNCustomerGateway `json:"items"`
}

// Repository type metadata.
var (
	VPNCustomerGateway_Kind             = "VPNCustomerGateway"
	VPNCustomerGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNCustomerGateway_Kind}.String()
	VPNCustomerGateway_KindAPIVersion   = VPNCustomerGateway_Kind + "." + CRDGroupVersion.String()
	VPNCustomerGateway_GroupVersionKind = CRDGroupVersion.WithKind(VPNCustomerGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNCustomerGateway{}, &VPNCustomerGatewayList{})
}
